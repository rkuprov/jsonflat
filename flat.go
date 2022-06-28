package jsonflat

import (
	"encoding/json"
	"fmt"
)

type flattened interface {
	string() string
	length() int
	contains(interface{}) bool
	append(interface{}) flattened
}

type JF struct {
	data map[string]flattened
}

func NewJSONFlat(b []byte) (*JF, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	fmt.Println("map:", m)
	if err != nil || m == nil {
		return nil, err
	}
	return flattenJSON(m), nil
}

// GetString returns the value of the given key in the JF object.
func (jf *JF) GetString(key string) string {
	if v, ok := jf.data[key]; ok {
		return v.string()
	}
	return ""
}

// Contains returns true if the JF object contains the given interface{}.
func (jf *JF) Contains(str string) bool {
	if _, ok := jf.data[str]; ok {
		return true
	}

	return false
}

func flattenJSON(m map[string]interface{}) *JF {
	jf := JF{data: make(map[string]flattened)}
	for k, v := range m {
		switch val := v.(type) {
		case string:
			if _, ok := jf.data[k]; !ok {
				jf.data[k] = flatString{val}
				continue
			}
			jf.data[k] = jf.data[k].append(val)
		case float64:
			if _, ok := jf.data[k]; !ok {
				jf.data[k] = flatFloat{val}
				continue
			}
			jf.data[k] = jf.data[k].append(val)
		case bool:
			if _, ok := jf.data[k]; !ok {
				jf.data[k] = flatBool{val}
				continue
			}
			jf.data[k] = jf.data[k].append(val)
		case map[string]interface{}:
			for k, val := range flattenJSON(val).data {
				if _, ok := jf.data[k]; !ok {
					jf.data[k] = val
					continue
				}
				jf.data[k].append(val)
			}
		case []interface{}:
			if _, ok := jf.data[k]; !ok {
				jf.data[k] = flattenSlice(val)
				continue
			}
			jf.data[k] = jf.data[k].append(flattenSlice(v.([]interface{})))
		default:
			return nil
		}
	}

	return &jf
}

func flattenSlice(s []interface{}) flattened {
	if len(s) == 0 {
		return nil
	}

	switch s[0].(type) {
	case string:
		ret := make(flatString, len(s))
		for _, v := range s {
			ret = append(ret, v.(string))
		}
		return ret
	case float64:
		ret := make(flatFloat, len(s))
		for _, v := range s {
			ret = append(ret, v.(float64))
		}
		return ret
	case bool:
		ret := make(flatBool, len(s))
		for _, v := range s {
			ret = append(ret, v.(bool))
		}
		return ret
	default:
		ret := make(flatString, len(s))
		for _, v := range s {
			ret.append(fmt.Sprintf("%v", v))
		}
	}

	return nil
}
