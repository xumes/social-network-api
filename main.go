package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Printf("Running the API at port: %d\n", config.Port)

	fmt.Printf("DB Connection string: %s\n", config.DatabaseConnectionString)

	appRouter := router.Generate()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), appRouter))
}
