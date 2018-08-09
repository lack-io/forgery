package forgery

// Random gender.
func Gender() (string, error) {
	slice, err := loader("genders")
	checkErr(err)
	return random(slice), err
}

// Random abbreviated gender.
func GenderAbbr() (string, error) {
	gender, _ := Gender()
	return gender[0:1], nil
}

// Random shirt size.
func ShirtSize() (string, error) {
	slice, err := loader("shirt_sizes")
	checkErr(err)
	return random(slice), nil
}

// Random race.
func Race() (string, error) {
	slice, err := loader("races")
	checkErr(err)
	return random(slice), nil
}

// Random language name. e.g. `Polish`
func Language() (string, error) {
	slice, err := loader("languages")
	checkErr(err)
	return random(slice), nil
}
