package main

import (
	"fmt"
	"os"
	"strings"
)

var kitaplar = []string{
	"Lord Of The Rings",
	"Harry Potter",
	"Dune",
	"Witcher"}

func List() {
	for _, b := range kitaplar {
		fmt.Println(b)
	}
}
func Search(k string) {
	var v string
	succ := 0 //succ for checking if we have the book, if it stays at 0 that means we don't have it.
	for _, v = range kitaplar {
		if strings.EqualFold(k, v) { // strings.EqualFold for case insensitive search.
			fmt.Printf("Success, we have %s", v)
			succ++
		}
	}
	if succ == 0 {
		fmt.Printf("Sorry, there is no such a book named: %s", k)
	}
}

func main() {

	if len(os.Args) <= 1 { // Checking if user entered a command and printing a guide message.
		panic("Error!! You did not enter any command.Please enter -list command for listing, -search <bookname> for searching books.")
	}

	switch os.Args[1] {
	case "list":
		List()
	case "search":
		arg := strings.Join(os.Args[2:], " ") // Use strings.Join for checking multi word books. eg: Lord of the rings
		Search(arg)
	default:
		fmt.Println("Error!! Enter -list command for listing, -search <bookname> for searching books.")
	}
}
