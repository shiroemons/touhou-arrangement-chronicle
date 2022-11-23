package domain

import (
	"context"

	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
)

type Product interface {
	All(ctx context.Context) ([]*model.Product, error)
}
