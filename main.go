package main

import "gopostgre/models"

func main() {

	product := models.Product{
		Title:       "Go ile WEB Programlama",
		Description: "Yazar: Furkan Sade Uçkun",
		Price:       76.99,
	}
	models.InsertProduct(product)

	/* product := models.Product{
		ID:          2,
		Title:       "System Interview",
		Description: "The system also is big!",
		Price:       134.99,
	}
	models.UpdateProduct(product) */

	// models.GetProductByID(2)
	// models.GetProducts()
}
