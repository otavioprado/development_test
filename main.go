package main

import (
	"floatme_development_test/pkg"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	pkg.Check(err)

	groceryStore, err := pkg.Setup(file)
	pkg.Check(err)

	timeResult, err := pkg.StartSimulation(groceryStore)
	pkg.Check(err)

	print(fmt.Sprintf("\nTime result is => %d \n", timeResult))
}
