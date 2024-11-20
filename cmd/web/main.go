package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vzhovtan/gweb/internal/controller"
	"github.com/vzhovtan/gweb/internal/db"
	"log"
	"net/http"
)

func main() {
	sq, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer sq.Close()

	if err := db.CreateTable(sq); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Path == "/" {
			controller.ShowIndex(writer, request)
		} else {
			controller.Proxy(sq)(writer, request)
		}
	})

	http.HandleFunc("/shorten", controller.Shorten(sq))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
