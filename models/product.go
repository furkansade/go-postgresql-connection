package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // auxiliary tool taht will work with database/sql
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "productsdb"
)

var db *sql.DB // database/sql package pointer's

// initialization func
func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
}

type Product struct {
	ID                 int
	Title, Description string
	Price              float32
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Etkilenen kayıt sayısı: %d\n", rowsAffected)
}

func DeleteProduct(id int) {
	result, err := db.Exec("DELETE FROM cars WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Silinen kayıt sayısı: %d\n", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET title=$2, description=$3, price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Güncellenen kayıt sayısı: %d\n", rowsAffected)
}

func GetProducts() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close() // son scope oncesi bunu calistir

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd) // prd's append to products
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, pr := range products {
		fmt.Printf("%d - %s | %s | %.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
} // son scope

func GetProductByID(id int) {
	var product string
	var price float32
	err := db.QueryRow("SELECT title, price FROM products WHERE id=$1", id).Scan(&product, &price)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product -> %s | %.2f₺\n", product, price)
	}
}
