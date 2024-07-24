package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rules := buildRules(buildArgs(os.Args[1:]))
	f, _ := os.Open("/Users/matt.love/dev/word/wordle.txt")
	r := bufio.NewReader(f)

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
