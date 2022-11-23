package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spkg/bom"
	"github.com/uptrace/bun"

	"github.com/shiroemons/touhou-arrangement-chronicle/internal/infrastructure/database"
)

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID           string  `csv:"id" bun:",pk"`
	Name         string  `csv:"name" bun:"name,nullzero,notnull"`
	ShortName    string  `csv:"short_name" bun:"short_name,nullzero,notnull"`
	ProductType  string  `csv:"product_type" bun:"product_type,nullzero,notnull"`
	SeriesNumber float64 `csv:"series_number" bun:"series_number,nullzero,notnull"`
}

type OriginalSong struct {
	bun.BaseModel `bun:"table:original_songs,alias:os"`

	ID          string `csv:"id" bun:",pk"`
	ProductID   string `csv:"product_id" bun:"product_id,nullzero,notnull"`
	Name        string `csv:"name" bun:"name,nullzero,notnull"`
	Composer    string `csv:"composer" bun:"composer,nullzero,notnull,default:''"`
	Arranger    string `csv:"arranger" bun:"arranger,nullzero,notnull,default:''"`
	TrackNumber int    `csv:"track_number" bun:"track_number,nullzero,notnull"`
	Duplicate   bool   `csv:"is_duplicate" bun:"is_duplicate,notnull"`
}

func main() {
	ctx := context.Background()
	db := database.New()
	defer db.Close()

	fn := func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(bom.NewReader(in)) // BOMの回避
		r.Comma = '\t'                        // 区切り文字をタブに変更
		r.Comment = '#'                       // #で始まる行はコメントと見なしスキップ
		return r
	}
	gocsv.SetCSVReader(fn)

	importProducts(ctx, db)
	importOriginalSongs(ctx, db)
}

func importProducts(ctx context.Context, db *database.DB) {
	log.Println("start products import.")

	f, err := os.Open("./data/products.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []Product
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	_, err = db.NewInsert().Model(&lines).
		On("CONFLICT (id) DO UPDATE").
		Set("name = EXCLUDED.name").
		Set("short_name = EXCLUDED.short_name").
		Set("product_type = EXCLUDED.product_type").
		Set("series_number = EXCLUDED.series_number").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("finish products import.")
}

func importOriginalSongs(ctx context.Context, db *database.DB) {
	log.Println("start original_songs import.")

	f, err := os.Open("./data/original_songs.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []OriginalSong
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	_, err = db.NewInsert().Model(&lines).
		On("CONFLICT (id) DO UPDATE").
		Set("product_id = EXCLUDED.product_id").
		Set("name = EXCLUDED.name").
		Set("composer = EXCLUDED.composer").
		Set("arranger = EXCLUDED.arranger").
		Set("track_number = EXCLUDED.track_number").
		Set("is_duplicate = EXCLUDED.is_duplicate").
		Exec(ctx)
	if err != nil {
		panic(err)
	}

	log.Println("finish original_songs import.")
}
