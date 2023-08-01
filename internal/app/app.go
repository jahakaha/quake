package app

import (
	"fmt"
	"quake-log-parser/internal/parser"
)

func Run() error {
	path := "./assets/game.log"
	fmt.Println(path)
	// data, err := parser.ReadFile(path)
	// if err != nil {
	// 	return err
	// }
	fmt.Println(path)
	game, err := parser.ParseData(path)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error ocured while parsing data %v", err)
	}
	fmt.Println(game)
	// str, err := FinalResponse(game)
	return nil
}
