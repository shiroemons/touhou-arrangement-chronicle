package loader

import (
	"context"
	"fmt"
	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
	"log"

	"github.com/graph-gophers/dataloader"

	"github.com/shiroemons/touhou-arrangement-chronicle/internal/repository"
)

type ProductLoader struct {
	productRepo *repository.Product
}

func (p *ProductLoader) BatchGetProducts(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	productIDs := make([]string, len(keys))
	for ix, key := range keys {
		productIDs[ix] = key.String()
	}

	productByID, err := p.productRepo.GetMapInIDs(ctx, productIDs)
	if err != nil {
		err = fmt.Errorf("fail get products, %w", err)
		log.Printf("%v\n", err)
		return nil
	}

	output := make([]*dataloader.Result, len(keys))
	for index, productKey := range keys {
		product, ok := productByID[productKey.String()]
		if ok {
			output[index] = &dataloader.Result{Data: product, Error: nil}
		} else {
			err = fmt.Errorf("product not found %s", productKey.String())
			output[index] = &dataloader.Result{Data: nil, Error: err}
		}
	}
	return output
}

func LoadProduct(ctx context.Context, productID string) (*model.Product, error) {
	loaders := GetLoaders(ctx)
	thunk := loaders.ProductLoader.Load(ctx, dataloader.StringKey(productID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	product := result.(*model.Product)
	return product, nil
}
