package helper

import (
	"crypto/rand"
	"encoding/base64"
)

// Tạo chuỗi ngẫu nhiên để sử dụng làm giá trị cho khóa
func GenerateRandomValue() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
