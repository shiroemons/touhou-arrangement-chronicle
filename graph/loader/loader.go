package loader

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader"

	"github.com/shiroemons/touhou-arrangement-chronicle/internal/infrastructure/database"
	"github.com/shiroemons/touhou-arrangement-chronicle/internal/repository"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

type Loaders struct {
	ProductLoader *dataloader.Loader
}

func NewLoaders(db *database.DB) *Loaders {
	productLoader := &ProductLoader{
		productRepo: repository.NewProduct(db),
	}
	loaders := &Loaders{
		ProductLoader: dataloader.NewBatchedLoader(productLoader.BatchGetProducts),
	}
	return loaders
}

func Middleware(loaders *Loaders) gin.HandlerFunc {
	loaders.ProductLoader.ClearAll()
	// return a middleware that injects the loader to the request context
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), loadersKey, loaders)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
