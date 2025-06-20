// logger_test.go - Tests for logger package

package logger

import (
	"errors"
	"testing"
)

// MockLogger implements Logger interface for testing
type MockLogger struct {
	Results      []string
	Errors       []string
	ShouldFail   bool
	CloseCalled  bool
}

func NewMockLogger() *MockLogger {
	return &MockLogger{
		Results:     make([]string, 0),
		Errors:      make([]string, 0),
		ShouldFail:  false,
		CloseCalled: false,
	}
}

func (ml *MockLogger) LogResult(deviceName string, result interface{}) error {
	if ml.ShouldFail {
		return errors.New("mock logging failure")
	}
	
	jsonStr, err := JSONFormatter(result)
	if err != nil {
		return err
	}
	
	ml.Results = append(ml.Results, jsonStr)
	return nil
}

func (ml *MockLogger) LogError(deviceName string, message string, err error) error {
	if ml.ShouldFail {
		return errors.New("mock error logging failure")
	}
	
	ml.Errors = append(ml.Errors, message)
	return nil
}

func (ml *MockLogger) Close() error {
	ml.CloseCalled = true
	if ml.ShouldFail {
		return errors.New("mock close failure")
	}
	return nil
}

func TestJSONFormatter(t *testing.T) {
	// Test with a simple map
	testData := map[string]string{
		"key": "value",
	}
	
	result, err := JSONFormatter(testData)
	if err != nil {
		t.Fatalf("JSONFormatter failed: %v", err)
	}
	
	expected := `{"key":"value"}`
	if result != expected {
		t.Errorf("JSONFormatter result mismatch: got %s, want %s", result, expected)
	}
	
	// Test with a struct
	type TestStruct struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}
	
	testStruct := TestStruct{
		Name:  "test",
		Value: 42,
	}
	
	result, err = JSONFormatter(testStruct)
	if err != nil {
		t.Fatalf("JSONFormatter failed with struct: %v", err)
	}
	
	expected = `{"name":"test","value":42}`
	if result != expected {
		t.Errorf("JSONFormatter result mismatch: got %s, want %s", result, expected)
	}
}
