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

// counts each letter once per word, so counts are the number of words that
// contain the letter. anything other than a-z is ignored.
func countWord(counts []entry, word string) {
	var seen [26]bool
	for _, b := range word {
		i := b - 'a'
		if i < 0 || i >= 26 || seen[i] {
			continue
		}
		seen[i] = true
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
	// same letters (e.g. anagrams like least and slate), fall back to alphabetical
	return first < second
}

func guess() {
	words := make([]string, 0)
	counts := make([]entry, 26)
	for i := range counts {
		counts[i].letter = 'a' + rune(i)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		words = append(words, s.Text())
		countWord(counts, s.Text())
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
