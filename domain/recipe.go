package domain

type Recipe struct {
	// todo id of some kind
	// todo some kind of timing
	Name         string
	Tags         []string
	Ingredients  []Ingredient
	IsVegan      bool
	IsVeggie     bool
	Servings     int
	Instructions string
}

func (r Recipe) TestIfVegan() bool {
	for _, v := range r.Ingredients {
		if !v.Food.IsVegan {
			return false
		}
	}
	return true
}

func (r Recipe) TestIfVeggie() bool {
	for _, v := range r.Ingredients {
		if !v.Food.IsVeggie {
			return false
		}
	}
	return true
}

// Scale returns a new Recipe with the desired number of servings.
// todo: this requires unit testing and is a good exercise for that
func (r Recipe) Scale(new_servings int) Recipe {
	var factor = float32(new_servings) / float32(r.Servings)
	var scaled = r
	for index := range scaled.Ingredients {
		scaled.Ingredients[index].Amount.Number *= factor
	}
	scaled.Servings = new_servings
	return scaled

}

// todo: scale servings
