package forgery

import (
	"testing"
	"strings"
)

func TestUserName(t *testing.T) {
	result, err := UserName()
	if result == "" {
		t.Errorf("Username() failed. Got user name: %s.", err.Error())
	}
}

func TestTopLevelDomain(t *testing.T) {
	result, err := TopLevelDomain()
	if result == "" {
		t.Errorf("TopLevelDomain() failed. Got top level domain: %s.", err.Error())
	}
}

func TestDomainName(t *testing.T) {
	result, err := DomainName()
	if result == "" {
		t.Errorf("DomainName() failed. Got domain name: %s.", err.Error())
	}
}

func TestEmail(t *testing.T) {
	result, _ := Email()
	if ok := strings.Contains(result, "@"); !ok {
		t.Errorf("Email() failed. Got email: %s.", result)
	}
}

func TestCountryCodeTopLevelDomains(t *testing.T) {
	result, err := CountryCodeTopLevelDomains()
	if result == "" {
		t.Errorf("CountryCodeTopLevelDomains() failed. Got country code top level domains: %s.", err.Error())
	}
}

func TestIpv4(t *testing.T) {
	result, _ := Ipv4()
	if ok := strings.Contains(result, "."); !ok {
		t.Errorf("Ipv4() failed. Got Ipv4: %s.", result)
	}
}