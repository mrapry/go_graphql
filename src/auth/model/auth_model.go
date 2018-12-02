package model

import "time"

const (
	// InvalidUsername constanta
	InvalidUsername = "username basic auth is not found"
	// InvalidAuth constanta
	InvalidAuth = "basic auth is invalid"
	// InvalidDeviceLogin constanta
	InvalidDeviceLogin = "device login parameter is invalid"
	// NowRowsError constanta
	NowRowsError = "sql: no rows in result set"
	// InvalidRefreshToken constanta
	InvalidRefreshToken = "refresh token is invalid"
	// InvalidGoogleCode constanta
	InvalidGoogleCode = "failed to get data from google, check the token you provide"
	// InvalidFacebookCode constanta
	InvalidFacebookCode = "activate your account first"
	// AuthorizationError constanta
	AuthorizationError = "authorization is empty"
	// FailedAzureToken constanta
	FailedAzureToken = "failed to get azure token"
)

// RequestToken data structure
type RequestToken struct {
	GrantType   string
	Username    string
	Password    string
	DeviceId    string
	DeviceLogin string
}
type RefreshToken struct {
	GrantType      string
	RefreshToken   string
	OldAccessToken string
	DeviceId       string
	DeviceLogin    string
}
type AnonymusToken struct {
	GrantType   string
	DeviceId    string
	DeviceLogin string
}
type GoogleToken struct {
	GrantType   string
	Code        string
	DeviceId    string
	DeviceLogin string
}
type FacebookToken struct {
	GrantType   string
	Code        string
	DeviceId    string
	DeviceLogin string
}
type AzureToken struct {
	GrantType   string
	Code        string
	DeviceId    string
	DeviceLogin string
	RedirectUri string
}

type AuthRespon struct {
	Message string `json:"message,omitempty"`
	Data    struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Email        string    `json:"email"`
			ExpiredTime  time.Time `json:"expiredTime"`
			FirstName    string    `json:"firstName"`
			HasPassword  bool      `json:"hasPassword"`
			LastName     string    `json:"lastName"`
			NewMember    bool      `json:"newMember"`
			RefreshToken string    `json:"refreshToken"`
			Token        string    `json:"token"`
			UserID       string    `json:"userId"`
		} `json:"attributes"`
	} `json:"data,omitempty"`
}

type DataMemberAuthService struct {
	Data struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Address struct {
				Province      string `json:"province"`
				ProvinceID    string `json:"provinceId"`
				City          string `json:"city"`
				CityID        string `json:"cityId"`
				District      string `json:"district"`
				DistrictID    string `json:"districtId"`
				SubDistrict   string `json:"subDistrict"`
				SubDistrictID string `json:"subDistrictId"`
				ZipCode       string `json:"zipCode"`
				Street1       string `json:"street1"`
				Street2       string `json:"street2"`
			} `json:"address"`
			BirthDate    string    `json:"birthDate"`
			Created      time.Time `json:"created"`
			Department   string    `json:"department"`
			Email        string    `json:"email"`
			Ext          string    `json:"ext"`
			FirstName    string    `json:"firstName"`
			Gender       string    `json:"gender"`
			HasPassword  bool      `json:"hasPassword"`
			IsAdmin      bool      `json:"isAdmin"`
			IsStaff      bool      `json:"isStaff"`
			JobTitle     string    `json:"jobTitle"`
			LastBlocked  time.Time `json:"lastBlocked"`
			LastLogin    time.Time `json:"lastLogin"`
			LastModified time.Time `json:"lastModified"`
			LastName     string    `json:"lastName"`
			Mobile       string    `json:"mobile"`
			Phone        string    `json:"phone"`
			SignUpFrom   string    `json:"signUpFrom"`
			SocialMedia  struct {
				FacebookID string `json:"facebookId"`
				GoogleID   string `json:"googleId"`
				AzureID    string `json:"azureId"`
			} `json:"socialMedia, omitempty"`
			Status  string `json:"status"`
			Version int    `json:"version"`
		} `json:"attributes"`
	} `json:"data"`
}

type DataMember struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	NewMember bool   `json:"newMember"`
}

// Token struct
type Token struct {
	Data         *DataMember `json:"data,omitempty"`
	AccessToken  string      `json:"accessToken,omitempty"`
	RefreshToken string      `json:"refreshToken,omitempty"`
}
