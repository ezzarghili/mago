package validator

import (
	"regexp"
	"strings"
)

// PhoneType used to destinguish between phone number types
type PhoneType int8

const (
	// PhoneAll when validating landline phone numbers
	PhoneAll PhoneType = iota
	// PhoneLandLine when validating landline phone numbers
	PhoneLandLine
	// PhoneMobile when validating landline phone numbers
	PhoneMobile
)

// Phone options to customize to validator
type Phone struct {
	Location string    // Provinces and prefectures ISO code
	Spacer   string    // regex expression
	Type     PhoneType // defaults to 0 == PhoneAll
}

// Validate if phone number is correct
func (validator *Phone) Validate(cine string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(cine, spacerReg))
	if validator.Type == PhoneLandLine {
		if validator.Location != "" {
			reg := regexp.MustCompile(`^(\+?212|0)5\d{8}$`)
			return reg.MatchString(value)
		} else {
			reg := regexp.MustCompile(`^(\+?212|0)5\d{8}$`)
			return reg.MatchString(value)
		}

		reg := regexp.MustCompile(`^[a-z]{1,2}\d+$`)
		return reg.MatchString(value)
	} else if validator.Type == PhoneMobile {
		reg := regexp.MustCompile(`^(\+?212|0)(6|7)\d{8}$`)
		return reg.MatchString(value)
	}
	combined := strings.Join(cityMap[strings.ToLower(validator.Location)], "|") // ex:  bes -> ta|tk
	if combined != "" {
		reg := regexp.MustCompile(`^(` + combined + `)\d+$`)
		return reg.MatchString(value)
	}
	return false
}
