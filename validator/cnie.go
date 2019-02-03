package validator

import (
	"regexp"
	"strings"
)

// CNIE ID validation
type CNIE struct {
	Spacer string // regex expression
}

// Validate if CNIE code is correct
func (validator *CNIE) Validate(CNIE string) bool {
	var spacerReg string
	if validator.Spacer != "" {
		spacerReg = validator.Spacer
	} else {
		spacerReg = spacers
	}
	value := strings.ToLower(clean(CNIE, spacerReg))
	reg := regexp.MustCompile(`^[a-z]{1,2}\d{5,6}$`)
	return reg.MatchString(value)
}
