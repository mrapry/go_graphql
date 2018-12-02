package middleware

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"strconv"
	"time"
)

const secretKey = "sup3rs3cr3t!!"

func HmacCode(appId string, random string) string {
	combineKey := secretKey + appId
	h := hmac.New(sha512.New, []byte(combineKey))
	h.Write([]byte(random))
	res := base64.URLEncoding.EncodeToString(h.Sum(nil))

	return res
}

func VerifyHmac(hmac string, appId string, random string) bool {
	if !CheckRandom(random) {
		return false
	}

	res := HmacCode(appId, random)

	if res != hmac {
		return false
	}

	return true
}

func CheckRandom(random string) bool {
	randomTime, _ := strconv.Atoi(random)
	timeNow := int(time.Now().Unix())

	if (timeNow - randomTime) > 120 {
		return false
	}
	return true
}
