package main

import (
	"fmt"
	"github.com/joedborg/whereami/lib"
)

func main() {
	data := whereami.GetLocationData()
	fmt.Printf("%s, %s\n", data.City, data.Country)
}
