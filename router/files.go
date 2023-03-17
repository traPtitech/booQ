package router

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/labstack/echo"
	"github.com/traPtitech/booQ/model"
	"github.com/traPtitech/booQ/storage"
)

// アップロードを許可するMIMEタイプ
var uploadableMimes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

// PostFile POST /files
func PostFile(c echo.Context) error {
	user := c.Get("user").(model.User)

	// フォームデータ確認
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// ファイルタイプ確認
	if !uploadableMimes[fileHeader.Header.Get(echo.HeaderContentType)] {
		return c.JSON(http.StatusBadRequest, errors.New("アップロードできないファイル形式です"))
	}

	// ファイルオープン
	file, err := fileHeader.Open()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	defer file.Close()

	// サムネイル画像化
	orig, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		// 不正な画像
		return c.JSON(http.StatusBadRequest, errors.New("不正なファイルです"))
	}
	// 背景を透明にする
	newImg := image.NewRGBA(orig.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), orig, orig.Bounds().Min, draw.Over)
	b := &bytes.Buffer{}
	_ = imaging.Encode(b, imaging.Fit(newImg, 360, 480, imaging.Linear), imaging.JPEG, imaging.JPEGQuality(85))

	// 保存
	f, err := model.CreateFile(user.ID, b, "jpg")
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// レスポンス
	return c.JSON(http.StatusCreated, map[string]interface{}{"id": f.ID, "url": fmt.Sprintf("/api/files/%d", f.ID)})
}

// GetFile GET /files/:id
func GetFile(c echo.Context) error {
	// IDチェック
	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	// ファイルオープン
	f, err := storage.Open(fmt.Sprintf("%d.jpg", ID))
	if err != nil {
		if err == storage.ErrFileNotFound {
			return c.NoContent(http.StatusNotFound)
		}
		return c.NoContent(http.StatusInternalServerError)
	}
	defer f.Close()

	// レスポンス
	c.Response().Header().Set("Cache-Control", "private, max-age=31536000")
	return c.Stream(http.StatusOK, "image/jpeg", f)
}
