package forgery


// Random currency description, e.g. `United Kingdom Pounds`.
func Description() (string, error) {
	slice, err := loader("currency_descriptions")
	checkErr(err)
	return random(slice), nil
}

// Random currency code, e.g. `GBP`
func Code() (string, error) {
	slice, err := loader("currency_codes")
	checkErr(err)
	return random(slice), nil
}
