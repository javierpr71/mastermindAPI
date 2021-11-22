package game

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis"
	models "github.com/javierpr71/mastermind/models"
	pRepo "github.com/javierpr71/mastermind/repository"
	log "github.com/sirupsen/logrus"
)

// NewRedisGameRepo retunrs implement of cotact repository interface
func NewRedisGameRepo(logger *log.Logger, Conn *redis.Client) pRepo.IGame {
	return &redisGameRepo{
		logger: logger,
		Conn:   Conn,
	}
}

type redisGameRepo struct {
	logger *log.Logger
	Conn   *redis.Client
}

func (m *redisGameRepo) Create(ctx context.Context, game *models.Game) (string, error) {

	json, err := json.Marshal(game)
	if err != nil {
		m.logger.Errorf("Create: error: %v", err)
		return "", err
	}

	m.logger.Infof("Create: Game ID: %s", game.ID)
	err = m.Conn.Set(game.ID, json, 0).Err()
	if err != nil {
		m.logger.Errorf("Create: error: %v", err)
		return "", err
	}
	return game.ID, nil
}

// GetById get a game by id
func (m *redisGameRepo) GetByID(ctx context.Context, id string) (*models.Game, error) {

	val, err := m.Conn.Get(id).Result()
	if err != nil {
		m.logger.Errorf("GetByID: error: %v", err)
		return nil, err
	}

	var game *models.Game
	err = json.Unmarshal([]byte(val), &game)
	if err != nil {
		m.logger.Errorf("GetByID: error: %v", err)
		return nil, err
	}

	return game, nil
}

// Update update game
func (m *redisGameRepo) Update(ctx context.Context, game *models.Game) (*models.Game, error) {

	json, err := json.Marshal(game)
	if err != nil {
		m.logger.Errorf("Update: error: %v", err)
		return nil, err
	}

	m.logger.Infof("Update: Game ID: %s", game.ID)
	err = m.Conn.Set(game.ID, json, 0).Err()
	if err != nil {
		m.logger.Errorf("Update: error: %v", err)
		return nil, err
	}
	return game, nil
}
