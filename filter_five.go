package main

import (
	"bufio"
	"fmt"
	"os"
)

func filter_five() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if len(s.Text()) == 5 {
			fmt.Println(s.Text())
		}
	}
}
