package helper

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// CHARS for setting short random string
	CHARS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// ErrorDataNotFound error message when data doesn't exist
	ErrorDataNotFound = "data %s not found"
	// ErrorParameterInvalid error message for parameter is invalid
	ErrorParameterInvalid = "%s parameter is invalid"
	// ErrorParameterRequired error message for parameter is missing
	ErrorParameterRequired = "%s parameter is required"
	// SuccessMessage message for success process
	SuccessMessage = "succeed to process data"
)

// GenerateRandomID function for generating shipping ID
func GenerateRandomID(length int, prefix ...string) string {
	var strPrefix string

	if len(prefix) > 0 {
		strPrefix = prefix[0]
	}

	yearNow, monthNow, _ := time.Now().Date()
	year := strconv.Itoa(yearNow)[2:len(strconv.Itoa(yearNow))]
	month := int(monthNow)
	RandomString := RandomString(length)

	id := fmt.Sprintf("%s%s%d%s", strPrefix, year, month, RandomString)
	return id
}

// RandomString function for random string
func RandomString(length int) string {
	rand.Seed(time.Now().UTC().UnixNano())

	charsLength := len(CHARS)
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = CHARS[rand.Intn(charsLength)]
	}
	return string(result)
}

// RandomStringBase64 function for random string and base64 encoded
func RandomStringBase64(length int) string {
	rb := make([]byte, length)
	_, err := rand.Read(rb)

	if err != nil {
		return ""
	}
	rs := base64.URLEncoding.EncodeToString(rb)

	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		return ""
	}

	return reg.ReplaceAllString(rs, "")
}

// ClearHTML function for validating HTML
func ClearHTML(src string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return re.ReplaceAllString(src, "")
}

// StringInSlice function for checking whether string in slice
// str string searched string
// list []string slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// SetLastName function for setting last name
func SetLastName(names []string) string {
	if len(names) <= 1 {
		return ""
	}

	names = append(names[:0], names[1:]...)

	return strings.Join(names, " ")
}
