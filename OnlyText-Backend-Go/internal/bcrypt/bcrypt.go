package bcrypt

import (
	"crypto/sha256"
	"encoding/base64"
)

type Hasher struct {
	Salt string
}

func (h *Hasher) Hash(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password + h.Salt))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (h *Hasher) CompareHash(password, hash string) bool {
	hasher := sha256.New()
	hasher.Write([]byte(password + h.Salt))

	return base64.URLEncoding.EncodeToString(hasher.Sum(nil)) == hash
}
