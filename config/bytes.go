package config

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type (
	DecodeFunc func(source []byte, target interface{}) error
	EncodeFunc func(i interface{}) ([]byte, error)
)

func DecodeJSON(bs []byte, i interface{}) error {
	return json.Unmarshal(bs, i)
}

func EncodeJSON(i interface{}) ([]byte, error) {
	return json.Marshal(i)
}

func EncodeJSONIndent(prefix, indent string) EncodeFunc {
	return func(i interface{}) ([]byte, error) {
		return json.MarshalIndent(i, prefix, indent)
	}
}

func DecodeXML(bs []byte, i interface{}) error {
	return xml.Unmarshal(bs, i)
}

func EncodeXML(i interface{}) ([]byte, error) {
	return xml.Marshal(i)
}

func EncodeXMLIndent(prefix, indent string) EncodeFunc {
	return func(i interface{}) ([]byte, error) {
		return xml.MarshalIndent(i, prefix, indent)
	}
}

func DecodeYAML(bs []byte, i interface{}) error {
	return yaml.Unmarshal(bs, i)
}

func EncodeYAML(i interface{}) ([]byte, error) {
	return yaml.Marshal(i)
}

func DecodeTOML(bs []byte, i interface{}) error {
	return toml.Unmarshal(bs, i)
}

func EncodeTOML(i interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	if err := toml.NewEncoder(&buf).Encode(i); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func EncodeTOMLIndent(indent string) EncodeFunc {
	return func(i interface{}) ([]byte, error) {
		buf := bytes.Buffer{}
		encoder := toml.NewEncoder(&buf)
		encoder.Indent = indent
		if err := encoder.Encode(i); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}
