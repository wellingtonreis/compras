package main

import (
	"compras/internal/app/handler"
	"log"
	"net/http"
)

func main() {
	router := handler.SetupRoutes()

	log.Println("Server starting on port 3000...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
