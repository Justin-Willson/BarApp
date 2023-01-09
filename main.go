package main

import (
	"fmt"
	"time"

	"github.com/Justin-Willson/BarApp/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"origin", "application/json"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/ingredients", app.GetIngredients)
	router.GET("/cocktails", app.GetCocktails)
	router.POST("/ingredients", app.AddIngredient)
	router.POST("/cocktails", app.AddCocktail)

	router.Run("localhost:8080")
	fmt.Println("running")
}
