package app

import (
	"net/http"

	"github.com/Justin-Willson/BarApp/dal"
	"github.com/Justin-Willson/BarApp/domain"
	"github.com/gin-gonic/gin"
)

func GetCocktails(c *gin.Context) {
	var cocktails = dal.GetCocktails()
	c.IndentedJSON(http.StatusOK, cocktails)
}

func AddCocktail(c *gin.Context) {
	var newCocktail domain.Cocktail

	//TODO: add error handling
	if err := c.BindJSON(&newCocktail); err != nil {
		return
	}
	//TODO add object verification

	dal.AddCocktail(newCocktail)
	c.IndentedJSON(http.StatusCreated, newCocktail)
}
