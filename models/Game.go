package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

const CODE_LEN int = 4
const VALUES string = "RBYGWO"

type RoundParam struct {
	ID    uuid.UUID `json:"id"`
	Guess string    `json:"guess"`
}

type Round struct {
	Round  int    `json:"round"`
	Guess  string `json:"guess"`
	Whites int    `json:"whites"`
	Blacks int    `json:"blacks"`
}

type GameShort struct {
	ID     string  `json:"id"`
	Rounds []Round `json:"rounds,omitempty"`
	Result string  `json:"result,omitempty"`
}

type Game struct {
	GameShort
	Code string `json:"code"`
}

func NewGame(player string) *Game {
	return &Game{
		GameShort: GameShort{
			ID:     uuid.NewString(),
			Rounds: []Round{},
		},
		Code: generateCode(),
	}
}

func generateCode() string {
	var result [CODE_LEN]byte
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < CODE_LEN; i++ {
		result[i] = byte(VALUES[rand.Intn(len(VALUES))])
	}
	return string(result[:])
}
