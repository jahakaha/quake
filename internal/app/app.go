package app

import (
	"fmt"
	"quake-log-parser/internal/parser"
)

func Run() error {
	path := "./assets/game.log"
	// data, err := parser.ReadFile(path)
	// if err != nil {
	// 	return err
	// }
	fmt.Println(data)
	game, err := parser.ParseData()
	if err != nil {
		return err
	}
	str, err := FinalResponse(game)
	return nil
}
