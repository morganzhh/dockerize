package main

import (
	"encoding/json"
	y "github.com/ghodss/yaml"
)

type yamlDecoder struct {
}

func (i *yamlDecoder) Decode(data []byte, v interface{}) error {
	bb, err := y.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, &v)
}