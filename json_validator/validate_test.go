package json_validator_test

import (
	"testing"

	"github.com/kristiansvalland/json-valid-benchmark/json_validator"
)

const JSON_STRING string = `
{
	"someval": 123,
	"object": {
		"a": "A",
		"b": "B",
		"c": 3
	},
	"longString": "fadklfjasfadshfkjashdfasdfjasdhfjkasfjalsfj"
}
`

func BenchmarkJsonValid(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if !json_validator.JsonValid(JSON_STRING) {
			panic("invalid json")
		}
	}
}

func BenchmarkJsonUnmarshalValid(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		if !json_validator.JsonMarshalValid(JSON_STRING) {
			panic("invalid json")
		}
	}
}
