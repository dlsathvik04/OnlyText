package jwt

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type tokenHeader struct {
	Typ       string
	Algorithm string
}

type TokenBody struct {
	Expiration time.Time
	Payload    interface{}
	Provider   string
}

func GenerateToken(payload interface{}, duration time.Duration, provider string) string {
	header := tokenHeader{
		Typ:       "JWT",
		Algorithm: "SHA256",
	}

	headerdata, err := json.Marshal(header)
	if err != nil {
		log.Panic(err)
	}

	body := TokenBody{
		Expiration: time.Now().Add(duration),
		Payload:    payload,
		Provider:   provider,
	}
	bodydata, err := json.Marshal(body)

	tokendata := base64.URLEncoding.EncodeToString(headerdata) + "." + base64.URLEncoding.EncodeToString(bodydata)

	if err != nil {
		log.Panic(err)
	}

	hasher := sha256.New()
	hasher.Write([]byte(tokendata))
	signature := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return tokendata + "." + signature

}

func AuthorizeToken(token string) (TokenBody, bool) {
	token_segments := strings.Split(token, ".")

	tokendata := token_segments[0] + "." + token_segments[1]

	hasher := sha256.New()
	hasher.Write([]byte(tokendata))
	hash := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	if hash == token_segments[2] {
		var tokenBody TokenBody
		data, err := base64.URLEncoding.DecodeString(token_segments[1])
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(data, &tokenBody)
		if err != nil {
			log.Fatal(err)
		}

		if time.Now().After(tokenBody.Expiration) {
			return TokenBody{}, false
		}

		return tokenBody, true
	}
	return TokenBody{}, false
}
