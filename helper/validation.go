package helper

import (
	"regexp"
	"errors"

)

const (

	// this block is for validating URL format
	email        string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	ip           string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	urlSchema    string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	urlUsername  string = `(\S+(:\S*)?@)`
	urlPath      string = `((\/|\?|#)[^\s]*)`
	urlPort      string = `(:(\d{1,5}))`
	urlIP        string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	urlSubdomain string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	url          string = `^` + urlSchema + `?` + urlUsername + `?` + `((` + urlIP + `|(\[` + ip + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + urlSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + urlPort + `?` + urlPath + `?$`
)

var (
	// ErrBadFormatURL variable for error of url format
	ErrBadFormatURL = errors.New("invalid url format")
	// ErrBadFormatMail variable for error of email format
	ErrBadFormatMail = errors.New("invalid email format")
	// emailRegexp regex for validate email
	emailRegexp = regexp.MustCompile(email)
	// URLRegexp regex for validate URL
	urlRegexp = regexp.MustCompile(url)
	// alphaNumericRegexp regex for validate uri
	alphaNumericRegexp = regexp.MustCompile("^[A-Za-z0-9][A-Za-z0-9]*$")
	alphaRegexp        = regexp.MustCompile("^[A-Za-z][A-Za-z ._-]*$")
	numericRegexp      = regexp.MustCompile("^[0-9][0-9 ._-]*$")
)

func ValidateNumeric(str string) bool {
	if !numericRegexp.MatchString(str) {
		return false
	}
	return true
}

func ValidateAlpha(str string) bool {
	if !alphaRegexp.MatchString(str) {
		return false
	}
	return true
}

func ValidateAlphaNumeric(str string) bool {
	if !alphaNumericRegexp.MatchString(str) {
		return false
	}
	return true
}

// ValidateEmail function for validating email
func ValidateEmail(email string) bool {
	if !emailRegexp.MatchString(email) {
		return false
	}
	return true
}

// ValidateAlphabetWithSpace func for check alphabet with space
func ValidateAlphabetWithSpace(str string) bool {
	var uppercase, lowercase, space, symbol int
	for _, r := range str {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = +1
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = +1
		} else if r == 32 { //code ascii for space
			space = +1
		} else { //except alphabet
			symbol = +1
		}
	}

	if symbol > 0 {
		return false
	}
	return uppercase >= 1 || lowercase >= 1 || space >= 1
}