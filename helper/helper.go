package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetHTTPRequestJSON(method string, url string, body io.Reader, target interface{}, headers ...map[string]string) (interface{}, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// iterate optional data of headers
	for _, header := range headers {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	timeout, _ := strconv.Atoi(os.Getenv("DEFAULT_TIMEOUT"))

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	e := json.NewDecoder(r.Body).Decode(target)
	if os.Getenv("DEVELOPMENT") == "1" {
		fmt.Println("code : " + r.Status)
		fmt.Println("url : " + url)
		fmt.Println("mtd : " + method)
		fmt.Printf("%v", headers)
		fmt.Println("rsp : ")
		fmt.Printf("%+v\n", target)
	}
	if e != nil {
		return r.Status, e
	}

	return r.Status, nil
}

type ErrorStruct struct {
	Message string `json:"message"`
}

func ErrorMessage(msg string) ErrorStruct {
	msgErr := ErrorStruct{}
	msgErr.Message = msg
	return msgErr
}

func GetResponseCode(code string) int {
	if !ValidateNumeric(code) {
		responCode := strings.Split(code, " ")[0]
		i, _ := strconv.Atoi(responCode)
		return i
	} else {
		i, _ := strconv.Atoi(code)
		return i
	}
}
