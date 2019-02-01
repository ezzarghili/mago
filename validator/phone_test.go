package validator

import "testing"

func TestPhoneValidateCorrectPhone(t *testing.T) {
	correct := []string{
		"+2120787766543",
		"+212787766543",
		"212787766543",
		"0787766543",
		"0587766543",
		"0687766543",
		"06-8776-6543",
		"6-8776-6543",
		"06 87 76 65 43",
		"00212787766543"}
	validator := Phone{}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct number `%s`", v)
		}
	}
}
func TestPhoneValidateIncorrectPhone(t *testing.T) {
	correct := []string{
		"+21207877665456",
		"0887766543",
		"(0687766543)"}
	validator := Phone{}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating correct number `%s`", v)
		}
	}
}

func TestPhoneValidateCorrectPhoneMobile(t *testing.T) {
	correct := []string{
		"+2120787766543",
		"+212787766543",
		"212787766543",
		"0787766543",
		"0687766543",
		"06-8776-6543",
		"6-8776-6543",
		"06 87 76 65 43",
		"00212787766543"}
	validator := Phone{Type: PhoneMobile}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct number `%s`", v)
		}
	}
}

func TestPhoneValidateIncorrectPhoneMobile(t *testing.T) {
	correct := []string{
		"+2120587766543",
		"+212587766543",
		"212587766543",
		"0587 766543",
		"05877_665 43",
		"0587766 543",
		"05-8776-6543",
		"5-8776-6543",
		"05 87 76 65 43",
		"09 87 76 65 43",
		"00212587766543"}
	validator := Phone{Type: PhoneMobile}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating correct number `%s`", v)
		}
	}
}

func TestPhoneValidateWithSpacers(t *testing.T) {
	correct := []string{
		"+21+207*87766 543",
		"+2127.87766543",
		"212787766543",
		"07.87766543",
		"0587766543",
		"0687766543",
	}
	validator := Phone{Spacer: `[*.+ ]`}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct spacers `%s`", v)
		}
	}
	incorrect := []string{
		"+21207_87766543",
		"+212787766543_",
		"212787-766543",
	}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect spacers `%s`", v)
		}
	}
}

func TestPhoneValidateCorrectPhoneLandline(t *testing.T) {
	correct := []string{
		"+2120523298543",
		"+212523298653",
		"212523296543",
		"0523 296543",
		"0523_295 943",
		"0523296 543",
		"05-2329-6543",
		"5-2329-6543",
		"05 23 29 65 43",
		"00212523296543"}
	validator := Phone{Type: PhoneLandLine}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct number `%s`", v)
		}
	}
}

func TestPhoneValidateIncorrectPhoneLandline(t *testing.T) {
	correct := []string{
		"+2120787766543",
		"+212787766543",
		"212787766543",
		"0787766543",
		"0687766543",
		"06-8776-6543",
		"6-8776-6543",
		"06 87 76 65 43",
		"00212787766543"}
	validator := Phone{Type: PhoneLandLine}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect number `%s`", v)
		}
	}

}
