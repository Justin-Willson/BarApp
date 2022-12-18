export class Ingredient {
    name: string;
    is_in_stock: boolean;
    notes: string;

    constructor(name: string, is_in_stock: boolean, notes: string) {
        this.name = name;
        this.is_in_stock = is_in_stock;
        this.notes = notes;
    }
}