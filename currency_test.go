package forgery

import "testing"

func TestDescription(t *testing.T) {
	result, err := Description()
	if result == "" {
		t.Errorf("Description() failed. Got description: %s.", err.Error())
	}
}

func TestCode(t *testing.T) {
	result, err := Code()
	if result == "" {
		t.Errorf("Code() failed. Got code: %s.", err.Error())
	}
}
