// Page for word.js. Guesses are added from the form, tiles are tapped to cycle through
// gray, yellow and green, and the suggestions update on every change. The guesses are
// mirrored in the URL hash (e.g. #crane01020lions00211) so a game can be reloaded or shared.

const SHOWN = 10 // same as the Go program's output

let guesses = []

const form = document.getElementById("guessForm")
const input = document.getElementById("guessInput")
const error = document.getElementById("error")
const board = document.getElementById("board")
const remaining = document.getElementById("remaining")
const suggestions = document.getElementById("suggestions")
const reset = document.getElementById("reset")

function addGuess(letters) {
    letters = letters.trim().toLowerCase()
    const message = validateGuess(letters, "00000")
    if (message) {
        error.textContent = message
        return false
    }
    error.textContent = ""
    guesses.push({ letters, numbers: "00000" })
    update()
    return true
}

function cycleTile(row, i) {
    const g = guesses[row]
    const next = (Number(g.numbers[i]) + 1) % 3
    g.numbers = g.numbers.slice(0, i) + next + g.numbers.slice(i + 1)
    update()
}

function removeGuess(row) {
    guesses.splice(row, 1)
    update()
}

function renderBoard() {
    board.innerHTML = ""
    guesses.forEach((g, row) => {
        const rowDiv = document.createElement("div")
        rowDiv.className = "row"
        for (let i = 0; i < 5; i++) {
            const tile = document.createElement("button")
            tile.type = "button"
            tile.className = `tile match${g.numbers[i]}`
            tile.textContent = g.letters[i]
            tile.title = "Tap to change color"
            tile.addEventListener("click", () => cycleTile(row, i))
            rowDiv.appendChild(tile)
        }
        const remove = document.createElement("button")
        remove.type = "button"
        remove.className = "remove"
        remove.textContent = "×"
        remove.title = "Remove guess"
        remove.setAttribute("aria-label", `Remove ${g.letters}`)
        remove.addEventListener("click", () => removeGuess(row))
        rowDiv.appendChild(remove)
        board.appendChild(rowDiv)
    })
    reset.hidden = guesses.length === 0
}

function renderSuggestions() {
    const words = rankWords(guesses, wordleList)
    remaining.textContent = words.length === 1 ? "1 possible word" : `${words.length} possible words`
    suggestions.innerHTML = ""
    for (const word of words.slice(0, SHOWN)) {
        const b = document.createElement("button")
        b.type = "button"
        b.className = "suggestion"
        b.textContent = word
        b.title = "Add as a guess"
        b.addEventListener("click", () => addGuess(word))
        suggestions.appendChild(b)
    }
}

function writeHash() {
    const hash = guesses.map(g => g.letters + g.numbers).join("")
    history.replaceState(null, "", hash ? `#${hash}` : location.pathname + location.search)
}

// ignores the hash if any guess in it is invalid
function readHash() {
    const hash = location.hash.slice(1).toLowerCase()
    const parsed = []
    for (let i = 0; i + 10 <= hash.length; i += 10) {
        const letters = hash.slice(i, i + 5)
        const numbers = hash.slice(i + 5, i + 10)
        if (validateGuess(letters, numbers)) {
            return []
        }
        parsed.push({ letters, numbers })
    }
    return hash.length % 10 === 0 ? parsed : []
}

function update() {
    renderBoard()
    renderSuggestions()
    writeHash()
}

form.addEventListener("submit", e => {
    e.preventDefault()
    if (addGuess(input.value)) {
        input.value = ""
    }
})

reset.addEventListener("click", () => {
    guesses = []
    error.textContent = ""
    update()
    input.focus()
})

window.addEventListener("hashchange", () => {
    guesses = readHash()
    update()
})

guesses = readHash()
update()
