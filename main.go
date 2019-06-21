package main

import (
	"os"

	"github.com/go-car-thing/src/reader"
	"github.com/go-car-thing/src/utils"
)

func main() {

	// Open the file for reading
	f, err := os.Open("resources/example_1.bag")
	utils.Check(err)
	defer f.Close()

	reader.ReadRosBag(f, true)
}
