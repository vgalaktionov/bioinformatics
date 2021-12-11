def minimum_skew(genome: str):
    skew = (len(genome) + 1) * [0]
    min_skew = 0
    for i, nucleotide in enumerate(genome):
        if nucleotide == "G":
            new_skew = skew[i] + 1
        elif nucleotide == "C":
            new_skew = skew[i] - 1
        else:
            new_skew = skew[i]

        skew[i + 1] = new_skew
        if new_skew < min_skew:
            min_skew = new_skew

    return [i for i, s in enumerate(skew) if s == min_skew]


def ba1f(params: list[str]):
    text = params[0]
    return minimum_skew(text)


def test_ba1f(benchmark):
    expected = [53, 97]
    result = benchmark(ba1f, ["CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG"])
    assert expected == result
