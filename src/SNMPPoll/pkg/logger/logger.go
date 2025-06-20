// logger.go - Core logger interface for SNMP Poll

package logger

import (
	"encoding/json"
	"fmt"
)

// Logger defines the interface for all logging implementations
type Logger interface {
	// LogResult logs an SNMP polling result
	LogResult(deviceName string, result interface{}) error
	
	// LogError logs an error message
	LogError(deviceName string, message string, err error) error
	
	// Close closes any open resources
	Close() error
}

// JSONFormatter provides common JSON formatting functionality
func JSONFormatter(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data to JSON: %w", err)
	}
	return string(bytes), nil
}
