package main

import (
	"fmt"
	"github.com/joedborg/whereami/lib"
)

func main() {
	data := whereami.GetLocationData()
	fmt.Println(data.City)
}
