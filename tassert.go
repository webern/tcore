// Copyright (c) 2019 by Matthew James Briggs, https://github.com/webern

package tcore

import (
	"fmt"
	"math"
	"strconv"
)

// TAssertString returns a standardized message regarding error. ("", true) when got == want, ("message", false) otherwise
func TAssertString(statement, got, want string) (failureMessage string, ok bool) {
	if got != want {
		return tmsg(statement, got, want), false
	}

	return "", true
}

// TAssertInt returns a standardized message regarding error. ("", true) when got == want, ("message", false) otherwise
func TAssertInt(statement string, got, want int) (failureMessage string, ok bool) {
	if got != want {
		return tmsg(statement, strconv.Itoa(got), strconv.Itoa(want)), false
	}

	return "", true
}

// TAssertFloat returns a standardized message regarding error. ("", true) when got == want, ("message", false) otherwise
func TAssertFloat(statement string, got, want, epsilon float64) (failureMessage string, ok bool) {
	if math.Abs(got-want) > epsilon {
		return tmsg(statement, fmt.Sprintf("%.10f", got), fmt.Sprintf("%.10f", want)), false
	}

	return "", true
}

// TAssertBool returns a standardized message regarding error. ("", true) when got == want, ("message", false) otherwise
func TAssertBool(statement string, got, want bool) (failureMessage string, ok bool) {
	if got != want {
		return tmsg(statement, fmt.Sprintf("%t", got), fmt.Sprintf("%t", want)), false
	}

	return "", true
}

// TAssertDeepEqual returns a standardized message regarding error. ("", true) when got == want, ("message", false) othw
func TAssertDeep(statement string, got, want interface{}) (failureMessage string, ok bool) {
	if got != want {
		return tmsg(statement, fmt.Sprintf("%v", got), fmt.Sprintf("%v", want)), false
	}

	return "", true
}

// tmsg formats a standard failure message
func tmsg(statement, got, want string) string {
	return fmt.Sprintf("%s = '%s', want '%s'", statement, got, want)
}
