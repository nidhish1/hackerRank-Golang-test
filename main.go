package main

//hackerRank-Golang-test
import (
	"fmt"
	"net/http"
	"os"

	"hackerRank-Golang-test/driver"

	ph "hackerRank-Golang-test/handler/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "root", dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(connection)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/", postRouter(pHandler))
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8005", r)
}

// A completely separate router for posts routes
func postRouter(pHandler *ph.Post) http.Handler {
	r := chi.NewRouter()

	r.Get("/search/{uid:[0-9]+}/{aid:[0-9]+}", pHandler.Search)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(""))
	})

	return r
}
