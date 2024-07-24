package main

import (
	"bufio"
	"fmt"
	"os"
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

// 0 no match
type excludes struct {
	letter byte
}

func (e excludes) passes(word string) bool {
	for i := range word {
		if word[i] == e.letter {
			return false
		}
	}
	return true
}

// 1 yellow match
type yellow struct {
	letter   byte
	position int
}

func (y yellow) passes(word string) bool {
	if word[y.position] == y.letter {
		return false
	}
	for i := range word {
		if word[i] == y.letter {
			return true
		}
	}
	return false
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
	for i := range a.letters {
		switch int(a.numbers[i]) - '0' {
		case NO_MATCH:
			r[i] = excludes{a.letters[i]}
		case YELLOW_MATCH:
			r[i] = yellow{a.letters[i], i}
		case GREEN_MATCH:
			r[i] = green{a.letters[i], i}
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

// TODO there is a bug where same letter has an exclude and not exclude, need to ignore the exclude or turn it yellow

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
