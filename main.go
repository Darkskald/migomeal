package main

import (
	"encoding/json"
	"migomeal/domain"
	"migomeal/rest"
	"net/http"
)

// todo; everything that is vegan shall automatically become vegetarian as well
// todo: the population of recipes in this way is extremelly tedious, this must be done by a name look-up


func fakeIngredients() domain.Ingredient {
	var tomato = domain.Ingredient{
		Food: domain.Food{"tomato", true, true},
		Amount: domain.Amount{
			Number: 4,
			Unit:   domain.Unit{"pcs", nil},
		},
	}
	return tomato
}

func mimicRecipe() func(w http.ResponseWriter, r *http.Request) {
	var recipe = domain.Recipe{}
	recipe.Name = "Chil con carne"
	recipe.Servings = 4
	recipe.Instructions = "Do some cooking and you are most probably done!"
	bytes,_ := json.Marshal(&recipe)
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(bytes)
	}
}

func main(){
	adapter := rest.NewAdapter()
	adapter.HandleFunc("/recipe", mimicRecipe())
	adapter.ListenAndServe()
}
