package domain

type DatabaseObject interface {
	TableName() string
}

type Ingredient struct {
	Name      string `json:"name"`
	IsInStock bool   `json:"is_in_stock"`
	Notes     string `json:"notes"`
}

func (_ Ingredient) TableName() string {
	return "ingredients"
}

type RecipeLine struct {
	Ingredient Ingredient `json:"ingredient"`
	Quantity   float64    `json:"quantity"`
	Units      string     `json:"units"`
}

type Cocktail struct {
	Name        string       `json:"name"`
	Ingredients []RecipeLine `json:"ingredients"`
	Notes       string       `json:"notes"`
}

func (_ Cocktail) TableName() string {
	return "cocktails"
}
