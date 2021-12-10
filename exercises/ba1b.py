from library.algorithms import frequent_words


def ba1b(params: list[str]):
    text = params[0]
    k = int(params[1])
    return frequent_words(text, k)


def test_ba1b(benchmark):
    expected = {"CATG", "GCAT"}
    result = benchmark(ba1b, ["ACGTTGCATGTCGCATGATGCATGAGAGCT", "4"])
    assert expected == result
