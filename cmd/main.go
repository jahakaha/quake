package main

import (
	"fmt"
	"os"
	"quake-log-parser/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
