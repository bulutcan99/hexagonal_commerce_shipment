package util

import (
	"encoding/json"
	"fmt"
)

func GenerateCacheKey(prefix string, params any) string {
	return fmt.Sprintf("%s:%v", prefix, params)
}

func GenerateCacheKeyParams(params ...any) string {
	var str string
	last := len(params) - 1
	for i, param := range params {
		str += fmt.Sprintf("%v", param)

		if i != last {
			str += ":"
		}
	}

	return str
}

func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

func Deserialize(data []byte, output any) error {
	return json.Unmarshal(data, output)
}
