package entity

import "strings"

type AvailableEmailDomain string

const (
	GmailDomain AvailableEmailDomain = "gmail.com"
)

type Auth struct {
	Email    string
	Password string
}

func (a *Auth) CheckAvailableEmail() bool {
	// @以降がgmail.comのみ許可
	userDomains := strings.Split(a.Email, "@")
	if len(userDomains) != 2 {
		return false
	}
	if userDomains[1] != string(GmailDomain) {
		return false
	}
	return true
}
