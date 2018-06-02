package main

import (
	"log"

	"github.com/FourSigma/bookmarks/internal/api"
)

func main() {
	log.Fatal(api.ListenAndServe("8080"))
}
