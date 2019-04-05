package utils

import (
	"os"
	"crypto/md5"
	"io"
	"encoding/hex"
	"github.com/sirupsen/logrus"
)

func HashMD5(f string) string {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		return ""
	}

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		logrus.Errorf("getfilemd5 error: %v", err)
		return ""
	}
	r := h.Sum(nil)[:16]
	hex.EncodeToString(r)
	return hex.EncodeToString(r)
}
