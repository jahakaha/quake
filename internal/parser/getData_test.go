package parser

import (
	"quake-log-parser/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitData(t *testing.T) {
	tests := []struct {
		inputLine      string
		expectedEvent  *models.Event
		expectedError  bool
		expectedErrMsg string
	}{
		{
			inputLine: `  20:37       InitGame:`,
			expectedEvent: &models.Event{
				Timestamp: "20:37",
				EventType: "InitGame",
			},
			expectedError: false,
		},
		{
			inputLine: `  20:38       Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT`,
			expectedEvent: &models.Event{
				Timestamp:      "20:38",
				EventType:      "Kill",
				AdditionalData: "1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			},
			expectedError: false,
		},
		{
			inputLine:      `Invalid log entry`,
			expectedError:  true,
			expectedErrMsg: "invalid log entry format Invalid log entry",
		},
	}

	for _, test := range tests {
		event, err := SplitData(test.inputLine)

		if test.expectedError {
			assert.Error(t, err)
			assert.Nil(t, event)
			assert.Contains(t, err.Error(), test.expectedErrMsg)
		} else {
			assert.NoError(t, err)
			assert.NotNil(t, event)
			assert.Equal(t, test.expectedEvent, event)
		}
	}
}
