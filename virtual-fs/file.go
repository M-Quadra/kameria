package virtualfs

import (
	"bytes"
	"io/fs"
)

type file struct {
	*bytes.Reader
	info fileInfo
}

func (f file) Stat() (fs.FileInfo, error) {
	return f.info, nil
}

func (f file) Readdir(count int) ([]fs.FileInfo, error) {
	return nil, nil
}

func (f file) Close() error {
	return nil
}
