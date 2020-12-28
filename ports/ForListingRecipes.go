package ports

import "migomeal/domain"

type ForListingRecipes interface {
	ListRecipes() ([]domain.Recipe, error)
}
