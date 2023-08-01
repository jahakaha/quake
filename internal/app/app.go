package app

import (
	"fmt"
	"io/ioutil"
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
	ressponse, err := FinalResponse(game)
	if err != nil {
		return err
	}
	fileName := "results.json"
	// Creating file
	err = ioutil.WriteFile(fileName, []byte(ressponse), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ressponse)
	fmt.Println("JSON data has been written to", fileName)
	return nil
}
