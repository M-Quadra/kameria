package virtualfs

import (
	"bytes"
	"errors"
	"io/fs"
)

// FS []byte to fs.FS
type FS map[string]*file

var _ fs.FS = (*FS)(nil)

// Open get virtual fs.File
func (fs FS) Open(name string) (fs.File, error) {
	if f, ok := fs[name]; ok {
		return f, nil
	}
	return nil, errors.New("no such file: " + name)
}

// NewFS []byte -> fs.FS, no copy
func NewFS(info map[string][]byte) *FS {
	fs := make(FS, len(info))
	for name := range info {
		data := info[name]
		if len(data) <= 0 {
			continue
		}

		fs[name] = &file{
			Reader: bytes.NewReader(data),
			info: fileInfo{
				name: name,
				data: data,
			},
		}
	}
	return &fs
}
