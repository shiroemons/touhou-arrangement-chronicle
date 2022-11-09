package model

import "github.com/uptrace/bun"

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID           string  `json:"id" bun:",pk"`
	Title        string  `json:"title" bun:"text"`
	ShortTitle   string  `json:"shortTitle" bun:"short_title"`
	ProductType  string  `json:"productType" bun:"product_type"`
	SeriesNumber float64 `json:"seriesNumber" bun:"series_number"`
}
