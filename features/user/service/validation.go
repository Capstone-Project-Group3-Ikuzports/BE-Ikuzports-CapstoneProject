package service

import (
	"errors"
	"regexp"
)

// ValidateEmail is a function to validate email
func emailFormatValidation(email string) error {
	//	Check syntax email address
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(email))
	if !matched {
		return errors.New("failed syntax email address")
	}
	return nil
}
