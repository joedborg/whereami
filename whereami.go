package main

import (
	"fmt"
	"github.com/joedborg/whereami/lib"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := whereami.GetLocationData()
	errorCheck(err)
	fmt.Printf("%s, %s\n", data.City, data.Country)
}
