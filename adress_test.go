package forgery

import (
	"testing"
	"reflect"
	"strings"
)

func TestStreetName(t *testing.T) {
	result, err := StreetName()
	if result == "" {
		t.Errorf("StreetName() failed. Got street name: %s.", err.Error())
	}
}

func TestStreetNumber(t *testing.T) {
	result, err := StreetNumber()
	if reflect.TypeOf(result).String() != "string" {
		t.Errorf("StreetNumber() failed. Got street number: %s.", err.Error())
	}
}

func TestStreetSuffix(t *testing.T) {
	result, err := StreetSuffix()
	if result == "" {
		t.Errorf("StreetSuffix() failed. Got street suffix: %s.", err.Error())
	}
}

func TestStreetAddress(t *testing.T) {
	result, _ := StreetAddress()
	if len(strings.Split(result, " ")) < 2 {
		t.Errorf("StreetAddress() failed. Got street address: %s.", result)
	}
}

func TestCity(t *testing.T) {
	result, err := City()
	if result == "" {
		t.Errorf("City() failed. Got ciry: %s.", err.Error())
	}
}

func TestState(t *testing.T) {
	result, err := State()
	if result == "" {
		t.Errorf("State() failed. Got state: %s.", err.Error())
	}
}

func TestStateAbbr(t *testing.T) {
	result, err := State()
	if result == "" {
		t.Errorf("StateAbbr() failed. Got state abbreviated: %s.", err.Error())
	}
}

func TestZipCode(t *testing.T) {
	result, err := ZipCode()
	if len(result) < 10 {
		t.Errorf("ZipCode() failed. Got zip code: %s.", err.Error())
	}
}


func TestPhone(t *testing.T) {
	result, err := Phone()
	if len(result) < 13 {
		t.Errorf("Phone() failed. Got phone: %s.", err.Error())
	}
}

func TestCountry(t *testing.T) {
	result, err := Country()
	if result == "" {
		t.Errorf("Country() failed. Got country: %s.", err.Error())
	}
}

func TestContinent(t *testing.T) {
	result, err := Continent()
	if result == "" {
		t.Errorf("Continent() failed. Got continent: %s.", err.Error())
	}
}
