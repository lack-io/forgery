package forgery

import "strings"

// Random male of female first name.
func FirstName() (string, error) {
	slice, err := loader("male_first_names")
	checkErr(err)
	slice1, err := loader("female_first_names")
	checkErr(err)
	slice = append(slice, slice1...)
	return random(slice), nil
}

// Random last name.
func LastName() (string, error) {
	slice, err := loader("last_names")
	checkErr(err)
	return random(slice), nil
}

// Random full name. Equivalent of `FirstName() + " " LastName()`
func FullName() (string, error) {
	firstname, err := FirstName()
	checkErr(err)
	lastname, err := LastName()
	checkErr(err)
	return firstname + " " + lastname, nil
}

// Random male first name.
func MaleFirstName() (string, error) {
	slice, err := loader("male_first_names")
	checkErr(err)
	return random(slice), nil
}

// Random female first name.
func FemaleFirstName() (string, error) {
	slice, err := loader("female_first_names")
	checkErr(err)
	return random(slice), nil
}

// Random company name.
func CompanyName() (string, error) {
	slice, err := loader("company_names")
	checkErr(err)
	return random(slice), nil
}

// Random job title.
func JobTitle() (string, error) {
	slice, err := loader("job_titles")
	checkErr(err)
	result := random(slice)
	suffix, _ := JobTitleSuffix()
	result = strings.Replace(result, "#{N}", suffix, -1)
	return result, nil
}

// Random job title suffix
func JobTitleSuffix() (string, error) {
	slice, err := loader("job_title_suffixes")
	checkErr(err)
	return random(slice), nil
}

// Random name title, e.g. `Mr`.
func NameTitle() (string, error) {
	slice, err := loader("name_titles")
	checkErr(err)
	return random(slice), nil
}

// Random name suffix, e.g. `Jr`
func Suffix() (string, error) {
	slice, err := loader("name_suffixes")
	checkErr(err)
	return random(slice), nil
}

// Random location name, e.g. `M16 Headquarters`
func Location() (string, error) {
	slice, err := loader("locations")
	checkErr(err)
	return random(slice), nil
}

// Random industry name.
func Industry() (string, error) {
	slice, err := loader("industries")
	checkErr(err)
	return random(slice), nil
}
