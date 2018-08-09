package forgery

import "testing"

func Testrandom(t testing.T) {
	slice, _ := loader("male_first_names")
	out := random(slice)
	out1 := random(slice)
	if out == "" || out1 == "" {
		t.Errorf("random failed. Got random value: %s.", out)
	}
	if out1 == out {
		t.Errorf("random failed. Got some value.")
	}
}

func Testloader(t testing.T) {
	slice, _ := loader("male_first_names")
	if len(slice) == 0 {
		t.Errorf("loader failed. Got empty value.")
	}
	if _, ok := dictionaries["male_first_names"]; !ok {
		t.Errorf("loader failed. Can not cache value.")
	}
}
