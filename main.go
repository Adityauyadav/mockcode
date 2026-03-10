package main

import (
	"fmt"
	"log"
	"net/http"

	"mockcode/db"
	"mockcode/handler"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Println("MockCode server starting on port 8080...")
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/submit", handler.Submit)
	http.ListenAndServe(":8080", nil)
}
