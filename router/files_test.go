package router

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traPtitech/booQ/model"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"testing"
)

var testJpeg = `/9j/4AAQSkZJRgABAQIAOAA4AAD/2wBDAP//////////////////////////////////////////////////////////////////////////////////////2wBDAf//////////////////////////////////////////////////////////////////////////////////////wAARCAABAAEDAREAAhEBAxEB/8QAHwAAAQUBAQEBAQEAAAAAAAAAAAECAwQFBgcICQoL/8QAtRAAAgEDAwIEAwUFBAQAAAF9AQIDAAQRBRIhMUEGE1FhByJxFDKBkaEII0KxwRVS0fAkM2JyggkKFhcYGRolJicoKSo0NTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uHi4+Tl5ufo6erx8vP09fb3+Pn6/8QAHwEAAwEBAQEBAQEBAQAAAAAAAAECAwQFBgcICQoL/8QAtREAAgECBAQDBAcFBAQAAQJ3AAECAxEEBSExBhJBUQdhcRMiMoEIFEKRobHBCSMzUvAVYnLRChYkNOEl8RcYGRomJygpKjU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6goOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4uPk5ebn6Onq8vP09fb3+Pn6/9oADAMBAAIRAxEAPwBKBH//2Q`

func TestPostFile(t *testing.T) {
	t.Parallel()
	e := echoSetupWithUser()

	t.Run("no form", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		req := httptest.NewRequest(echo.POST, "/api/files", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("invalid file type", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		go func() {
			defer writer.Close()
			part, _ := writer.CreateFormFile("file", "test.txt")
			_, _ = part.Write([]byte("test text file"))
		}()

		req := httptest.NewRequest(echo.POST, "/api/files", pr)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("bad image", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		go func() {
			defer writer.Close()
			h := textproto.MIMEHeader{}
			h.Set("Content-Disposition", `form-data; name="file"; filename="test.jpg"`)
			h.Set("Content-Type", "image/jpeg")
			part, _ := writer.CreatePart(h)
			_, _ = part.Write([]byte("test text file"))
		}()

		req := httptest.NewRequest(echo.POST, "/api/files", pr)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusBadRequest, rec.Code)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)
		go func() {
			defer writer.Close()
			h := textproto.MIMEHeader{}
			h.Set("Content-Disposition", `form-data; name="file"; filename="test.jpg"`)
			h.Set("Content-Type", "image/jpeg")
			part, _ := writer.CreatePart(h)
			img, _ := base64.RawStdEncoding.DecodeString(testJpeg)
			_, _ = part.Write(img)
		}()

		req := httptest.NewRequest(echo.POST, "/api/files", pr)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
	})
}

func TestGetFile(t *testing.T) {
	t.Parallel()
	e := echoSetupWithUser()

	t.Run("not found", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		req := httptest.NewRequest(echo.GET, "/api/files/a", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusNotFound, rec.Code)
	})

	t.Run("not found2", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		req := httptest.NewRequest(echo.GET, "/api/files/999999", nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusNotFound, rec.Code)
	})

	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		assert := assert.New(t)

		img, _ := base64.RawStdEncoding.DecodeString(testJpeg)
		f, err := model.CreateFile(1, bytes.NewReader(img), "jpg")
		require.NoError(t, err)

		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/api/files/%d", f.ID), nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		assert.Equal(http.StatusOK, rec.Code)
		assert.EqualValues(img, rec.Body.Bytes())
	})
}
