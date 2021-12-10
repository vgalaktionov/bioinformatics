from library.algorithms import reverse_complement


def ba1c(params: list[str]):
    pattern = params[0]
    return reverse_complement(pattern)


def test_ba1c(benchmark):
    expected = "ACCGGGTTTT"
    result = benchmark(ba1c, ["AAAACCCGGT"])
    assert expected == result
