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
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
}

type photo struct {
	AlbumID   int    `json:"albumId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Thumbnail string `json:"thumbnailUrl"`
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
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var albums []album
	json.Unmarshal(responseData, &albums)

	// for _, al := range albums {

	// 	stmt, _ := db.Prepare("insert into album(id, userId, title) values (?, ?, ?)")
	// 	_, _ = stmt.Exec(al.ID, al.UserID, al.Title)

	// }

	// for _, al := range albums {
	// 	url := "https://jsonplaceholder.typicode.com/photos?albumId=" + fmt.Sprintf("%d", al.ID)
	// 	fmt.Println(url)
	// 	response, err := http.Get(url)
	// 	if err != nil {
	// 		fmt.Print(err.Error())
	// 		os.Exit(1)
	// 	}
	// 	defer response.Body.Close()
	// 	var photos []photo
	// 	json.Unmarshal(responseData, &photos)
	// 	fmt.Println(photos)

	// }

	for i, _ := range albums {
		response, err = http.Get("https://jsonplaceholder.typicode.com/photos?albumId=" + fmt.Sprintf("%d", i+1))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		defer response.Body.Close()

		responseData, err = ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var photos []photo
		json.Unmarshal(responseData, &photos)
		fmt.Println(photos)

		for _, ph := range photos {

			stmt, _ := db.Prepare("insert into photo(id, albumId, title,url,thumbnailUrl) values (?, ?, ?,?,?)")
			_, _ = stmt.Exec(ph.ID, ph.AlbumID, ph.Title, ph.URL, ph.Thumbnail)
			defer stmt.Close()

		}

	}

}
