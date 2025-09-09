package api

import "encoding/json"

func JsonToBytes(loc any) ([]byte, error) {
	return json.Marshal(loc)
}

func BytesToJson[T any](data []byte) (T, error) {
	var loc T
	err := json.Unmarshal(data, &loc)
	return loc, err
}
