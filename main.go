package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUser() []*User {
	db, err := sql.Open("mysql", "root:root/@tcp(db:3036)/@database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	results, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}
	var users []*User
	for results.Next() {
		var u User
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}
		users = append(users, &u)
	}
	return users
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
	fmt.Printf("hit the home page endpoint")
}
func userPage(w http.ResponseWriter, r *http.Request) {
	users := getUser()
	fmt.Fprintf(w, "Welcome to user page")
	fmt.Printf("hit the users page endpoint")
	json.NewEncoder(w).Encode(users)
}
func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", userPage)
	log.Fatal(http.ListenAndServe(":8000", nil))
	log.Println("welcome")
}
