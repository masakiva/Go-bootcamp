package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	nFlag := flag.Bool("n", false, "omits the trailing newline")
	sFlag := flag.String("s", " ", "separator (default: \" \")")
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sFlag))
	if !*nFlag {
		fmt.Println()
	}
}
