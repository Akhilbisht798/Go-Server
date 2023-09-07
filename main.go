package main

import (
	"booking-app/pkg/handlers"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/user", handlers.CreateUserHandler)
	fmt.Println("Server is running on :3000")
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from go web server")
}