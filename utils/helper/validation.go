package helper

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// ValidateEmail is a function to validate email
func EmailFormatValidation(email string) error {
	//	Check syntax email address
	pattern := `^\w+@\w+\.\w+$`
	matched, _ := regexp.Match(pattern, []byte(email))
	if !matched {
		return errors.New("failed syntax email address")
	}
	return nil
}

func ValidatePassword(pass string) string {
	var (
		upp, low, num, space, special bool
		tot                           uint8
		character                     = `#%'()+/:;<=>?[\]^{|}~_-., !@$&*`
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsSpace(char):
			space = true
			tot++
		case strings.ContainsRune(character, char):
			special = true
			tot++
		default:
			return "tidak ada password"
		}
	}

	if !upp {
		return "password must contain uppercase"
	} else if !low {
		return "password must contain lowercase"
	} else if !num {
		return "password must contain numeric"
	} else if tot < 8 {
		return "password must have minumum 8 character"
	} else if space {
		return "password cannot be filled with space"
	} else if !special {
		return "password must contain special character"
	}

	return "Valid"
}
