from library.algorithms import pattern_count


def ba1a(params: list[str]):
    text, pattern = params
    return pattern_count(text, pattern)


def test_ba1a(benchmark):
    expected = 2
    result = benchmark(ba1a, ["GCGCG", "GCG"])
    assert expected == result


def test_ba1a_large(benchmark):
    benchmark(ba1a, open("data/rosalind_ba1a.txt").read().strip().split())
