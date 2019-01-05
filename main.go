package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"database/sql"

	_ "github.com/lib/pq"
)

type Account struct {
	Username string
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	connStr := "host=db user=docker password=docker dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(fmt.Sprintf("A Bunch of farts happened: %v", err)))
		return
	}

	rows, err := db.Query("SELECT username FROM account")
	if err != nil {
		fmt.Printf("pq error: %v\n", err)
	}
	i := 0
	for rows.Next() {
		i++
		fmt.Println(i)
		var account Account
		rows.Scan(&account.Username)
		fmt.Printf("row: %s\n", account.Username)
	}

	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
