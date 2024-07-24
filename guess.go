package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type entry struct {
	letter rune
	count  int
	//rank   int // sorted by count, 0 being most 25 being least
}

func countWord(counts []entry, word string) {
	for _, b := range word {
		i := b - 'a'
		counts[i].count = counts[i].count + 1
	}
}

func wordContains(word string, char rune) bool {
	for _, c := range word {
		if c == char {
			return true
		}
	}
	return false
}

// return true if first is better than second
func cmpWords(counts []entry, first string, second string) bool {
	for _, count := range counts {
		firstContains := wordContains(first, count.letter)
		secondContains := wordContains(second, count.letter)
		if firstContains && !secondContains {
			return true
		} else if secondContains && !firstContains {
			return false
		}
	}
	// unreachable
	return true
}

func guess() {
	words := make([]string, 0)
	counts := make([]entry, 26)
	for i := range counts {
		counts[i].letter = 'a' + rune(i)
	}

	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	for err == nil {
		words = append(words, string(line))
		countWord(counts, string(line))
		line, _, err = r.ReadLine()
	}

	sort.Slice(counts, func(i int, j int) bool {
		return counts[i].count > counts[j].count
	})

	sort.Slice(words, func(i int, j int) bool {
		return cmpWords(counts, words[i], words[j])
	})

	for _, word := range words {
		fmt.Println(word)
	}

}
