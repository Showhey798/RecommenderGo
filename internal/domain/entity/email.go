package entity

import "strings"

type AvailableEmailDomain string

const (
	GmailDomain AvailableEmailDomain = "gmail.com"
)

type Email string

func (e Email) CheckAvailableEmail() bool {
	// @以降がgmail.comのみ許可
	userDomains := strings.Split(string(e), "@")
	if len(userDomains) != 2 {
		return false
	}
	if userDomains[1] != string(GmailDomain) {
		return false
	}
	return true
}
