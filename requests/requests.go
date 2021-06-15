package requests

import (
	"fmt"
	"io"
)

func Get(path string) Response {
	return NewBuilder().Get(path).Do()
}

func GetText(path string) (string, error) {
	return NewBuilder().Get(path).Do().Text()
}

func GetJSON(path string, result interface{}) error {
	return NewBuilder().Get(path).Do().JSON(result)
}

func SaveFile(path string, filename string) error {
	return NewBuilder().Get(path).Do().SaveFile(filename)
}

func GetStream(path string) (io.ReadCloser, error) {
	return NewBuilder().Get(path).Do().Stream()
}

func Post(path string, reader io.Reader, contentType string) Response {
	return NewBuilder().Post(path).Header("Content-Type", contentType).WriteBody(reader).Do()
}

func PostJSON(path string, body interface{}, result interface{}) error {
	return NewBuilder().Post(path).WriteJSON(body).Do().JSON(result)
}

func PostFile(path string, fileName string) Response {
	return NewBuilder().Post(path).WriteFile(fileName).Do()
}

func PostFormFile(path string, formName string, fileName string) Response {
	return NewBuilder().Post(path).WriteFormFile(formName, fileName).Do()
}

func Put(path string, reader io.Reader, contentType string) Response {
	return NewBuilder().Put(path).Header("Content-Type", contentType).WriteBody(reader).Do()
}

func PutJSON(path string, body interface{}, result interface{}) error {
	return NewBuilder().Post(path).WriteJSON(body).Do().JSON(result)
}

func Patch(path string, reader io.Reader, contentType string) Response {
	return NewBuilder().Patch(path).Header("Content-Type", contentType).WriteBody(reader).Do()
}

func PatchJSON(path string, body interface{}, result interface{}) error {
	return NewBuilder().Post(path).WriteJSON(body).Do().JSON(result)
}

func newError(module string, format string, args ...interface{}) error {
	return fmt.Errorf("%s: %s: %s", "requests", module, fmt.Sprintf(format, args...))
}
