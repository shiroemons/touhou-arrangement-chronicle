// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Node interface {
	IsNode()
	GetID() string
}

type ProductType string

const (
	ProductTypePc98                ProductType = "pc98"
	ProductTypeWindows             ProductType = "windows"
	ProductTypeZunsMusicCollection ProductType = "zuns_music_collection"
	ProductTypeAkyusUntouchedScore ProductType = "akyus_untouched_score"
	ProductTypeCommercialBooks     ProductType = "commercial_books"
	ProductTypeOther               ProductType = "other"
)

var AllProductType = []ProductType{
	ProductTypePc98,
	ProductTypeWindows,
	ProductTypeZunsMusicCollection,
	ProductTypeAkyusUntouchedScore,
	ProductTypeCommercialBooks,
	ProductTypeOther,
}

func (e ProductType) IsValid() bool {
	switch e {
	case ProductTypePc98, ProductTypeWindows, ProductTypeZunsMusicCollection, ProductTypeAkyusUntouchedScore, ProductTypeCommercialBooks, ProductTypeOther:
		return true
	}
	return false
}

func (e ProductType) String() string {
	return string(e)
}

func (e *ProductType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductType", str)
	}
	return nil
}

func (e ProductType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
