# bioinformatics

> A repo with implementations of the exercises in the excellent Bioinformatics Algorithms: An Active Learning Approach (2E).

## Running an exercise

The entrypoint expects the name of the exercise as the first and only CLI argument, and for the data to be passed on STDIN.

E.g:

```bash
# pbpaste is how you write clipboard contents to the terminal on MacOS
pbpaste | npm run ba1a

# running the downloaded dataset for rosalind
npm run ex ba1b <data/rosalind_ba1b.txt
```

## Generating an exercise scaffold

The entrypoint expects the name of the exercise as the first and only CLI argument.

```bash
npm run gen ba1b
```

## Tests + benchmarks

```bash
npm t
```

## Development setup

```bash
npm ci
```
