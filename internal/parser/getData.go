package parser

import (
	"bufio"
	"fmt"
	"os"
	"quake-log-parser/internal/models"
	"regexp"
	"strings"
)

// ReadFile reads the data from the given file path and returns a slice of events and an error.
func ReadFile(path string) ([]models.Event, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("in")
		return nil, err
	}
	defer file.Close()
	events := []models.Event{}
	scanner := bufio.NewScanner(file)
	result := ""
	// Scanning the entire file
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

// SplitData splits the input line by events and returns a pointer to the Event object and an error.
func SplitData(line string) (*models.Event, error) {
	// Regexp to split every line by events
	// Split it by 3 events: timestamp, eventType, additionalData
	fmt.Println("line", line)
	re := regexp.MustCompile(`^\s*(\d+:\d+)\s+([^:]+)(?::\s*(.*))?$`)

	// Find all the submatches in the entry line
	matches := re.FindStringSubmatch(line)

	if len(matches) < 3 {
		// fmt.Println(matches)
		return nil, fmt.Errorf("invalid log entry format %s", line)
	}

	event := &models.Event{
		Timestamp: matches[1],
		EventType: matches[2],
	}

	// Shutdown can only be 2 events
	// If it's not a Shutdown type, we have AdditionalData
	if len(matches) == 4 {
		event.AdditionalData = matches[3]
	}
	return event, nil
}
