package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "g97.nidhish@gmail.com"
	password = "g97.nidhish@gmail.com"
	hostname = "143.110.190.177:3306"
	dbname   = "g97.nidhish@gmail.com"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

type album struct {
	UserID int `json:"userId"`
	ID     int `json:"id"`
	Title  int `json:"title"`
}

func main() {

	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return
	}
	defer db.Close()

	db.Ping()

	response, err := http.Get("https://jsonplaceholder.typicode.com/albums")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var albums []album
	json.Unmarshal(responseData, &albums)
	fmt.Printf("Birds : %+v", albums)

}
