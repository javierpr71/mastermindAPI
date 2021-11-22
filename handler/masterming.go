package handler

import (
	"github.com/javierpr71/mastermind/driver"
	"github.com/javierpr71/mastermind/repository"
	"github.com/javierpr71/mastermind/repository/game"
	log "github.com/sirupsen/logrus"
)

func NewMasterMindHandler(logger *log.Logger, db *driver.DB) *Game {
	return &Game{
		Repo: game.NewRedisGameRepo(logger, db.SQL),
	}
}

// Game ...
type Game struct {
	Repo repository.IGame
}
