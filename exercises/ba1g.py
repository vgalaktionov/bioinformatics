from library.algorithms import hamming_distance


def ba1g(params: list[str]):
    textA, textB = params
    return hamming_distance(textA, textB)


def test_ba1g(benchmark):
    expected = 3
    result = benchmark(ba1g, ["GGGCCGTTGGT", "GGACCGTTGAC"])
    assert expected == result
