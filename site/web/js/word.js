// Filtering and ranking, mirroring the Go program in legacy/. A guess is a 5 letter word
// and a result of 5 digits: 0 gray, 1 yellow, 2 green, e.g. crane 01020.

const NO_MATCH = 0
const YELLOW_MATCH = 1
const GREEN_MATCH = 2

// letter is not at position; from a yellow or gray match
const notAt = (letter, position) => word => word[position] !== letter

// letter is at position; from a green match
const green = (letter, position) => word => word[position] === letter

// how many times letter appears in the word, based on the yellow and green
// matches for it in a single guess. if the guess also had a gray match for
// the letter, the word has exactly that many, otherwise at least that many.
const letterCount = (letter, count, exact) => word => {
    const n = word.split(letter).length - 1
    return exact ? n === count : n >= count
}

function toRules(letters, numbers) {
    const rules = []
    const counts = {}
    const grays = {}
    for (let i = 0; i < letters.length; i++) {
        const letter = letters[i]
        switch (Number(numbers[i])) {
            case NO_MATCH:
                rules.push(notAt(letter, i))
                grays[letter] = true
                break
            case YELLOW_MATCH:
                rules.push(notAt(letter, i))
                counts[letter] = (counts[letter] || 0) + 1
                break
            case GREEN_MATCH:
                rules.push(green(letter, i))
                counts[letter] = (counts[letter] || 0) + 1
                break
        }
    }
    for (const letter in counts) {
        rules.push(letterCount(letter, counts[letter], !!grays[letter]))
    }
    for (const letter in grays) {
        if (!counts[letter]) {
            rules.push(letterCount(letter, 0, true))
        }
    }
    return rules
}

// guesses is a list of { letters, numbers }
function buildRules(guesses) {
    return guesses.flatMap(g => toRules(g.letters, g.numbers))
}

const passes = (rules, word) => rules.every(rule => rule(word))

// returns an error message, or "" if the guess and result are valid
function validateGuess(letters, numbers) {
    if (!/^[a-z]{5}$/.test(letters)) {
        return `guess "${letters}" must be 5 letters`
    }
    if (!/^[012]{5}$/.test(numbers)) {
        return `result "${numbers}" for guess "${letters}" must be 5 digits of 0, 1 or 2`
    }
    return ""
}

// counts each letter once per word, so counts are the number of words that
// contain the letter. anything other than a-z is ignored.
function countWord(counts, word) {
    const seen = {}
    for (const letter of word) {
        if (letter in counts && !seen[letter]) {
            seen[letter] = true
            counts[letter]++
        }
    }
}

// letters ordered from most to fewest words containing them
function letterOrder(words) {
    const counts = {}
    for (let i = 0; i < 26; i++) {
        counts[String.fromCharCode(97 + i)] = 0
    }
    for (const word of words) {
        countWord(counts, word)
    }
    return Object.keys(counts).sort((a, b) => counts[b] - counts[a])
}

// negative if first is better than second: whichever has the most common letter
// the other lacks wins. same letters (e.g. anagrams like least and slate), fall
// back to alphabetical
function cmpWords(order, first, second) {
    for (const letter of order) {
        const firstContains = first.includes(letter)
        const secondContains = second.includes(letter)
        if (firstContains && !secondContains) {
            return -1
        } else if (secondContains && !firstContains) {
            return 1
        }
    }
    return first < second ? -1 : first > second ? 1 : 0
}

// the words that fit every guess, best first
function rankWords(guesses, wordList) {
    const rules = buildRules(guesses)
    const words = wordList.filter(word => passes(rules, word))
    const order = letterOrder(words)
    return words.sort((a, b) => cmpWords(order, a, b))
}
