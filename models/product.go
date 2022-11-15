package main

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
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname) // connection string

	db, err = sql.Open("postgres", connString) // open the connection

	if err != nil {
		log.Fatal(err)
	}

}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float32
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES ($1, $2, $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Etkilenen kayıt sayısı: (%d)", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET title=$2, description=$3, price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Printf("Etkilenen kayıt sayısı: (%d)", rowsAffected)
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

	for _, value := range products {
		fmt.Printf("%d - %s, %s, %.2f\n", value.ID, value.Title, value.Description, value.Price)
	}
} // son scope
