package validator

import "regexp"

const spacers = "[-_ ]" // default regex expression

// remove separation characters `spacers` from the value string
func clean(value string, spacers string) string {
	reg := regexp.MustCompile(spacers)
	return string(reg.ReplaceAll([]byte(value), []byte{}))
}
