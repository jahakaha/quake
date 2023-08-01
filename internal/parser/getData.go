package parser

import (
	"bufio"
	"fmt"
	"os"
	"quake-log-parser/internal/models"
	"regexp"
	"strings"
)

// readFile func to get data from given path
func ReadFile(path string) ([]models.Event, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
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
			return nil, err
		}

		events = append(events, *event)

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return events, nil
}

// SplitData func to stplit line by events
func SplitData(line string) (*models.Event, error) {
	// regexp to split every line by events
	// split it by 3 events; timestamp, eventType, additionalData
	re := regexp.MustCompile(`^\s*(\d+:\d+)\s+([^:]+)(?::\s*(.*))?$`)

	// Find all the submatches in the data entry
	matches := re.FindStringSubmatch(line)

	if len(matches) < 3 {
		// fmt.Println(matches)
		return nil, fmt.Errorf("invalid log entry format")
	}

	event := &models.Event{
		Timestamp: matches[1],
		EventType: matches[2],
	}

	// Shutdown can be only 2 events
	// if it's not Shutdown type, we've got AdditionalData
	if len(matches) == 4 {
		event.AdditionalData = matches[3]
	}
	return event, nil
}
