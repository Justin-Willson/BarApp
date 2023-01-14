package app

import (
	"net/http"

	"github.com/Justin-Willson/BarApp/domain"

	"github.com/Justin-Willson/BarApp/dal"
	"github.com/gin-gonic/gin"
)

func GetIngredients(c *gin.Context) {
	var ingredients = dal.GetIngredients()
	c.IndentedJSON(http.StatusOK, ingredients)
}

func AddIngredient(c *gin.Context) {
	var newIngredient domain.Ingredient

	//TODO: add error handling
	if err := c.BindJSON(&newIngredient); err != nil {
		return
	}
	//TODO add object verification

	dal.AddIngredient(newIngredient)
	//TODO: maybe return object from db?
	c.IndentedJSON(http.StatusCreated, newIngredient)
}

func DeleteIngredientById(c *gin.Context) {
	id := c.Param("id")
	result := dal.DeleteIngredient(id)
	//TODO: add some sort of interpretation of result
	c.IndentedJSON(http.StatusOK, result)
}
