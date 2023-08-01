package parser

import (
	"bufio"
	"os"
	"quake-log-parser/internal/models"
	"strings"
)

// readFile func to get data from given path
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	events := []models.Event{}
	scanner := bufio.NewScanner(file)
	result := ""
	// scannig for a whole file
	for scanner.Scan() {
		line := scanner.Text()
		// Remove prefix whitespace
		result = strings.TrimLeft(line, " ")
		// Split data by 3 events
		event, err := SplitData(result)
		if err != nil {
			return "", err
		}

		events = append(events, *event)

	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return result, nil
}

func SplitData(data string) (*models.Event, error) {

	return event, nil
}
