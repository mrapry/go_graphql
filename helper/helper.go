package helper

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
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
	// ErrorParameterLength error message for parameter length is invalid
	ErrorParameterLength = "length of %s parameter exceeds the limit %d"
	// ErrorUnauthorized error message for unauthorized user
	ErrorUnauthorized = "you are not authorized"
	// SuccessMessage message for success process
	SuccessMessage = "succeed to process data"
	// ErrorRedis redis error
	ErrorRedis = "redis: nil"
	// ErrorLoginAttempt message for reaching login attempt limit
	ErrorLoginAttempt = "already reach login attempt limit"

	// FormatDateDB for inserting date to database
	FormatDateDB = "2006-01-02"
	//FormatDOB for parsing dob from form value
	FormatDOB = "02/01/2006"
	// DefaultDOB default value
	DefaultDOB = "01/01/0001"
)

// GetHTTPNewRequest function for getting json API
func GetHTTPNewRequest(method, url string, body io.Reader, target interface{}, headers ...map[string]string) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	e := json.NewDecoder(r.Body).Decode(target)
	if e != nil {
		return e
	}

	return nil
}

// GetHTTPNewReqBasicAuth function for getting json API with basic auth
func GetHTTPNewReqBasicAuth(method, url, userName, password string, body io.Reader, target interface{}, headers ...map[string]string) error {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	req.SetBasicAuth(userName, password)
	client := &http.Client{Timeout: 10 * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	e := json.NewDecoder(r.Body).Decode(target)
	if e != nil {
		return e
	}

	return nil
}

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

// GenerateMemberID function for generating member id
// based on member max id
func GenerateMemberID(lastUserID string) string {
	preffix := "USR"
	var strNextID string

	intLastUserID, _ := strconv.Atoi(lastUserID)

	if intLastUserID < 1000 {
		strNextID = "0" + strconv.Itoa(intLastUserID+1)
	} else if intLastUserID == 99999999 {
		strNextID = "00000001"
	} else {
		strNextID = strconv.Itoa(intLastUserID + 1)
	}

	now := time.Now()
	dt := now.Format("06") + now.Format("01")

	return preffix + dt + strNextID
}

// GenerateTokenByString function for generate token by string
func GenerateTokenByString(str string) string {
	var token string

	// generate random string
	RandomString := RandomString(30)

	mix := str + "-" + RandomString
	md := md5.New()
	io.WriteString(md, mix)
	sumMd5 := md.Sum(nil)
	hash := hex.EncodeToString(sumMd5[:])

	sh := sha1.New()
	io.WriteString(sh, hash)
	sumSha := sh.Sum(nil)
	token = hex.EncodeToString(sumSha[:])

	return token
}
