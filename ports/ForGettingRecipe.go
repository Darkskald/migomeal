package ports

import "migomeal/domain"

type ForGettingRecipe interface {
	GetRecipe(id string) (domain.Recipe, error)
}
