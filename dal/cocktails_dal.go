package dal

import (
	"github.com/Justin-Willson/BarApp/domain"
)

var ginPair = domain.IngredientQuantityPair{Ingredient: gin, Quantity: 3.0, Units: "ounces"}
var vermouthPair = domain.IngredientQuantityPair{Ingredient: vermouth, Quantity: 0.75, Units: "ounces"}

var martini = domain.Cocktail{Ingredients: []domain.IngredientQuantityPair{ginPair, vermouthPair},
	Name:  "Martini",
	Notes: "The best cocktial",
}
var cocktails = []domain.Cocktail{martini}

func GetCocktails() []domain.Cocktail {
	return cocktails
}

func AddCocktail(cocktial domain.Cocktail) {
	cocktails = append(cocktails, cocktial)
}
