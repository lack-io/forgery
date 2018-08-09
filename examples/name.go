package main

import (
	"github.com/xingyys/forgery"
	"fmt"
)

func main() {
	firstname, _ := forgery.FirstName()
	fmt.Printf("Get the random firstname: %s.\n", firstname)
	lastname, _ := forgery.LastName()
	fmt.Printf("Get the random lastname: %s.\n", lastname)
	fullname, _ := forgery.FullName()
	fmt.Printf("Got the random fullname: %s.\n", fullname)
}
