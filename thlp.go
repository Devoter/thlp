package thlp

import (
	"bytes"
	"reflect"
	"regexp"
)

// Asserter is an interface that provides necessary methods for assert functions.
type Asserter interface {
	// Fatalf is equivalent to Logf followed by FailNow.
	Fatalf(format string, arguments ...interface{})
}

// Equal calls t.Fatalf if expected and got arguments are not equal.
func Equal(t Asserter, expected interface{}, got interface{}, format string) {
	if expected != got {
		t.Fatalf(format+"\n", expected, got)
	}
}

// DeepEqual calls t.Fatalf if expected and got arguments are not equal deeply.
func DeepEqual(t Asserter, expected interface{}, got interface{}, format string) {
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf(format+"\n", expected, got)
	}
}

// Bytes calls t.Fatalf if expected and got bytes slices are not equal.
func Bytes(t Asserter, expected []byte, got []byte, format string) {
	if !bytes.Equal(expected, got) {
		t.Fatalf(format+"\n", expected, got)
	}
}

// Ok calls t.Fatalf if ok is not true.
func Ok(t Asserter, ok bool, format string) {
	if !ok {
		t.Fatalf(format + "\n")
	}
}

// Cmp calls t.Fatalf if compare function returns false.
func Cmp(t Asserter, cmp func(e interface{}, g interface{}) bool, expected interface{}, got interface{}, format string) {
	if !cmp(expected, got) {
		t.Fatalf(format+"\n", expected, got)
	}
}

// Err calls t.Fatalf if the pattern (regexp) is matched into error.
// If pattern is `""` and error is nil the assertion is true.
func Err(t Asserter, pattern string, err error, format string) {
	if !compareError(pattern, err) {
		t.Fatalf(format+"\n", pattern, err)
	}
}

func compareError(pattern string, err error) bool {
	if err == nil {
		return pattern == ""
	}

	if pattern == "" {
		return false
	}

	matched, _ := regexp.MatchString(pattern, err.Error())

	return matched
}
