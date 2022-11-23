package repository

import (
	"context"

	"github.com/uptrace/bun"

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

func (p *Product) GetMapInIDs(ctx context.Context, ids []string) (map[string]*model.Product, error) {
	products := make([]*model.Product, 0)
	err := p.db.NewSelect().Model(&products).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, err
	}

	productById := map[string]*model.Product{}
	for _, product := range products {
		p := product
		productById[p.ID] = p
	}
	return productById, nil
}
