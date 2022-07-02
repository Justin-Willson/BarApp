package dal

import (
	"github.com/Justin-Willson/BarApp/domain"
)

var ingredients = []domain.Ingredient{
	{Name: "Gin", IsInStock: true, Notes: "Better than vodka"},
	{Name: "Dry Vermouth", IsInStock: true, Notes: "Great for martinis"},
}
var gin = ingredients[0]
var vermouth = ingredients[1]

func GetIngredients() []domain.Ingredient {
	return ingredients
}

func AddIngredient(ingredient domain.Ingredient) {
	ingredients = append(ingredients, ingredient)
}
