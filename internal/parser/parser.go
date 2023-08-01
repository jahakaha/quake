package parser

import (
	"fmt"
	"quake-log-parser/internal/models"
)

// Parse data by final result
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
	return nil
}

// SerializeData func to marshall events by Game struct
func SerializeData(events []models.Event) ([]models.Game, error) {
	fmt.Println("serializeData")
	currentGameID := 0
	games := []models.Game{}

	for _, event := range events {
		// fmt.Println(event.EventType)
		if event.EventType == models.InitGame {
			// id
			currentGameID++
			game := &models.Game{
				KillMethod: make(map[string]int),
				Players:    map[int]*models.Player{},
			}
			game.ID = currentGameID + 1
			games = append(games, *game)
		} else if event.EventType == models.ClientConnect {

		} else if event.EventType == models.Kill {

		} else if event.EventType == models.Divider {
			continue
		}
	}
	return games, nil
}

func FinalResponse(games []models.Game) (string, error) {
	response := ""
	return response, nil
}
