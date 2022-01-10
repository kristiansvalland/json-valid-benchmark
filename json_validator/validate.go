// Package json_validator contains methods for validating json
package json_validator

import "encoding/json"

func JsonValid(str string) bool {
	return json.Valid([]byte(str))
}

func JsonMarshalValid(str string) bool {
	var x interface{}
	err := json.Unmarshal([]byte(str), &x)
	return err == nil
}
