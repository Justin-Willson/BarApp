package dal

import (
	"github.com/Justin-Willson/BarApp/domain"
)

const INGREDIENTS_TABLE_NAME = "ingredients"

func GetIngredients() []*domain.Ingredient {
	return GetAll[domain.Ingredient](INGREDIENTS_TABLE_NAME)
}

func AddIngredient(ingredient domain.Ingredient) {
	InsertOne(ingredient, INGREDIENTS_TABLE_NAME)
}
