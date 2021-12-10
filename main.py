import argparse
import importlib
import re
import sys
from pathlib import Path

import httpx
from bs4 import BeautifulSoup
from loguru import logger

EXERCISE_PATTERN = re.compile(r"ba(\d+)[a-z]")


def exercise(arg: str):
    if EXERCISE_PATTERN.search(arg):
        return arg
    raise argparse.ArgumentError(f"invalid exercise: {arg}")


def run(exercise: str):
    logger.info(f"executing exercise {exercise} with input from stdin...")

    module = importlib.import_module(f"exercises.{exercise}")
    impl = getattr(module, exercise)
    params = sys.stdin.read().split()
    result = impl(params)

    print(result)


def generate(exercise: str):
    logger.info(f"generating exercise {exercise} scaffold...")

    response = httpx.get(f"https://rosalind.info/problems/{exercise}/")
    soup = BeautifulSoup(response.text, features="html.parser")
    sample_dataset = soup.select_one("#sample-dataset + .codehilite>pre").text.strip().split()
    sample_output = soup.select_one("#sample-output + .codehilite>pre").text.strip()
    try:
        sample_output = int(sample_output)
    except ValueError:
        pass

    with open(Path(__file__).parent / "exercises" / f"{exercise}.py", "w") as f:
        f.write(
            f"""
def foobar(text: str, pattern: str):
    return 1


def {exercise}(params: list[str]):
    text, pattern = params
    return foobar(text, pattern)


def test_{exercise}():
    expected = {sample_output if type(sample_output) == int else f"'{sample_output}'"}
    result = {exercise}({sample_dataset})
    assert expected == result
"""
        )


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Bioinformatics Algorithms CLI")
    subparsers = parser.add_subparsers(title="commands")
    run_parser = subparsers.add_parser("run")
    run_parser.add_argument("exercise", help="Rosalind exercise code", type=exercise)
    run_parser.set_defaults(func=run)
    generate_parser = subparsers.add_parser("generate")
    generate_parser.add_argument("exercise", help="Rosalind exercise code", type=exercise)
    generate_parser.set_defaults(func=generate)
    args = parser.parse_args()

    args.func(args.exercise)
