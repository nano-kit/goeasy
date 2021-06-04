package json

import "encoding/json"

// Stringify v to the JSON string representation
func Stringify(v interface{}) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
