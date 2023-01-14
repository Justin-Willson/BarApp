package dal

import (
	"github.com/Justin-Willson/BarApp/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

const INGREDIENTS_TABLE_NAME = "ingredients"

func GetIngredients() []*domain.Ingredient {
	return GetAll[domain.Ingredient](INGREDIENTS_TABLE_NAME)
}

func AddIngredient(ingredient domain.Ingredient) {
	InsertOne(ingredient, INGREDIENTS_TABLE_NAME)
}

func DeleteIngredient(id string) *mongo.DeleteResult {
	return DeleteById(INGREDIENTS_TABLE_NAME, id)
}
