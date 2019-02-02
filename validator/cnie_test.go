package validator

import "testing"

func TestCNIEValidateCorrectEmptyLocation(t *testing.T) {
	correct := []string{"TA 343434", "T 78309", "_C-798083", " C_109783 "}
	validator := CNIE{}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct token `%s`", v)
		}
	}
}

func TestCNIEValidateWrongEmptyLocation(t *testing.T) {
	incorrect := []string{"3433434", "78323fdf", "_Ccc 783", " Cdf "}
	validator := CNIE{}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect token `%s`", v)
		}
	}
}

func TestCNIEValidateWithSpacers(t *testing.T) {
	correct := []string{"TA*343434", "C.70983", "*C+70983", " C+70983 "}
	validator := CNIE{Spacer: `[*.+ ]`}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct token `%s`", v)
		}
	}
	incorrect := []string{"TA_343434", "C-70983", "_C-70983", " 70983 "}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect token `%s`", v)
		}
	}
}
