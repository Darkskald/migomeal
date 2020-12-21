package domain


type Recipe struct {
	// todo id of some kind
	Name string
	Tags []string
	Ingredients []Ingredient
	IsVegan bool
	IsVeggie bool
	Servings int
}

func (r Recipe) TestIfVegan () bool {
	for _, v := range r.Ingredients {
		if !v.Food.isVegan {
			return false
		}
	}
	return true
}

func (r Recipe) TestIfVeggie () bool {
	for _, v := range r.Ingredients {
		if !v.Food.isVeggie {
			return false
		}
	}
	return true
}

// todo: scale servings
