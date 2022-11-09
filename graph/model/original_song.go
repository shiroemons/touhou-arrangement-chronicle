package model

import "github.com/uptrace/bun"

type OriginalSong struct {
	bun.BaseModel `bun:"table:original_songs,alias:os"`

	ID          string   `json:"id" bun:",pk"`
	Product     *Product `json:"product"`
	Title       string   `json:"title" bun:"title"`
	Composer    string   `json:"composer" bun:"composer"`
	Arranger    string   `json:"arranger" bun:"arranger"`
	TrackNumber int      `json:"trackNumber" bun:"track_number"`
	Duplicate   bool     `json:"duplicate" bun:"is_duplicate"`
}
