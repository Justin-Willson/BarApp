package dal

import (
	"github.com/Justin-Willson/BarApp/domain"
)

const COCKTAIL_TABLE_NAME = "cocktails"

func GetCocktails() []*domain.Cocktail {
	return GetAll[domain.Cocktail](COCKTAIL_TABLE_NAME)
}

func AddCocktail(cocktail domain.Cocktail) {
	InsertOne(cocktail, COCKTAIL_TABLE_NAME)
}
