package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shiroemons/touhou-arrangement-chronicle/graph/generated"
	"github.com/shiroemons/touhou-arrangement-chronicle/graph/model"
)

// Product is the resolver for the product field.
func (r *originalSongResolver) Product(ctx context.Context, obj *model.OriginalSong) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// OriginalSongs is the resolver for the originalSongs field.
func (r *queryResolver) OriginalSongs(ctx context.Context) ([]*model.OriginalSong, error) {
	panic(fmt.Errorf("not implemented: OriginalSongs - originalSongs"))
}

// OriginalSong returns generated.OriginalSongResolver implementation.
func (r *Resolver) OriginalSong() generated.OriginalSongResolver { return &originalSongResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type originalSongResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
