package model

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID           string      `json:"id" bun:",pk"`
	Name         string      `json:"name" bun:"name"`
	ShortName    string      `json:"shortName" bun:"short_name"`
	ProductType  ProductType `json:"productType" bun:"product_type"`
	SeriesNumber float64     `json:"seriesNumber" bun:"series_number"`
}

func (Product) IsNode()            {}
func (this Product) GetID() string { return this.ID }
