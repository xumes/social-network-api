package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Running the API")

	appRouter := router.Generate()

	log.Fatal(http.ListenAndServe(":5080", appRouter))
}
