package repository

import (
	"context"

	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
	"github.com/shiroemons/touhou-arrangement-chronicle/internal/infrastructure/database"
)

type OriginalSong struct {
	db *database.DB
}

func NewOriginalSong(db *database.DB) *OriginalSong {
	return &OriginalSong{db: db}
}

func (os *OriginalSong) All(ctx context.Context) ([]*model.OriginalSong, error) {
	originalSongs := make([]*model.OriginalSong, 0)
	err := os.db.NewSelect().Model(&originalSongs).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return originalSongs, nil
}
