package main

import (
	"cli-net-tool/app"
	"log"
	"os"
)

func main() {
	app := app.Generate()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}