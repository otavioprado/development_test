package main

import (
	"floatme_development_test/pkg"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	groceryStore := pkg.Setup(file)
	timeResult := pkg.StartSimulation(groceryStore)
	print(fmt.Sprintf("\nTime result is => %d \n", timeResult))
}
