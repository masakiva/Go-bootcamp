package main

import (
	"fmt"
	"os"
	"strconv"
)

const ERROR_MSG string = "Arguments is invalid."

func calculationStr(args []string) (string, bool) {
	var s string

	if len(args) != 3 {
		return "", false
	}
	a, err := strconv.Atoi(args[1])
	if err != nil {
		return "", false
	}
	b, err := strconv.Atoi(args[2])
	if err != nil || b == 0 {
		return "", false
	}

	s += "sum: " + strconv.Itoa(a+b) + "\n"
	s += "difference: " + strconv.Itoa(a-b) + "\n"
	s += "product: " + strconv.Itoa(a*b) + "\n"
	s += "quotient: " + strconv.Itoa(a/b) + "\n"

	return s, true
}

func main() {
	s, ok := calculationStr(os.Args)
	if !ok {
		fmt.Println(ERROR_MSG)
		os.Exit(1)
	}
	fmt.Print(s)
}
