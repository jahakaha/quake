package app

import (
	"fmt"
	"os"
	"quake-log-parser/internal/parser"
)

func Run() error {
	path := "./assets/game.log"
	game, err := parser.ParseData(path)
	if err != nil {
		return fmt.Errorf("error ocured while parsing data %v", err)
	}
	response, err := FinalResponse(game)
	if err != nil {
		return err
	}
	fileName := "results.json"
	// Creating file
	err = os.WriteFile(fileName, []byte(response), 0644)
	if err != nil {
		return err
	}
	fmt.Println("JSON data has been written to", fileName)
	return nil
}
