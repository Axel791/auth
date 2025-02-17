package domains

import (
	"fmt"
	"regexp"
)

const (
	ValidatePasswordError = "password validation failed"
	ValidateLoginError    = "login validation failed"
)

type User struct {
	id       int64
	login    string
	password string
}

func (u *User) ValidateLogin() error {
	if u.login == "" {
		return fmt.Errorf(ValidateLoginError)
	}
	latinRegex := regexp.MustCompile(`^[A-Za-z]+$`)
	if !latinRegex.MatchString(u.login) {
		return fmt.Errorf(ValidateLoginError)
	}

	return nil
}

func (u *User) ValidatePassword() error {
	if len(u.password) < 6 {
		return fmt.Errorf(ValidatePasswordError)
	}
	return nil
}
