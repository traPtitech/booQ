package storage

import (
	"errors"
	"io"
	"os"
	"path/filepath"
)

// Local ローカルストレージ
type Local struct {
	localDir string
}

// SetLocalStorage ローカルストレージをカレントストレージに設定します
func SetLocalStorage(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return errors.New("dir doesn't exist")
	}
	if !fi.IsDir() {
		return errors.New("dir is not a directory")
	}

	current = &Local{localDir: dir}
	return nil
}

func (l Local) Save(filename string, src io.Reader) error {
	file, err := os.Create(l.getFilePath(filename))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, src)
	return err
}

func (l Local) Open(filename string) (io.ReadCloser, error) {
	r, err := os.Open(l.getFilePath(filename))
	if err != nil {
		return nil, ErrFileNotFound
	}
	return r, nil
}

func (l Local) Delete(filename string) error {
	path := l.getFilePath(filename)
	if _, err := os.Stat(path); err != nil {
		return ErrFileNotFound
	}
	return os.Remove(path)
}

func (l Local) getFilePath(filename string) string {
	return filepath.Join(l.localDir, filename)
}
