package main

import (
	"fmt"

	"github.com/Justin-Willson/BarApp/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ingredients", app.GetIngredients)
	router.GET("/cocktails", app.GetCocktails)
	router.POST("/ingredients", app.AddIngredient)
	router.POST("/cocktails", app.AddCocktail)

	fmt.Println("running")
	router.Run("localhost:8080")
}
