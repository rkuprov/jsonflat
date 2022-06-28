package jsonflat

import "fmt"

type flatBool []bool

var _ flattened = flatBool(nil)

// string returns the string value of the flatBool.
func (f flatBool) string() string {
	return fmt.Sprintf("%t", f)
}

// length returns the length of the flatBool.
func (f flatBool) length() int {
	return len(f)
}

// contains returns true if the flatBool contains the given interface{}.
func (f flatBool) contains(i interface{}) bool {
	switch v := i.(type) {
	case bool:
		for _, f := range f {
			if f == v {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// append adds the given interface{} to the flatBool.
func (f flatBool) append(i interface{}) flattened {
	switch v := i.(type) {
	case bool:
		return append(f, v)
	}
	return nil
}
