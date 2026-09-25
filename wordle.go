package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed wordle.txt
var wordleList string

func main() {
	rules := rulesFromCommandLine()
	r := bufio.NewReader(strings.NewReader(wordleList))

	words := make([]string, 0)
	counts := make([]entry, 26)
	for i := range counts {
		counts[i].letter = 'a' + rune(i)
	}

	line, _, err := r.ReadLine()
	for err == nil {
		if rules.passes(string(line)) {
			words = append(words, string(line))
			countWord(counts, string(line))
		}
		line, _, err = r.ReadLine()
	}

	sort.Slice(counts, func(i int, j int) bool {
		return counts[i].count > counts[j].count
	})

	sort.Slice(words, func(i int, j int) bool {
		return cmpWords(counts, words[i], words[j])
	})

	for i := 0; i < 10 && i < len(words); i++ {
		fmt.Println(words[i])
	}
}
