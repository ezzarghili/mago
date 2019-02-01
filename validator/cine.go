package validator

import (
	"regexp"
	"strings"
)

// CINE ID validation
type CINE struct {
	Spacer string // regex expression
}

// Validate if CINE code is correct
func (validator *CINE) Validate(cine string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(cine, spacerReg))
	reg := regexp.MustCompile(`^[a-z]{1,2}\d+$`)
	return reg.MatchString(value)
}
