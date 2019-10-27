package storage

import (
	"errors"
	"io"
)

var current Storage = &Memory{files: map[string][]byte{}}

var (
	// ErrFileNotFound 指定されたファイル名のファイルは存在しません
	ErrFileNotFound = errors.New("not found")
)

// Storage ストレージインターフェース
type Storage interface {
	Save(filename string, src io.Reader) error
	Open(filename string) (io.ReadCloser, error)
	Delete(filename string) error
}

// Save 指定したファイル名で保存します。同名のファイルは上書きされます。
func Save(filename string, src io.Reader) error {
	return current.Save(filename, src)
}

// Open 指定したファイルを開きます
func Open(filename string) (io.ReadCloser, error) {
	return current.Open(filename)
}

// Delete 指定したファイルを削除します
func Delete(filename string) error {
	return current.Delete(filename)
}
