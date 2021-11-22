package repository

import (
	"context"

	"github.com/javierpr71/mastermind/models"
)

// IGame interface
type IGame interface {
	Create(ctx context.Context, game *models.Game) (string, error)
	GetByID(ctx context.Context, id string) (*models.Game, error)
	Update(ctx context.Context, game *models.Game) (*models.Game, error)
}
