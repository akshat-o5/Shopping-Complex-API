package main

import (
	"log"
	"net/http"

	"github.com/akshat-o5/shopping-complex-api/models"
	"github.com/akshat-o5/shopping-complex-api/routes"
)

func main() {
	models.ConnectDB()

	r := routes.SetUpRoutes()

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
