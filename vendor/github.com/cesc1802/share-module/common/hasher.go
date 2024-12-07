package common

import (
	"crypto/md5"
	"encoding/hex"
)

type Hasher interface {
	Hash(plainText string) string
}

type MD5Hasher struct{}

func NewMD5Hasher() *MD5Hasher {
	return &MD5Hasher{}
}

func (h *MD5Hasher) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
