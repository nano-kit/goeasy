package json

import "encoding/json"

func Stringify(v interface{}) string {
	buf, _ := json.Marshal(v)
	return string(buf)
}
