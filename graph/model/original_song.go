package model

import "github.com/uptrace/bun"

type OriginalSong struct {
	bun.BaseModel `bun:"table:original_songs,alias:os"`

	ID          string   `json:"id" bun:",pk"`
	ProductID   string   `json:"productID" bun:"product_id"`
	Product     *Product `json:"product" bun:"-"`
	Name        string   `json:"name" bun:"name"`
	Composer    string   `json:"composer" bun:"composer"`
	Arranger    string   `json:"arranger" bun:"arranger"`
	TrackNumber int      `json:"trackNumber" bun:"track_number"`
	Duplicate   bool     `json:"duplicate" bun:"is_duplicate"`
}
