// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package tcore

import "fmt"

// TErr returns a standardized message regarding error. Returns "", true when err == nil, and "some message", false
// when err is not nil
func TErr(statement string, err error) (failureMessage string, ok bool) {
	if err == nil {
		return "", true
	}

	msg := fmt.Sprintf("an error occurred during %s: %s", statement, err.Error())
	return msg, false
}
