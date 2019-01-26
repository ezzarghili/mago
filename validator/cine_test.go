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

func TestMaCINEValidateWrongWithLocation(t *testing.T) {
	correct := []string{"C 3433434", "EL 783", "_-783", " td_783 "}
	validator := CINE{Location: "BES"}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect token `%s`", v)
		}
	}
}
func TestCINEValidateCorrectWithLocation(t *testing.T) {
	correct := []string{"TA 3433434", "TK 783", "_tk-783", " ta_783 "}
	validator := CINE{Location: "BES"}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct token `%s`", v)
		}
	}
}

func TestCINEValidateCorrectWithUndefinedLocation(t *testing.T) {
	token := "3433434"
	validator := CINE{Location: "BfS"}
	if validator.Validate(token) {
		t.Errorf("failed validating undefined Location `%s`", validator.Location)
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
