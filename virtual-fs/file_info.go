package virtualfs

import (
	"io/fs"
	"time"
)

type fileInfo struct {
	fs.FileInfo

	name       string
	data       []byte
	createTime time.Time
}

func (fi fileInfo) Name() string {
	return fi.name
}

func (fi fileInfo) Size() int64 {
	return int64(len(fi.data))
}

func (fi fileInfo) Mode() fs.FileMode {
	return fs.FileMode(0444)
}

func (fi fileInfo) ModTime() time.Time {
	return fi.createTime
}

func (fi fileInfo) IsDir() bool {
	return false
}

func (fi fileInfo) Sys() interface{} {
	return nil
}
