package forgery

import "testing"

func TestGender(t *testing.T) {
	result, err := Gender()
	if result == "" {
		t.Errorf("Gender() failed. Got gender: %s.", err.Error())
	}
}

func TestGenderAbbr(t *testing.T) {
	result, err := GenderAbbr()
	if result == "" {
		t.Errorf("GenderAbbr() failed. Got abbreviated gender: %s.", err.Error())
	}
}

func TestShirtSize(t *testing.T) {
	result, err := ShirtSize()
	if result == "" {
		t.Errorf("ShirtSize() failed. Got shirt size: %s.", err.Error())
	}
}

func TestRace(t *testing.T) {
	result, err := Race()
	if result == "" {
		t.Errorf("Race() failed. Got race: %s.", err.Error())
	}
}

func TestLanguage(t *testing.T) {
	result, err := Language()
	if result == "" {
		t.Errorf("Language() failed. Got language: %s.", err.Error())
	}
}
