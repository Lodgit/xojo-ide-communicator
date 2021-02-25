package xojo

import (
	"bytes"
	"fmt"
)

// xojoErrors defines Xojo error types to be checked.
var xojoErrors = [][]byte{
	// Xojo IDE
	[]byte("buildError"),
	[]byte("loadError"),
	[]byte("openErrors"),
	[]byte("scriptError"),
}

// CheckForErrorResponse just checks many possible errors during commands execution.
func CheckForErrorResponse(inJsonb []byte, inErr error) (jsonb []byte, err error) {
	err = inErr
	if err != nil {
		return
	}
	jsonb = inJsonb
	for _, errb := range xojoErrors {
		if bytes.Contains(inJsonb, errb) {
			err = fmt.Errorf(
				"Xojo IDE commands execution error occurred, please check the response output",
			)
			return
		}
	}
	return
}
