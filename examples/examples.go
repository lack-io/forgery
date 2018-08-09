package main

import (
	"xingyys/forgery"
	"fmt"
)

func main() {
	name, _ := forgery.Sentences()
	fmt.Println(name)
}
