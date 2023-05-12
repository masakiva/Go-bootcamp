package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid argument")
	}
	for _, arg := range os.Args[1:] {
		match, _ := regexp.MatchString(`[\w\-\._]+@[\w\-\._]+\.[A-Za-z]+`, arg)
		if match && len(arg) <= 256 {
			fmt.Println(arg, "is a valid email address.")
		} else {
			fmt.Println(arg, "is NOT a valid email address.")
		}
	}
}
