package parser

import "fmt"

func ParseData() error {
	data, err := ReadFile(path)
	if err != nil {
		return err
	}
	game, err := SerializeData(data)
	if err != nil {
		return err
	}
	result, err := FinalResponse(game)
	if err != nil {
		fmt.Println(err)
	}
}

func SerializeData(events []Event) ([]Game, error) {
	return games, nil
}

func FinalResponse(games []Game) (string, error) {
	return response, nil
}
