package main

import (
	"fmt"
	"net/http"

    g "github.com/gin-gonic/gin"
	"github.com/Justin-Willson/BarApp/domain"
)

var ingredients = []domain.Ingredient{
	{ Name: "Gin", IsInStock: true, Notes: "Better than vodka" },
	{ Name: "Dry Vermouth", IsInStock: true, Notes: "Great for martinis" },
}
var gin = ingredients[0]
var vermouth = ingredients[1]

var ginPair = domain.IngredientQuantityPair{ Ingredient: gin, Quantity: 3.0, Units: "ounces"}
var vermouthPair = domain.IngredientQuantityPair{ Ingredient: vermouth, Quantity: 0.75, Units: "ounces"}

var martini = domain.Cocktail{Ingredients: []domain.IngredientQuantityPair{ginPair, vermouthPair},
							  Name: "Martini",
  							  Notes: "The best cocktial",
}

func getIngredients(c *g.Context) {
    c.IndentedJSON(http.StatusOK, ingredients)
}

func main() {
	router := g.Default()
    router.GET("/ingredients", getIngredients)

	fmt.Println("running")
    router.Run("localhost:8080")
}