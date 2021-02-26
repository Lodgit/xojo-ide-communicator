package xojotesting

import (
	"bytes"
	"fmt"
)

// xojoUnitGlobalError defines a XojoUnit global scope error.
// Those exception are captured and displayed directly rather than those per test.
// Test result ones have a "error_scope_test_result" scope error instead
var xojoUnitGlobalError = []byte("error_scope_global")

// CheckForGlobalErrorResponse just checks a possible "global scope error" during commands execution.
func CheckForGlobalErrorResponse(inJsonb []byte, inErr error) (jsonb []byte, err error) {
	err = inErr
	if err != nil {
		return
	}
	jsonb = inJsonb
	if bytes.Contains(inJsonb, xojoUnitGlobalError) {
		err = fmt.Errorf(
			"XojoUnit execution error occurred, please check the response output",
		)
		return
	}
	return
}
