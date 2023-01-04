import axios from "axios";

export class Ingredient {
    name: string;
    is_in_stock: boolean;
    notes: string;
    private static readonly INGREDIENTS_ENDPOINT: string = "http://localhost:8080/ingredients"

    constructor(name: string, is_in_stock: boolean, notes: string) {
        this.name = name;
        this.is_in_stock = is_in_stock;
        this.notes = notes;
    }

    static async getIngredients(): Promise<Ingredient[]>  {
        try {
            const ingredients = axios.get(Ingredient.INGREDIENTS_ENDPOINT)
                                    .then((response) => response.data)
                                    .then((data) => data as Ingredient[])
            
            return ingredients   
            
          } catch (error) {
            console.log(error);
          }
          return []
    }
}