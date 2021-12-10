from itertools import chain

from library.algorithms import frequent_words


def find_clumps(genome: str, k: int, L: int, t: int):
    last_index = len(genome) - k - L + 1
    return set(chain.from_iterable(frequent_words(genome[i : i + L], k, t) for i in range(last_index)))


def ba1e(params: list[str]):
    genome, *rest = params
    k, L, t = map(int, rest)
    return find_clumps(genome, k, L, t)


def test_ba1e(benchmark):
    expected = {"CGACA", "GAAGA", "AATGT"}
    result = benchmark(ba1e, ["CGGACTCGACAGATGTGAAGAAATGTGAAGACTGAGTGAAGAGAAGAGGAAACACGACACGACATTGCGACATAATGTACGAATGTAATGTGCCTATGGC", "5", "75", "4"])
    assert expected == result
