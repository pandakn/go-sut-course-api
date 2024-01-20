package utils

import "encoding/json"

// Serialize marshals the input data into an array of bytes
func Serialize(data any) ([]byte, error) {
	return json.Marshal(data)
}

// Deserialize unmarshals the input data into the output interface
func Deserialize(data []byte, output any) error {
	return json.Unmarshal(data, output)
}
