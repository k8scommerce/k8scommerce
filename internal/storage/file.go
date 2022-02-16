package storage

import (
	"bytes"
)

type AssetKind int

const (
	AssetKind_Image    AssetKind = 0
	AssetKind_Document AssetKind = 1
	AssetKind_Audio    AssetKind = 2
	AssetKind_Video    AssetKind = 3
)

type File struct {
	Name      string
	AssetKind AssetKind
	Mime      string

	buffer *bytes.Buffer
}

func NewFile() *File {
	return &File{
		buffer: &bytes.Buffer{},
	}
}

func (f *File) Write(chunk []byte) error {
	_, err := f.buffer.Write(chunk)

	return err
}
