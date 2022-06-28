package jsonflat

import (
	"fmt"
	"testing"
)

var john = []byte(`{
						"person": {
							"name": "John",
							"age": "30",
							"city": "New York"
							},
						"id": "1"
				}`)

func TestNewJSONFlat(t *testing.T) {
	flat, err := NewJSONFlat(john)
	if err != nil {
		t.Error(err)
	}
	if flat == nil {
		t.Error("NewJSONFlat returned nil")
	}
	fmt.Println(flat.Contains("name"))
	fmt.Println(flat.Contains("age"))
	fmt.Println(flat.Contains("city"))
	fmt.Println(flat.GetString("id"))
	fmt.Println(flat.Contains("id"))
	fmt.Println(flat.data)
}
