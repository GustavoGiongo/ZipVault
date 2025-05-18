package main

import (
	"ZipVault/api"
	"fmt"
)

func main() {

	fmt.Println(Working())
	api.StartServer()
}

func Working() string {
	return "Working!"
}
