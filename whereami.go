package main

import (
	"fmt"
	"github.com/joedborg/whereami/lib"
)

func main() {
	data, _ := whereami.GetLocationData()
	fmt.Printf("%s, %s\n", data.City, data.Country)
}
