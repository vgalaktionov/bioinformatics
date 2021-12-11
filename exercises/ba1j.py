from library.algorithms import frequent_words_with_mismatches_reverse_complements


def ba1j(params: list[str]):
    text, *rest = params
    k, d = map(int, rest)
    return frequent_words_with_mismatches_reverse_complements(text, k, d)


def test_ba1j(benchmark):
    expected = {"ATGT", "ACAT"}
    result = benchmark(ba1j, ["ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"])
    assert expected == result
