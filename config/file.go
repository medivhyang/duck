package config

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{path: path}
}

func (f *File) Load(contentType ContentType, target interface{}) error {
	return LoadFile(contentType, f.path, target)
}

func (f *File) Store(contentType ContentType, source interface{}) error {
	return StoreFile(contentType, source, f.path)
}

func (f *File) LoadOrStore(contentType ContentType, value interface{}) error {
	return LoadOrStoreFile(contentType, f.path, value)
}
