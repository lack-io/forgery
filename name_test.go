package forgery

import (
	"testing"
	"strings"
)

func TestFirstName(t *testing.T) {
	firstName, err := FirstName()
	if firstName == "" {
		t.Errorf("FirstName() failed. Got firstname: %s.", err.Error())
	}
}

func TestLastName(t *testing.T) {
	lastname, err := LastName()
	if lastname == "" {
		t.Errorf("LastName() failed. Got firstname: %s.", err.Error())
	}
}

func TestFullName(t *testing.T) {
	name, err := FullName()
	if name == "" {
		t.Errorf("FullName() failed. Got full name: %s.", err.Error())
	}
	nameSlice := strings.Split(name, " ")
	if len(nameSlice) < 2 {
		t.Errorf("FullName() failed. Error format: %s.", name)
	}
}

func TestMaleFirstName(t *testing.T) {
	name, err := MaleFirstName()
	if name == "" {
		t.Errorf("MaleFirstName() failed. Got male first name: %s.", err.Error())
	}
}

func TestFemaleFirstName(t *testing.T) {
	name, err := FemaleFirstName()
	if name == "" {
		t.Errorf("FemaleFirstName() failed. Got female first name: %s.", err.Error())
	}
}

func TestCompanyName(t *testing.T) {
	name, err := CompanyName()
	if name == "" {
		t.Errorf("CompanyName() failed. Got company name: %s.", err.Error())
	}
}

func TestJobTitle(t *testing.T) {
	name, err := JobTitle()
	if name == "" {
		t.Errorf("JobTitle() failed. Got job title: %s.", err.Error())
	}
}

func TestJobTitleSuffix(t *testing.T) {
	name, err := JobTitleSuffix()
	if name == "" {
		t.Errorf("JobTitleSuffix() failed. Got job title suffix: %s", err.Error())
	}
}

func TestNameTitle(t *testing.T) {
	name, err := Title()
	if name == "" {
		t.Errorf("NameTitle() failed. Got title: %s.", err.Error())
	}
}

func TestSuffix(t *testing.T) {
	name, err := Suffix()
	if name == "" {
		t.Errorf("Suffix() failed. Got title: %s.", err.Error())
	}
}

func TestLocation(t *testing.T) {
	name, err := Location()
	if name == "" {
		t.Errorf("Location() failed. Got location: %s.", err.Error())
	}
}

func TestIndustry(t *testing.T) {
	name, err := Industry()
	if name == "" {
		t.Errorf("Industry() failed. Got industry: %s.", err.Error())
	}
}
