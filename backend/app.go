package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"database/sql"

	_ "github.com/lib/pq"
)

type Account struct {
	Username string
	Password string
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	env := os.Getenv("APP_ENV")
	host := ""
	if env == "docker" {
		host = "db"
	} else {
		host = "localhost"
	}

	connStr := fmt.Sprintf("host=%s user=docker password=docker dbname=postgres sslmode=disable", host)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		w.Write([]byte(fmt.Sprintf("A Bunch of farts happened: %v", err)))
		return
	}

	rows, err := db.Query("SELECT username FROM account")
	if err != nil {
		fmt.Printf("pq error: %v\n", err)
		w.Write([]byte(fmt.Sprintf("pq error: %v\n", err)))
		return
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
