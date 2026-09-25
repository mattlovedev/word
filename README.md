# word

A Wordle helper: enter your guesses and the colors you got, and it lists the words
that are still possible, best guesses first. Live at
[mattlove.dev/word](https://mattlove.dev/word/).

Results are written as 5 digits, one per letter: `0` gray, `1` yellow, `2` green.
For example `crane 01020` means R was yellow, N was green, and C, A and E were gray.

## Site (`site/`)

Type a guess, then tap its tiles to cycle each one through gray, yellow and green.
The page shows how many words are still possible and the top 10 suggestions; tap a
suggestion to add it as your next guess.

The guesses are kept in the URL, so a game can be reloaded or linked directly, e.g.
[`/#crane01020`](https://mattlove.dev/word/#crane01020).

It's plain static HTML/JS with no build step. To run it locally:

```bash
python3 -m http.server -d site 8000   # then open http://localhost:8000
```

Pushes to `main` that touch the site are deployed to GitHub Pages by
[`.github/workflows/pages.yml`](.github/workflows/pages.yml).

## How words are ranked

Only words that fit every guess are kept. Repeated letters follow Wordle's rules: in
`speed 21100`, one E is yellow and the other gray, so the answer has exactly one E.

The remaining words are then ranked by letter:

1. Count how many of the remaining words contain each letter.
2. A word containing the most common letter ranks above one without it. If both or
   neither have it, the next most common letter decides, and so on down the list.
3. Words with the same letters (like `least` and `slate`) are ordered alphabetically,
   as are letters with the same count.

## Legacy command-line version (`legacy/`)

The original Go program, which the site mirrors. It takes guesses and results as
pairs of arguments and prints the top 10 words:

```bash
go run ./legacy crane 01020 lions 01120
go build -o word ./legacy && ./word crane 01020
```

It should give the same results as the site. If you change the ranking, change both
`legacy/` and `site/web/js/word.js`.

## Word list

`legacy/wordle.txt` is the original list of 2,315 Wordle answers, from before the NYT
took over the game. Answers added since then aren't in it.

The site's copy, `site/web/js/wordle.js`, is generated from it:

```bash
{ echo "// The Wordle answer list, generated from legacy/wordle.txt."; echo "const wordleList = ["; sed 's/.*/    "&",/' legacy/wordle.txt; echo "]"; } > site/web/js/wordle.js
```

## Layout

| Path | What |
|---|---|
| `site/` | The web version (deployed) |
| `site/web/js/word.js` | Filtering and ranking |
| `site/web/js/page.js` | The page |
| `site/web/js/wordle.js` | Word list, generated from `legacy/wordle.txt` |
| `legacy/` | The original Go command-line program and word list |
