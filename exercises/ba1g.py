def hamming_distance(textA: str, textB: str):
    return sum(textA[i] != textB[i] for i in range(len(textA)))


def ba1g(params: list[str]):
    textA, textB = params
    return hamming_distance(textA, textB)


def test_ba1g(benchmark):
    expected = 3
    result = benchmark(ba1g, ["GGGCCGTTGGT", "GGACCGTTGAC"])
    assert expected == result
