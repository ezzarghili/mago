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

// Phone options to customize the validator
type Phone struct {
	Spacer string    // regex expression
	Type   PhoneType // defaults to 0 == PhoneAll
}

// Validate if phone number is correct
func (validator *Phone) Validate(phone string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(phone, spacerReg))
	if validator.Type == PhoneLandLine {
		reg := regexp.MustCompile(`^((\+|00)?212)?0?5\d{8}$`)
		return reg.MatchString(value)
	} else if validator.Type == PhoneMobile {
		reg := regexp.MustCompile(`^((\+|00)?212)?0?(6|7)\d{8}$`)
		return reg.MatchString(value)
	}
	reg := regexp.MustCompile(`^((\+|00)?212)?0?(5|6|7)\d{8}$`)
	return reg.MatchString(value)
}
