package forgery

import "testing"

func TestHexColor(t *testing.T) {
	result, _ := HexColor()
	if len(result) < 6 {
		t.Errorf("HexColor() failed. Got hex color: %s.", result)
	}
}


func TestHexColorShort(t *testing.T) {
	result, _ := HexColorShort()
	if len(result) < 3 {
		t.Errorf("HexColorShort() failed. Got short hex color: %s", result)
	}
}

func TestText(t *testing.T) {
	result, _ := Text()
	if len(result) < 15 || len(result) > 20 {
		t.Errorf("Text() failed. Got text: %s.", result)
	}
}