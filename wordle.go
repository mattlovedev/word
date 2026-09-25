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
	s := bufio.NewScanner(strings.NewReader(wordleList))

	words := make([]string, 0)
	counts := make([]entry, 26)
	for i := range counts {
		counts[i].letter = 'a' + rune(i)
	}

	for s.Scan() {
		if rules.passes(s.Text()) {
			words = append(words, s.Text())
			countWord(counts, s.Text())
		}
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
