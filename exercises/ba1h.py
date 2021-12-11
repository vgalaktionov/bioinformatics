from library.algorithms import pattern_find_approx


def ba1h(params: list[str]):
    pattern, text, d = params
    d = int(d)
    return pattern_find_approx(text, pattern, d)


def test_ba1h(benchmark):
    expected = [6, 7, 26, 27, 78]
    result = benchmark(ba1h, ["ATTCTGGA", "CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAATGCCTAGCGGCTTGTGGTTTCTCCTACGCTCC", "3"])
    assert expected == result
