package service

import "net/url"

type ResultinService struct {
	Result     interface{}
	HTTPStatus interface{}
	Error      error
}

type AuthService interface {
	AuthPassword(form url.Values, headerAuth string) <-chan ResultinService
	GetDetailMember(headerAuth string) <-chan ResultinService
}