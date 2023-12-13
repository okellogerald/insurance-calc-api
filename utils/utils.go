package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func StructToMap(v any) (map[string]interface{}, error) {
	var result map[string]interface{}
	inrec, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(inrec, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetEnvVariable(key string) (string, error) {
	v := os.Getenv(key)
	if len(v) == 0 {
		return "", errors.New("we couldn't find the signing key")
	}

	return v, nil
}

