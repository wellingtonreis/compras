package main

import (
	"log"
	"net/http"

	"github.com/wellingtonreis/compras/internal/app/handler"
)

func main() {
	router := handler.SetupRoutes()

	log.Println("Server starting on port 3000...")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
