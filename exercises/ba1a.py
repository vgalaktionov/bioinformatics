from library.algorithms import pattern_count


def ba1a(params: list[str]):
    text, pattern = params
    return pattern_count(text, pattern)


def test_ba1a(benchmark):
    expected = 2
    result = benchmark(ba1a, ["GCGCG", "GCG"])
    assert expected == result
