package parser

import (
	"testing"

	"quake-log-parser/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {

	actualGames := []models.Game{
		{
			ID:         1,
			TotalKills: 2,
			Players: map[int]*models.Player{
				0: {ID: 0, Name: "<world>", KillCount: 0},
				1: {ID: 1, Name: "player1", KillCount: 1},
				2: {ID: 2, Name: "player2", KillCount: 1},
			},
			KillMethod: map[string]int{
				models.MOD_TRIGGER_HURT: 2,
			},
		},
	}

	expectedGames := []models.Game{
		{
			ID:         1,
			TotalKills: 2,
			Players: map[int]*models.Player{
				0: {ID: 0, Name: "<world>", KillCount: 0},
				1: {ID: 1, Name: "player1", KillCount: 1},
				2: {ID: 2, Name: "player2", KillCount: 1},
			},
			KillMethod: map[string]int{
				models.MOD_TRIGGER_HURT: 2,
			},
		},
	}

	assert.Equal(t, expectedGames, actualGames, "Parsed games do not match expected games")
}

func TestAddKill(t *testing.T) {
	game := &models.Game{
		Players: map[int]*models.Player{
			1: {ID: 1, Name: "player1", KillCount: 0},
			2: {ID: 2, Name: "player2", KillCount: 0},
		},
		KillMethod: make(map[string]int),
	}

	killData := "1 2 MOD_RAILGUN"
	_, err := addKill(killData, game)
	assert.NoError(t, err, "Unexpected error")

	expectedGame := &models.Game{
		Players: map[int]*models.Player{
			1: {ID: 1, Name: "player1", KillCount: 1},
			2: {ID: 2, Name: "player2", KillCount: -1},
		},
		KillMethod: map[string]int{
			models.MOD_RAILGUN: 1,
		},
	}

	assert.Equal(t, expectedGame, game, "Updated game data do not match expected data")
}

func TestGetPlayerName(t *testing.T) {
	name := `n\player1\t\`
	expectedName := "player1"
	result, err := getPlayerName(name)
	assert.NoError(t, err, "Unexpected error")

	assert.Equal(t, expectedName, result, "Parsed player name does not match expected name")
}

func TestGetPlayerNameInvalid(t *testing.T) {
	name := `invalid_data`
	_, err := getPlayerName(name)
	assert.Error(t, err, "Expected an error for invalid player name")
}
