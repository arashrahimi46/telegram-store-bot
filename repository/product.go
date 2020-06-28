package repository

import (
	"fmt"
	"log"
)

type Product struct {
	Id    int
	Title string
	Price int
}

func GetProducts(query string) []Product {
	var id int
	var title string
	var price int
	fmt.Println(query , fmt.Sprintf("%%%s%%", query))
	db := NewMysqlConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products WHERE title LIKE ?" ,fmt.Sprintf("%%%s%%", query))
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		err := rows.Scan(&id, &title, &price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, Product{Id: id, Title: title, Price: price})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return products
}

func GetProductByTitle(name string) Product {
	var id int
	var title string
	var price int
	db := NewMysqlConnection()
	defer db.Close()
	err := db.QueryRow("select * from products where title = ?", name).Scan(&id , &title , &price)
	if err != nil {
		fmt.Println(err)
	}
	return Product{Id: id , Title: title , Price: price}
}
