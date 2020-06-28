package repository

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"


func NewMysqlConnection() *sql.DB {
	db, err := sql.Open("mysql", "arashrahimi46:fazilatschool1@/test_store")
	if err != nil {
		fmt.Println(err)
	}
	return db
}