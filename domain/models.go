package domain

type Ingredient struct {
    Name  string  `json:"name"`
    IsInStock bool  `json:"is_in_stock"`
    Notes  string `json:"notes"`
}

type IngredientQuantityPair struct {
	Ingredient Ingredient `json:"ingredient"`
	Quantity float64 `json:"quantity"`
	Units string `json:"units"`
}

type Cocktail struct {
	Name string `json:"name"`
	Ingredients []IngredientQuantityPair `json:"ingredients"`
	Notes string `json:"notes"`
}