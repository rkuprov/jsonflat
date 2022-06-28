package jsonflat

import "fmt"

type flatFloat []float64

var _ flattened = flatFloat(nil)

// string returns the string value of the flatFloat.
func (f flatFloat) string() string {
	return fmt.Sprintf("%f", f)
}

// length returns the length of the flatFloat.
func (f flatFloat) length() int {
	return len(f)
}

// contains returns true if the flatFloat contains the given interface{}.
func (f flatFloat) contains(i interface{}) bool {
	switch v := i.(type) {
	case float64:
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

// append adds the given interface{} to the flatFloat.
func (f flatFloat) append(i interface{}) flattened {
	switch v := i.(type) {
	case float64:
		return append(f, v)
	}
	return nil
}
