// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package tcore

import "fmt"

// TErr returns a standardized message regarding error. Returns "", true when err == nil, and "some message", false
// when err is not nil
func TExpectErr(statement string, err error) (failureMessage string, ok bool) {
	if err != nil {
		return "", true
	}

	msg := fmt.Sprintf("an error was expected but did not occur during %s", statement)
	return msg, false
}
