package main

import (
	"github.com/joer-projects/fingo/setup"
)

func main() {
	err := setup.InitPostgres()
	if err != nil {
		panic(err)
	}
}
