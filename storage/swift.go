package storage

import (
	"github.com/ncw/swift"
	"io"
)

// Swift Swiftオブジェクトストレージ
type Swift struct {
	container string
	conn      *swift.Connection
}

// SetSwiftStorage Swiftオブジェクトストレージをカレントストレージに設定します
func SetSwiftStorage(container, userName, apiKey, tenant, tenantID, authURL string) error {
	conn := &swift.Connection{
		AuthUrl:  authURL,
		UserName: userName,
		ApiKey:   apiKey,
		Tenant:   tenant,
		TenantId: tenantID,
	}

	// 認証
	if err := conn.Authenticate(); err != nil {
		return err
	}

	// コンテナの存在を確認
	if _, _, err := conn.Container(container); err != nil {
		return err
	}

	current = &Swift{
		container: container,
		conn:      conn,
	}
	return nil
}

func (s *Swift) Save(filename string, src io.Reader) error {
	_, err := s.conn.ObjectPut(s.container, filename, src, true, "", "", swift.Headers{})
	return err
}

func (s *Swift) Open(filename string) (io.ReadCloser, error) {
	file, _, err := s.conn.ObjectOpen(s.container, filename, true, nil)
	if err != nil {
		if err == swift.ObjectNotFound {
			return nil, ErrFileNotFound
		}
		return nil, err
	}
	return file, nil
}

func (s *Swift) Delete(filename string) error {
	err := s.conn.ObjectDelete(s.container, filename)
	if err != nil {
		if err == swift.ObjectNotFound {
			return ErrFileNotFound
		}
		return err
	}
	return nil
}
