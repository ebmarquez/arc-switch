// syslog.go - Syslog implementation with JSON formatting

package logger

import (
	"encoding/json"
	"fmt"
	"log/syslog"
	"time"
)

// SyslogLogger implements Logger interface for syslog
type SyslogLogger struct {
	writer *syslog.Writer
	tag    string
}

// SyslogMessage represents the structure of the JSON message sent to syslog
type SyslogMessage struct {
	Timestamp  string      `json:"timestamp"`
	DeviceName string      `json:"device_name"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Message    string      `json:"message,omitempty"`
}

// NewSyslogLogger creates a new syslog logger with the specified tag
func NewSyslogLogger(tag string) (*SyslogLogger, error) {
	writer, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, tag)
	if err != nil {
		return nil, fmt.Errorf("failed to create syslog writer: %w", err)
	}
	
	return &SyslogLogger{
		writer: writer,
		tag:    tag,
	}, nil
}

// LogResult logs an SNMP polling result to syslog in JSON format
func (sl *SyslogLogger) LogResult(deviceName string, result interface{}) error {
	msg := SyslogMessage{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		DeviceName: deviceName,
		Data:       result,
	}
	
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal result to JSON: %w", err)
	}
	
	return sl.writer.Info(string(jsonData))
}

// LogError logs an error message to syslog in JSON format
func (sl *SyslogLogger) LogError(deviceName string, message string, err error) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	
	msg := SyslogMessage{
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		DeviceName: deviceName,
		Message:    message,
		Error:      errMsg,
	}
	
	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal error to JSON: %w", err)
	}
	
	return sl.writer.Err(string(jsonData))
}

// Close closes the syslog writer
func (sl *SyslogLogger) Close() error {
	if sl.writer != nil {
		return sl.writer.Close()
	}
	return nil
}
