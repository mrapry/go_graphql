package middleware

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestBearerVerify(t *testing.T) {
	expiredToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG0iOmZhbHNlLCJhdWQiOiJiaGlubmVrYS1taWNyb3NlcnZpY2VzLWIxMzcxNC01MzEyMTE1IiwiYXV0Ijp0cnVlLCJkaWQiOiJjMGI0ZDFiNGM0NDc0IiwiZGxpIjoiV0VCIiwiZXhwIjoxNTIxMDIyNTI0LCJpYXQiOjE1MjEwMjI0NjQsImlzcyI6ImJoaW5uZWthLmNvbSIsInN1YiI6IlVTUjE4MDIxODE3NSJ9.gw_tfbPHq6XIWuVT3ksHFovWkYZteUuuLGepkGeAnAGP41pF_AlEhcB120Jao1FOi74li7f3ab6kN_hBcMNMuYmhlkP4pa78QKFCY7uZpLUm6LIc2AOHf1VRm0poQnvH0AnDDw1_bU8NFe0GKr48Cf88934txTCJRQ75Sw4pnbk"
	validAccessToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG0iOnRydWUsImF1ZCI6ImJoaW5uZWthLW1pY3Jvc2VydmljZXMtYjEzNzE0LTUzMTIxMTUiLCJhdXRob3Jpc2VkIjp0cnVlLCJkaWQiOiJjMGI0ZDFiNGM0NDc0IiwiZGxpIjoiV0VCIiwiZXhwIjoxNTQ2MjQzNjg0LCJpYXQiOjE1MjQ2NDM2ODQsImlzcyI6ImJoaW5uZWthLmNvbSIsInN1YiI6IlVTUjE4MDMxODE3NiJ9.FICCcN1wBzBIQqYlfeyhovboTaIQb3LCqK662xRUGoKw0RcAwEp18BEZsKX13KRLvN1AysnW-1FMeA1hze6tdTxB04vWJKRcq9atxPSVmOuzRqsDCJZLaYbIjLkiDnOPvcpQixbqbbqeOhH3NUfwybm2dnRlCCInVaFWWWh21bU"
	anonymousAccessToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG0iOmZhbHNlLCJhdWQiOiJiaGlubmVrYS1taWNyb3NlcnZpY2VzLWIxMzcxNC01MzEyMTE1IiwiYXV0aG9yaXNlZCI6ZmFsc2UsImRpZCI6ImMwYjRkMWI0YzQ0NzQiLCJkbGkiOiJXRUIiLCJleHAiOjE1NDYyNDE2ODUsImlhdCI6MTUyNDY0MTY4NSwiaXNzIjoiYmhpbm5la2EuY29tIiwic3ViIjoiYmhpbm5la2EtbWljcm9zZXJ2aWNlcy1iMTM3MTQtNTMxMjExNSJ9.U0XGYR_ZiYUDA1aRHgvbYqp0zF8saN-O4_H7793Ou8OfQKsU-t5NRqC6cyeImBN8ayh3o3s35_4AvAuGH9uLrlq3MuUxPZnsmU6SlJ0ODen0O8ak6i3PZF_6dltLA5ZsTgVt4YvOydSQusTqHoN-jyFsBFtZHKHCFoFVK9Hqu-Q"

	t.Run("Test Positive case", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, validAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "bhinneka.com"))
		})

		mw := BearerVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Test Positive case with Anonymous Access Token", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, anonymousAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "bhinneka.com"))
		})

		mw := BearerVerify(verifyKey, false)(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Test Negative case, Will error 'need an authorised user'", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, anonymousAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "bhinneka.com"))
		})

		mw := BearerVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.Error(t, err)
	})

	t.Run("Test Negative Will error 'Access Token Expired'", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, expiredToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "bhinneka.com"))
		})

		mw := BearerVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.Error(t, err)
	})

}

const ValidPublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCoqzL5JrMzed4tb8uEoLKd42EO
sYmb0HpbicGt/OUeJxaHtt59Ew0BbpreBeiuugXweEa5xctQOxGYr27h4ZOnR0hW
Si+h5Y35CKzMEmZnzQwzQphgqww0U+e9/OAvVfCW1xWvVFr0WbhIRn+w/9DUvp+6
jKz3fIj3yQaHWVMMNQIDAQAB
-----END PUBLIC KEY-----`

func getPublicKey(publicKey string) (*rsa.PublicKey, error) {
	r := strings.NewReader(publicKey)
	verifyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}
