package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/shiroemons/touhou-arrangement-chronicle/internal/infrastructure/database"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *database.DB
}
