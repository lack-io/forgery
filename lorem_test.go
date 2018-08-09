package forgery

import (
	"testing"
	"strings"
)

func TestWords(t *testing.T) {
	result, err := Words()
	if result == "" {
		t.Errorf("Words() failed. Got word: %s.", err.Error())
	}
	if len(word) < 2 {
		t.Errorf("Words() failed. the words is empty.")
	}
}

func TestWord(t *testing.T) {
	result, _ := Word()
	if len(result) < 1 {
		t.Errorf("Word() failed. Got word: %s.", result)
	}
}

func TestTitle(t *testing.T) {
	result, _ := Title()
	if !strings.Contains(result, "?") &&
		!strings.Contains(result, ".") &&
		!strings.Contains(result, "!") {
		t.Errorf("Title() failed. Got title: %s.", result)
	}
}

func TestSentences(t *testing.T) {
	result, _ := Sentences()
	if !strings.Contains(result, ".") {
		t.Errorf("Sentences() failed. Got sentences: %s.", result)
	}
}

func TestSentence(t *testing.T) {
	result, _ := Sentences()
	if !strings.Contains(result, ".") {
		t.Errorf("Sentence() failed. Got sentence: %s.", result)
	}
}


func TestParagraphs(t *testing.T) {
	result, _ := Paragraphs()
	if !strings.Contains(result, "\n\n") {
		t.Errorf("Paragraphs() failed. Got paragraphs: %s.", result)
	}
}

func TestParagraph(t *testing.T) {
	result, _ := Paragraph()
	if !strings.Contains(result, "\n\n") {
		t.Errorf("Paragraph() failed. Got paragraph: %s.", result)
	}
}