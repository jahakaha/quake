package parser_test

import (
	"testing"

	"quake-log-parser/internal/models"
	"quake-log-parser/internal/parser"

	"github.com/stretchr/testify/assert"
)

func TestParseData(t *testing.T) {
	tests := []struct {
		filePath       string
		expectedOutput []models.Game
		expectedError  bool
	}{
		// Test case 1: Valid input with one game
		{
			filePath: "test_files/test_log_1.log",
			expectedOutput: []models.Game{
				{
					ID:         1,
					TotalKills: 1,
					Players: map[int]*models.Player{
						2: {
							ID:        2,
							Name:      "Isgalamido",
							KillCount: -1,
						},
					},
					KillMethod: map[string]int{
						"MOD_TRIGGER_HURT": 1,
					},
				},
			},
			expectedError: false,
		},
		// Test case 2: Valid input with multiple games
		{
			filePath: "test_files/test_log_2.log",
			expectedOutput: []models.Game{
				{
					ID:         1,
					TotalKills: 1,
					Players: map[int]*models.Player{
						2: {
							ID:        2,
							Name:      "Isgalamido",
							KillCount: -1,
						},
					},
					KillMethod: map[string]int{
						"MOD_TRIGGER_HURT": 1,
					},
				},
				{
					ID:         2,
					TotalKills: 1,
					Players: map[int]*models.Player{
						3: {
							ID:        3,
							Name:      "Player2",
							KillCount: -1,
						},
					},
					KillMethod: map[string]int{
						"MOD_TRIGGER_HURT": 1,
					},
				},
			},
			expectedError: false,
		},
		// Add more test cases as needed
	}

	for _, test := range tests {
		output, err := parser.ParseData(test.filePath)

		if test.expectedError {
			assert.Error(t, err)
			assert.Nil(t, output)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, output)
			assert.Equal(t, test.expectedOutput, output)
		}
	}
}
