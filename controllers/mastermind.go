package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/javierpr71/mastermind/models"
	"github.com/javierpr71/mastermind/responses"
)

const MARK_CHAR byte = 0
const MAX_ROUNDS int = 10

func (server *Server) NewGame(w http.ResponseWriter, r *http.Request) {

	game := models.NewGame("Javier")

	id, err := server.pGameHandler.Repo.Create(r.Context(), game)
	if err != nil {
		responses.ERROR(r, w, http.StatusNoContent, err)
		return
	}

	responses.JSON(w, http.StatusOK, &models.GameShort{ID: id})
}

func (server *Server) RoundStatus(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	game, err := server.pGameHandler.Repo.GetByID(r.Context(), id)
	if err != nil {
		responses.ERROR(r, w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusOK, game.GameShort)
}

func (server *Server) Round(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(r, w, http.StatusUnprocessableEntity, err)
	}
	defer r.Body.Close()

	var roundParam *models.RoundParam
	if roundParam, err = getRoundParam(body); err != nil {
		responses.ERROR(r, w, http.StatusUnauthorized, err)
		return
	}

	server.logger.Infof("Round: Id: %s. Guess: %s", roundParam.ID, roundParam.Guess)

	// Get Game By Id from DB
	game, err := server.pGameHandler.Repo.GetByID(r.Context(), roundParam.ID.String())
	if err != nil {
		responses.ERROR(r, w, http.StatusInternalServerError, err)
		return
	}
	// Check if game is ended
	if game.Result != "" {
		responses.JSON(w, http.StatusOK, game)
		return
	}

	// Check guess for this round
	game.Rounds, game.Result = checkRound(game.Rounds, game.Code, roundParam.Guess)

	// Update game in DB with this round
	_, err = server.pGameHandler.Repo.Update(r.Context(), game)
	if err != nil {
		responses.ERROR(r, w, http.StatusInternalServerError, err)
		return
	}
	result := game.Rounds[len(game.Rounds)-1]
	server.logger.Infof("Game: $s, Round: %v", game.ID, result)
	responses.JSON(w, http.StatusOK, result)
}

func getRoundParam(body []byte) (*models.RoundParam, error) {
	roundParam := &models.RoundParam{}
	if err := json.Unmarshal(body, &roundParam); err != nil {
		return nil, err
	}
	return roundParam, nil
}

func checkRound(rounds []models.Round, code, guess string) ([]models.Round, string) {
	whites, blacks := check([]byte(code), []byte(guess))
	round := models.Round{
		Round:  len(rounds) + 1,
		Guess:  guess,
		Whites: whites,
		Blacks: blacks,
	}
	newRounds := append(rounds, round)
	// Check if this round end the game
	var result string
	if blacks == 4 {
		result = "You Win!"
	} else if len(rounds) >= MAX_ROUNDS {
		result = "You Loose!"
	}
	return newRounds, result
}

func check(code, guess []byte) (w, b int) {

	if len(code) != len(guess) {
		return 0, 0
	}

	// Check for whites
	for k, v := range guess {
		if v == code[k] {
			b++
			code[k] = MARK_CHAR
			guess[k] = MARK_CHAR
		}
	}

	// Check for blacks
	for _, v := range guess {
		if v == MARK_CHAR {
			continue
		}
		for k1, v1 := range code {
			if v == v1 {
				w++
				code[k1] = MARK_CHAR
				break
			}
		}
	}
	return w, b
}
