package domain


type Recipe struct {
	// todo id of some kind
	Name string
	Tags []string
	Ingredients []Ingredient
	isVegan bool
	isVeggie bool
	servings int

}
