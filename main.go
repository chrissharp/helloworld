package main

import (
	"fmt"
)

var (
	projectVersion = "dev"
	projectBuild   = "dev"
)

func main() {
	fmt.Printf("Hello wercker world %s, %s\n", projectVersion, projectBuild)
}
