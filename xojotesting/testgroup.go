package xojotesting

import "encoding/json"

// Test defines a JSON test object.
type Test struct {
	Name          string `json:"name"`
	Duration      string `json:"duration"`
	Passed        bool   `json:"passed"`
	FailedMessage string `json:"failed_message"`
}

// TestGroup defines a JSON testing group object.
type TestGroup struct {
	Name           string `json:"name"`
	Failures       int    `json:"failures"`
	Duration       string `json:"duration"`
	PassedCount    int    `json:"passed_count"`
	Total          int    `json:"total"`
	Skipped        int    `json:"skipped"`
	NotImplemented int    `json:"not_implemented"`
	Tests          []Test `json:"tests"`
}

// TestResult defines a JSON testing result object.
type TestResult struct {
	XojoVersion     string      `json:"xojo_version"`
	XojoUnitVersion string      `json:"xojo_unit_version"`
	StartTime       string      `json:"start_time"`
	Duration        string      `json:"duration"`
	Total           string      `json:"total"`
	PassedCount     string      `json:"passed_count"`
	Failures        string      `json:"failures"`
	NotImplemented  string      `json:"not_implemented"`
	Skipped         string      `json:"skipped"`
	Groups          []TestGroup `json:"groups"`
}

// ParseTestResult parses JSON bytes of a XojoUnit result and return the corresponding TestResult object.
func ParseTestResult(jsonb []byte) (testResult TestResult, err error) {
	err = json.Unmarshal(jsonb, &testResult)
	return
}
