package repository

import (
	"context"

	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
	"github.com/shiroemons/touhou-arrangement-chronicle/internal/infrastructure/database"
)

type Product struct {
	db *database.DB
}

func NewProduct(db *database.DB) *Product {
	return &Product{db: db}
}

func (p *Product) All(ctx context.Context) ([]*model.Product, error) {
	products := make([]*model.Product, 0)
	err := p.db.NewSelect().Model(&products).Order("id ASC").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
