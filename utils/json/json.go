package json

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
)

// Json well it's a json.
type Json = map[string]any

// Parse decodes the given json string into its corresponding type.
func Parse[T any](s string) (T, error) {
	buf := strings.NewReader(s)
	return ParseFromReader[T](buf)
}

// ParseFromReader decodes the given json string from an io.Reader into its corresponding type.
func ParseFromReader[T any](reader io.Reader) (T, error) {
	dec := json.NewDecoder(reader)

	var target T

	err := dec.Decode(&target)
	if err != nil {
		return target, err
	}

	return target, nil
}

// Stringify encodes the given object into json string.
func Stringify[T any](obj T) (string, error) {
	out := bytes.NewBuffer([]byte{})
	err := StringifyToWriter(out, obj)
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

// Stringify encodes the given object into a json string to the given io.Writer.
func StringifyToWriter[T any](writer io.Writer, obj T) error {
	enc := json.NewEncoder(writer)

	err := enc.Encode(obj)
	if err != nil {
		return err
	}

	return nil
}

func Jsonify[T any](obj T) (Json, error) {
	jsonStr, err := Stringify(obj)
	if err != nil {
		return nil, err
	}
	mappedJson, err := Parse[Json](jsonStr)
	if err != nil {
		return nil, err
	}
	return mappedJson, nil
}
