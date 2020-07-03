package model

import (
	"fmt"

	"github.com/traPtitech/booQ/storage"
	"io"
)

// File アップロードファイルの構造体
type File struct {
	GormModel
	UploadUserID uint `gorm:"type:int;not null"`
}

// TableName dbのテーブル名を指定する
func (File) TableName() string {
	return "files"
}

// CreateFile Fileを作成する
func CreateFile(uploadUserID uint, src io.Reader, ext string) (File, error) {
	// トランザクション開始
	tx := db.Begin()
	defer tx.RollbackUnlessCommitted()

	f := File{UploadUserID: uploadUserID}

	// DBにレコード作成
	if err := tx.Create(&f).Error; err != nil {
		return File{}, err
	}

	// ストレージに保存
	if err := storage.Save(fmt.Sprintf("%d.%s", f.ID, ext), src); err != nil {
		return File{}, err
	}

	// コミット
	return f, tx.Commit().Error
}
