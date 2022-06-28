package jsonflat

import (
	"fmt"
)

type flatString []string

var _ flattened = flatString(nil)

// string returns the string value of the flatString.
func (f flatString) string() string {
	return fmt.Sprintf("%s", f)
}

// length returns the length of the flatString.
func (f flatString) length() int {
	return len(f)
}

// contains returns true if the flatString contains the given interface{}.
func (f flatString) contains(i interface{}) bool {
	switch str := i.(type) {
	case string:
		for _, s := range f {
			if s == str {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// append adds the given interface{} to the flatString.
func (f flatString) append(i interface{}) flattened {
	switch str := i.(type) {
	case string:
		return append(f, str)
	}
	return nil
}
