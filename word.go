package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NO_MATCH = iota
	YELLOW_MATCH
	GREEN_MATCH
)

// 0 no match
// 1 yellow match
// 2 green match

//grape 00101
//grape 00101 navel 11010

// letter is not at position; from a yellow or gray match
type notAt struct {
	letter   byte
	position int
}

func (n notAt) passes(word string) bool {
	return word[n.position] != n.letter
}

// how many times letter appears in the word, based on the yellow and green
// matches for it in a single guess. if the guess also had a gray match for
// the letter, the word has exactly that many, otherwise at least that many.
type letterCount struct {
	letter byte
	count  int
	exact  bool
}

func (c letterCount) passes(word string) bool {
	n := strings.Count(word, string(c.letter))
	if c.exact {
		return n == c.count
	}
	return n >= c.count
}

// 2 green match
type green struct {
	letter  byte
	poition int
}

func (g green) passes(word string) bool {
	return word[g.poition] == g.letter
}

type rule interface {
	passes(word string) bool // returns true if word passes rule
}

type rules []rule

func (rs rules) passes(word string) bool {
	for _, r := range rs {
		if !r.passes(word) {
			return false
		}
	}
	return true
}

type arg struct {
	letters string
	numbers string
}

func (a arg) toRules() []rule {
	r := make([]rule, 5)
	counts := make(map[byte]int)
	grays := make(map[byte]bool)
	for i := range a.letters {
		letter := a.letters[i]
		switch int(a.numbers[i]) - '0' {
		case NO_MATCH:
			r[i] = notAt{letter, i}
			grays[letter] = true
		case YELLOW_MATCH:
			r[i] = notAt{letter, i}
			counts[letter]++
		case GREEN_MATCH:
			r[i] = green{letter, i}
			counts[letter]++
		}
	}
	for letter, count := range counts {
		r = append(r, letterCount{letter, count, grays[letter]})
	}
	for letter := range grays {
		if counts[letter] == 0 {
			r = append(r, letterCount{letter, 0, true})
		}
	}
	return r
}

func buildRules(args []arg) rules {
	rules := make([]rule, 0, len(args)*5)
	for _, arg := range args {
		rules = append(rules, arg.toRules()...)
	}
	return rules
}

func buildArgs(a []string) []arg {
	n := len(a) / 2

	args := make([]arg, n)

	for i := 0; i < n; i++ {
		args[i].letters = a[2*i]
		args[i].numbers = a[2*i+1]
	}
	return args
}

func word() {
	rules := buildRules(buildArgs(os.Args[1:]))
	r := bufio.NewReader(os.Stdin)
	line, _, err := r.ReadLine()
	for err == nil {
		if rules.passes(string(line)) {
			fmt.Println(string(line))
		}
		line, _, err = r.ReadLine()
	}
}
