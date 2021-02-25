package xojotesting

import (
	"encoding/json"
)

// RuntimeError defines a JSON testing runtime error (Xojo Runtime Exception) object.
type RuntimeError struct {
	ErrorType    string   `json:"error_type"`
	ErrorNumber  int      `json:"error_number"`
	ErrorMessage string   `json:"error_message"`
	ErrorReason  string   `json:"error_reason"`
	ErrorStack   []string `json:"error_stack"`
}

// RuntimeErrorResult defines a JSON result which ccontains one runtime error (Xojo Runtime Exception) object.
type RuntimeErrorResult struct {
	RuntimeError RuntimeError `json:"runtime_error"`
}

// Test defines a JSON test object.
type Test struct {
	Name          string       `json:"name"`
	Duration      string       `json:"duration"`
	Passed        bool         `json:"passed"`
	Type          string       `json:"type"`
	FailedMessage string       `json:"failed_message"`
	RuntimeError  RuntimeError `json:"runtime_error"`
}

// TestGroup defines a JSON testing group object.
type TestGroup struct {
	Name                string `json:"name"`
	Duration            string `json:"duration"`
	Total               int    `json:"total"`
	PassedCount         int    `json:"passed_count"`
	FailuresCount       int    `json:"failures_count"`
	SkippedCount        int    `json:"skipped_count"`
	NotImplementedCount int    `json:"not_implemented_count"`
	Tests               []Test `json:"tests"`
}

// TestResult defines a JSON testing result object.
type TestResult struct {
	XojoVersion         string      `json:"xojo_version"`
	XojoUnitVersion     string      `json:"xojo_unit_version"`
	StartTime           string      `json:"start_time"`
	Duration            string      `json:"duration"`
	Total               string      `json:"total"`
	PassedCount         string      `json:"passed_count"`
	FailuresCount       string      `json:"failures_count"`
	SkippedCount        string      `json:"skipped_count"`
	NotImplementedCount string      `json:"not_implemented_count"`
	Groups              []TestGroup `json:"groups"`
}

// ParseTestResult parses JSON bytes of a XojoUnit result and return the corresponding `TestResult` object.
func ParseTestResult(jsonb []byte) (testResult TestResult, err error) {
	err = json.Unmarshal(jsonb, &testResult)
	return
}

// ParseRuntimeErrorResult parses JSON bytes of a XojoUnit result and return the corresponding `RuntimeErrorResult` object.
func ParseRuntimeErrorResult(jsonb []byte) (runtimeErrorResult RuntimeErrorResult, err error) {
	err = json.Unmarshal(jsonb, &runtimeErrorResult)
	return
}
