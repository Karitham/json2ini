package main

import (
	"bytes"
	"testing"

	"github.com/go-json-experiment/json"
)

func TestTransform(t *testing.T) {
	t.Parallel()

	var ini = map[string]map[string]string{
		"section1": {
			"key1": "value1",
			"key2": "value2",
		},
		"section2": {
			"key3": "value3",
			"key4": "value4",
		},
	}

	var expected = `[section1]
key1=value1
key2=value2
[section2]
key3=value3
key4=value4
`

	in := &bytes.Buffer{}
	json.MarshalWrite(in, ini)

	out := &bytes.Buffer{}
	err := Transform(in, out)
	if err != nil {
		t.Fatal(err)
	}

	if out.String() != expected {
		t.Errorf("Expected %q, got %q", expected, out.String())
	}
}
