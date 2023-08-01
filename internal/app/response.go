package app

import (
	"encoding/json"
	"fmt"
	"quake-log-parser/internal/models"
)

// FinalResponse func takes a slice of Game
// and returns a formatted string containing the response for each game
func FinalResponse(games []models.Game) (string, error) {
	response := ""
	// Loop through each game in the input slice
	for _, game := range games {
		// Format the game name with its ID (e.g., "game-1")
		gameName := fmt.Sprintf("game-%d", game.ID)
		players := []string{}
		kills := make(map[string]int)
		// Loop through each player in the current game
		for _, player := range game.Players {
			// Add the player's name to the players slice
			players = append(players, player.Name)

			// Store the player's kill count in the kills map with the player's name as the key
			kills[player.Name] = player.KillCount
		}

		// Create a GameResponse struct to hold the game data for the response
		gameResponse := &models.GameReponse{
			TotalKills: game.TotalKills,
			Players:    players,
			Kills:      kills,
			KillMeans:  game.KillMethod,
		}

		// Marshal the GameResponse struct into a formatted JSON byte slice
		gameResponseBytes, err := json.MarshalIndent(gameResponse, "", "\t")
		if err != nil {
			return "", err
		}
		// Append the formatted response for the current game to the overall response string
		response += fmt.Sprintf("%s: %s\n", gameName, string(gameResponseBytes))

	}
	return response, nil
}
