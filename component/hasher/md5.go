package hasher

import (
	"crypto/md5"
	"encoding/hex"
)

type md5hash struct {
}

func NewMd5Hash() *md5hash {
	return &md5hash{}
}

func (h *md5hash) Hash(data string) string {
	hasher := md5.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}
