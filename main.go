package main

import (
	"log"
	"os"

	"ip-cli-go/app"
)

func main() {
	application := app.Generate()
	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
