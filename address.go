package forgery

import (
	"math/rand"
	"time"
	"strconv"
)

// Random street name.
func StreetName() (string, error) {
	slice, err := loader("street_names")
	checkErr(err)
	return random(slice), nil
}

// Random street number.
func StreetNumber() (string, error) {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(2000)), nil
}

// Random street suffix.
func StreetSuffix() (string, error) {
	slice, err := loader("street_suffixes")
	checkErr(err)
	return random(slice), nil
}

// Random street address.
// Equivalent of `StreetNumber() + StreetName() + StreetSuffix()`
func StreetAddress() (string, error) {
	streetNumber, err := StreetNumber()
	checkErr(err)
	streetName, err := StreetName()
	checkErr(err)
	streetSuffix, err := StreetSuffix()
	checkErr(err)
	return streetNumber + " " + streetName + " " + streetSuffix, nil
}

// Random city name.
func City() (string, error) {
	slice, err := loader("cities")
	checkErr(err)
	return random(slice), nil
}

// Random US state name.
func State() (string, error) {
	slice, err := loader("states")
	checkErr(err)
	return random(slice), nil
}

// Random US abbreviated state name.
func StateAbbr() (string, error) {
	slice, err := loader("state_abbrevs")
	checkErr(err)
	return random(slice), nil
}

// Random ZIP code, either in `#####` or `#####-####` format.
func ZipCode() (string, error) {
	rand.Seed(time.Now().UnixNano())
	format := "#####-####"
	result := ""
	for i := 0; i < len(format); i++ {
		if format[i:i+1] == "#" {
			result += strconv.Itoa(rand.Intn(10))
		} else {
			result += format[i:i+1]
		}
	}
	return result, nil
}

// Random phone number in `#-(###)###-####` format.
func Phone() (string, error) {
	rand.Seed(time.Now().UnixNano())
	format := "###-####-####"
	result := ""
	for i := 0; i < len(format); i++ {
		if format[i:i+1] == "#" {
			result += strconv.Itoa(rand.Intn(10))
		} else {
			result += format[i:i+1]
		}
	}
	return result, nil
}

// Random country name.
func Country() (string, error) {
	slice, err := loader("countries")
	checkErr(err)
	return random(slice), nil
}

// Random continent name.
func Continent() (string, error) {
	slice, err := loader("continents")
	checkErr(err)
	return random(slice), nil
}