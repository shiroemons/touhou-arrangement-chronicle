package main

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/spkg/bom"
	"io"
	"log"
	"os"
)

type Product struct {
	ID           string `csv:"id"`
	Name         string `csv:"name"`
	ShortName    string `csv:"short_name"`
	ProductType  string `csv:"product_type"`
	SeriesNumber string `csv:"series_number"`
}

type OriginalSong struct {
	ID          string `csv:"id"`
	ProductID   string `csv:"product_id"`
	Name        string `csv:"name"`
	Composer    string `csv:"composer"`
	Arranger    string `csv:"arranger"`
	TrackNumber int64  `csv:"track_number"`
	Duplicate   bool   `csv:"is_duplicate"`
}

func main() {
	fn := func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(bom.NewReader(in)) // BOMの回避
		r.Comma = '\t'                        // 区切り文字をタブに変更
		r.Comment = '#'                       // #で始まる行はコメントと見なしスキップ
		return r
	}
	gocsv.SetCSVReader(fn)

	products()
	originalSongs()
}

func products() {
	f, err := os.Open("./data/products.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []Product
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	for _, v := range lines {
		fmt.Printf("%+v\n", v)
	}
}

func originalSongs() {
	f, err := os.Open("./data/original_songs.tsv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []OriginalSong
	if err = gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	for _, v := range lines {
		fmt.Printf("%+v\n", v)
	}
}
