package validator

import (
	"regexp"
	"strings"
)

// ZipCode options to customize the validator
type ZipCode struct {
	// Location Area   // Provinces and prefectures ISO code
	Spacer string // regex expression
}

// Validate if zip code is correct
func (validator *ZipCode) Validate(zipCode string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(zipCode, spacerReg))
	reg := regexp.MustCompile(`^[0-9]{5}$`)
	return reg.MatchString(value)
}
