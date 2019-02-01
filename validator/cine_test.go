package validator

import "testing"

func TestCINEValidateCorrectEmptyLocation(t *testing.T) {
	correct := []string{"TA 3433434", "C 783", "_C-783", " C_783 "}
	validator := CINE{}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct token `%s`", v)
		}
	}
}

func TestCINEValidateWrongEmptyLocation(t *testing.T) {
	incorrect := []string{"3433434", "78323fdf", "_Ccc 783", " Cdf "}
	validator := CINE{}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect token `%s`", v)
		}
	}
}

func TestCINEValidateWithSpacers(t *testing.T) {
	correct := []string{"TA*3433434", "C.783", "*C+783", " C+783 "}
	validator := CINE{Spacer: `[*.+ ]`}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct token `%s`", v)
		}
	}
	incorrect := []string{"TA_3433434", "C-783", "_C-783", " C_783 "}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect token `%s`", v)
		}
	}
}
