from library.algorithms import pattern_find


def ba1d(params: list[str]):
    pattern, text = params
    return pattern_find(text, pattern)


def test_ba1d(benchmark):
    expected = [1, 3, 9]
    result = benchmark(ba1d, ["ATAT", "GATATATGCATATACTT"])
    assert expected == result
