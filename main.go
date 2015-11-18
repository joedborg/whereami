package main

import (
	"fmt"
	"github.com/joedborg/whereami/lib"
	"log"
)

func main() {
	data, err := whereami.GetLocationData()
	if err != nil {
		log.Fatalf("Error getting location: %v", err)
	}
	fmt.Printf("%s, %s\n", data.City, data.Country)
}
