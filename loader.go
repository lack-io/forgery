package forgery

import (
	"os"
	"io"
	"bufio"
	"math/rand"
	"time"
	"strings"
)

var dictionaries map[string][]string = make(map[string][]string)

func random(slice []string) string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(slice))
	return strings.TrimSpace(slice[n])
}

func loader(name string) (slice []string, err error) {
	slice, ok := dictionaries[name]
	if ok {
		return slice, nil
	}
	file, err := os.Open("./dictionaries/" + name)
	if err != nil {
		return slice, err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	for {
		line, err := rd.ReadString('\n')
		slice = append(slice, line)
		if err != nil || io.EOF == err {
			break
		}
	}
	dictionaries[name] = slice
	return slice, nil
}

func checkErr(err error) (string, error) {
	return "", err
}