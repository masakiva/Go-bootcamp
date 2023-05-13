package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	asterisk_max_nb, err := strconv.Atoi(os.Args[1])
	line_asterisk_nb := 1

	if err != nil {
		fmt.Println("strconv.Atoi(): error")
		return
	}

	for asterisk_max_nb > 0 {
		if line_asterisk_nb <= asterisk_max_nb {
			for i := 0; i < line_asterisk_nb; i++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
		asterisk_max_nb -= line_asterisk_nb
		line_asterisk_nb++
	}
}
