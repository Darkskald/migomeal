package ports

import "migomeal/domain"

type ForAddingRecipe interface {
	AddRecipe(recipe domain.Recipe) (domain.Recipe, error)
}
