package main

import (
	"log"
	"rest-api/books"
	"rest-api/server"
)

func main() {
	lib := books.NewLibrary()
	hand := server.NewHTTPHandler(&lib)
	server := server.NewServer(hand)

	if err := server.StartServer(); err != nil {
		log.Fatalln(err)
	}
}
