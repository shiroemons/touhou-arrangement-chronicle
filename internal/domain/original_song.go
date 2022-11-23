package domain

import (
	"context"

	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
)

type OriginalSong interface {
	All(ctx context.Context) ([]*model.OriginalSong, error)
}
