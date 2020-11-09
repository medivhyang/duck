package config

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

type ContentType string

const (
	JSON ContentType = "JSON"
	XML  ContentType = "XML"
	YAML ContentType = "YAML"
	TOML ContentType = "TOML"
)

func (ct ContentType) Valid() bool {
	switch ct {
	case JSON, XML, YAML, TOML:
		return true
	default:
		return false
	}
}

var ErrUnsupportedContentType = errors.New("config: unsupported content type")

var (
	Pretty = true
	Prefix = ""
	Indent = "  "
)

func Decode(contentType ContentType, source []byte, target interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	switch contentType {
	case JSON:
		if err := json.Unmarshal(source, target); err != nil {
			return err
		}
	case XML:
		if err := xml.Unmarshal(source, target); err != nil {
			return err
		}
	case YAML:
		if err := yaml.Unmarshal(source, target); err != nil {
			return err
		}
	case TOML:
		if err := toml.Unmarshal(source, target); err != nil {
			return err
		}
	}
	return nil
}

func DecodeString(contentType ContentType, source string, target interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	return Decode(contentType, []byte(source), target)
}

func DecodeReader(contentType ContentType, source io.Reader, target interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	switch contentType {
	case JSON:
		return json.NewDecoder(source).Decode(target)
	case XML:
		return xml.NewDecoder(source).Decode(target)
	case YAML:
		return yaml.NewDecoder(source).Decode(target)
	case TOML:
		_, err := toml.DecodeReader(source, target)
		return err
	}
	return nil
}

func Encode(contentType ContentType, source interface{}) ([]byte, error) {
	if !contentType.Valid() {
		return nil, ErrUnsupportedContentType
	}
	switch contentType {
	case JSON:
		if Pretty {
			return json.MarshalIndent(source, Prefix, Indent)
		}
		return json.Marshal(source)
	case XML:
		if Pretty {
			return xml.MarshalIndent(source, Prefix, Indent)
		}
		return xml.Marshal(source)
	case YAML:
		return yaml.Marshal(source)
	case TOML:
		buf := bytes.Buffer{}
		encoder := toml.NewEncoder(&buf)
		if Pretty {
			encoder.Indent = Indent
		}
		if err := encoder.Encode(source); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
	return nil, nil
}

func EncodeToString(contentType ContentType, source interface{}) (string, error) {
	if !contentType.Valid() {
		return "", ErrUnsupportedContentType
	}
	bs, err := Encode(contentType, source)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func EncodeToWriter(contentType ContentType, source interface{}, target io.Writer) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	switch contentType {
	case JSON:
		encoder := json.NewEncoder(target)
		if Pretty {
			encoder.SetIndent(Prefix, Indent)
		}
		encoder.Encode(source)
	case XML:
		encoder := xml.NewEncoder(target)
		if Pretty {
			encoder.Indent(Prefix, Indent)
		}
		encoder.Encode(source)
	case YAML:
		return yaml.NewEncoder(target).Encode(source)
	case TOML:
		encoder := toml.NewEncoder(target)
		if Pretty {
			encoder.Indent = Indent
		} else {
			encoder.Indent = ""
		}
		encoder.Encode(source)
	}
	return nil
}

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvInt(key string) (int, error) {
	return strconv.Atoi(GetEnv(key))
}

func GetEnvInt64(key string) (int64, error) {
	return strconv.ParseInt(GetEnv(key), 10, 64)
}

func GetEnvFloat64(key string) (float64, error) {
	return strconv.ParseFloat(GetEnv(key), 64)
}

func GetEnvBool(key string) (bool, error) {
	return strconv.ParseBool(GetEnv(key))
}

func GetEnvOrDefault(key string, defaultValue ...string) string {
	s, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return s
}

func GetEnvIntOrDefault(key string, defaultValue ...int) (int, error) {
	v, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return strconv.Atoi(v)
}

func GetEnvInt64OrDefault(key string, defaultValue ...int64) (int64, error) {
	v, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return strconv.ParseInt(GetEnv(v), 10, 64)
}

func GetEnvFloat64OrDefault(key string, defaultValue ...float64) (float64, error) {
	v, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return strconv.ParseFloat(GetEnv(v), 64)
}

func GetEnvBoolOrDefault(key string, defaultValue ...bool) (bool, error) {
	v, ok := os.LookupEnv(key)
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0], nil
	}
	return strconv.ParseBool(GetEnv(v))
}

func LoadEnv(contentType ContentType, key string, target interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	content := os.Getenv(key)
	if content == "" {
		return nil
	}
	return DecodeString(contentType, content, target)
}

func LoadFile(contentType ContentType, filePath string, target interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return DecodeReader(contentType, file, target)
}

func StoreFile(contentType ContentType, source interface{}, filePath string) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return EncodeToWriter(contentType, source, file)
}

func LoadOrStoreFile(contentType ContentType, filePath string, value interface{}) error {
	if !contentType.Valid() {
		return ErrUnsupportedContentType
	}
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return StoreFile(contentType, value, filePath)
	}
	return LoadFile(contentType, filePath, value)
}
