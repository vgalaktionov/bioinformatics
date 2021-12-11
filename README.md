# bioinformatics

> A repo with implementations of the exercises in the excellent Bioinformatics Algorithms: An Active Learning Approach (2E).

## Running an exercise

The entrypoint expects the name of the exercise as the first and only CLI argument, and for the data to be passed on STDIN.

E.g:

```bash
# pbpaste is how you write clipboard contents to the terminal on MacOS
pbpaste | python main.py run ba1a

# running the downloaded dataset for rosalind
cat ~/Downloads/rosalind_ba1b.txt | python main.py run ba1b
```

## Generating an exercise scaffold

The entrypoint expects the name of the exercise as the first and only CLI argument.

```bash
python main.py generate ba1b
```

## Tests + benchmarks

```bash
pytest
```

## Development setup

```bash
python -m venv .venv && . .venv/bin/activate && pip install -r requirements.txt
```
