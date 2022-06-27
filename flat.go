package jsonflat

type flattened interface {
	string() string
	length() int
	contains(interface{}) bool
}

type JF struct {
	data map[string]flattened
}

func NewJSONFlat(b []byte) (JF, error) {

}
