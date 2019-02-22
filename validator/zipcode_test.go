package validator

import "testing"

func TestZipCodeValidateCorrect(t *testing.T) {
	correct := []string{
		"13000",
		"13 000",
		"25_998"}
	validator := ZipCode{}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct zip code `%s`", v)
		}
	}
}
func TestZipCodeValidateIncorrect(t *testing.T) {
	correct := []string{
		"879",
		"df879",
		"(98090)"}
	validator := ZipCode{}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect zip code `%s`", v)
		}
	}
}

func TestZipCodeValidateStrictModeIncorrect(t *testing.T) {
	correct := []string{
		"11011",
		"13004",
		"16005"}
	validator := ZipCode{StrictMode: true}
	for _, v := range correct {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect zip code `%s`", v)
		}
	}
}

func TestZipCodeValidateStrictModeCorrect(t *testing.T) {
	correct := []string{
		"11022",
		"13003",
		"16004"}
	validator := ZipCode{StrictMode: true}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct zip code `%s`", v)
		}
	}
}

func TestZipCodeValidateWithSpacers(t *testing.T) {
	correct := []string{
		"13*000",
		"13+000",
		"25.998"}
	validator := ZipCode{Spacer: `[*.+ ]`}
	for _, v := range correct {
		if !validator.Validate(v) {
			t.Errorf("failed validating correct spacers `%s`", v)
		}
	}
	incorrect := []string{"13-000",
		"13-000",
		"25_998"}
	for _, v := range incorrect {
		if validator.Validate(v) {
			t.Errorf("failed validating incorrect spacers `%s`", v)
		}
	}
}
