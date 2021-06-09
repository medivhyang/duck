package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func LoadFile(filePath string, target interface{}, decode DecodeFunc) error {
	if decode == nil {
		return errors.New("require decode func")
	}
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return decode(bs, target)
}

func StoreFile(i interface{}, filePath string, encode EncodeFunc) error {
	if encode == nil {
		return errors.New("require encode func")
	}
	bs, err := encode(i)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, bs, os.ModePerm)
}

func LoadOrStoreFile(filePath string, value interface{}, decode DecodeFunc, encode EncodeFunc) error {
	_, err := os.Stat(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		return StoreFile(value, filePath, encode)
	}
	return LoadFile(filePath, value, decode)
}
