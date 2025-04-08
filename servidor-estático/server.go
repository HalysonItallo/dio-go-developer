package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./servidor-estático/web"))
	http.Handle("/", fs)

	log.Print("Listen on 3000 ...", nil)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
