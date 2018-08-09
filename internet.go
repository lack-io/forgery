package forgery

import (
	"strings"
	"math/rand"
	"time"
	"strconv"
)

// Random user name
func UserName() (string, error) {
	result, err := FirstName()
	checkErr(err)
	return strings.ToLower(result), nil
}

// Random top level domain
func TopLevelDomain() (string, error) {
	slice, err := loader("top_level_domains")
	checkErr(err)
	return random(slice), nil
}

// Random domain name
func DomainName() (string, error) {
	slice, err := loader("company_names")
	checkErr(err)
	tld, err := TopLevelDomain()
	checkErr(err)
	result := random(slice) + "." + tld
	return strings.ToLower(result), nil
}

// Random email address
func Email() (string, error) {
	user, err := UserName()
	checkErr(err)
	domain, err := DomainName()
	return user + "@" + domain, err
}

// Random country code top level domains
func CountryCodeTopLevelDomains() (string, error) {
	slice, err := loader("country_code_top_level_domains")
	checkErr(err)
	return random(slice), nil
}

// Random IPv4 address
func Ipv4() (string, error) {
	ip := make([]string, 0)
	for i := 0; i < 4; i++ {
		rand.Seed(time.Now().UnixNano())
		ip = append(ip, strconv.Itoa(rand.Intn(255)))
	}
	return strings.Join(ip, "."), nil
}