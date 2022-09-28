package main

import (
	"os"
	"testProject/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
