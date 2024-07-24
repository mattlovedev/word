package main

import (
	"bufio"
	"fmt"
	"os"
)

func filter_five() {
	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	for err == nil {
		if len(string(line)) == 5 {
			fmt.Println(string(line))
		}
		line, _, err = r.ReadLine()
	}
}
