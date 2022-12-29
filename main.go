package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gr8warrior/mongomock/router"
)

func main() {
	fmt.Print("Mongo DB")

	r := router.Router()
	fmt.Println("Mongo Mock Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Print("Server listening at port 4000")

}
