package forgery

import (
	"strings"
	"math/rand"
	"time"
)

var word []string = make([]string, 0)

func sample(words []string, quantity int) (result []string) {
	length := len(words)
	for i := 0; i < quantity; i++ {
		rand.Seed(time.Now().UnixNano())
		rd := rand.Intn(length)
		result = append(result, words[rd:rd+1]...)
	}
	return
}

// Random words
func Words() (string, error) {
	if len(word) < 2 {
		slice, err := loader("lorem_ipsum")
		checkErr(err)
		w := strings.Join(slice, " ")
		w = strings.Replace(w, "\n", "", -1)
		if strings.Contains(w, ".") {
			w = strings.Replace(w, ".", "", -1)
		}
		if strings.Contains(w, "|") {
			w = strings.Replace(w, "|", "", -1)
		}
		if strings.Contains(w, ",") {
			w = strings.Replace(w, ",", "", -1)
		}
		if strings.Contains(w, ";") {
			w = strings.Replace(w, ";", "", -1)
		}
		if strings.Contains(w, "/") {
			w = strings.Replace(w, "/", "", -1)
		}
		word = strings.Split(w, " ")
	}
	return strings.Join(sample(word, 10), " "), nil
}

// Random word
func Word() (string, error) {
	w, _ := Words()
	return strings.Join(sample(strings.Split(w, " "), 1), " "), nil
}

// Random sentence to be used as e.g. an e-mail subject.
func Title() (string, error) {
	w, _ := Words()
	result := strings.Join(sample(strings.Split(w, " "), 4), " ")
	result += strings.Join(sample([]string{"?", ".", "!"}, 1), "")
	return result, nil
}

// Random sentences.
func Sentences() (string, error) {
	slice, err := loader("lorem_ipsum")
	checkErr(err)
	return strings.Replace(
		strings.Join(sample(slice, 2), " "), "\n", "", -1), nil
}

// Random sentence.
func Sentence() (string, error) {
	sen, _ := Sentences()
	return strings.Join(sample(strings.Split(sen, " "), 1), " "), nil
}

// Random paragraphs.
func Paragraphs() (string, error) {
	sen, _ := Sentences()
	return sen + "\n\n", nil
}

// Random paragraph.
func Paragraph() (string, error) {
	sen, _ := Sentence()
	return sen + "\n\n", nil
}
