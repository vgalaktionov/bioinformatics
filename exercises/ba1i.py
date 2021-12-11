from library.algorithms import frequent_words_with_mismatches


def foobar(text: str, pattern: str):
    return 1


def ba1i(params: list[str]):
    text, *rest = params
    k, d = map(int, rest)
    return frequent_words_with_mismatches(text, k, d)


def test_ba1i(benchmark):
    expected = {"GATG", "ATGC", "ATGT"}
    result = benchmark(ba1i, ["ACGTTGCATGTCGCATGATGCATGAGAGCT", "4", "1"])
    assert expected == result


def test_ba1i_partial():
    expected = {"GG"}
    result = ba1i(["AGGGT", "2", "1"])
    assert expected == result


def test_ba1i_d0():
    expected = {"GG"}
    result = ba1i(["AGGGT", "2", "0"])
    assert expected == result


def test_ba1i_multiple():
    expected = {"AGG", "GGC", "GCG", "CGG"}
    result = ba1i(["AGGCGG", "3", "0"])
    assert expected == result


# def test_ba1i_large():
#     expected = {"GCACACAGAC", "GCGCACACAC"}
#     result = ba1i(
#         [
#             "CACAGTAGGCGCCGGCACACACAGCCCCGGGCCCCGGGCCGCCCCGGGCCGGCGGCCGCCGGCGCCGGCACACCGGCACAGCCGTACCGGCACAGTAGTACCGGCCGGCCGGCACACCGGCACACCGGGTACACACCGGGGCGCACACACAGGCGGGCGCCGGGCCCCGGGCCGTACCGGGCCGCCGGCGGCCCACAGGCGCCGGCACAGTACCGGCACACACAGTAGCCCACACACAGGCGGGCGGTAGCCGGCGCACACACACACAGTAGGCGCACAGCCGCCCACACACACCGGCCGGCCGGCACAGGCGGGCGGGCGCACACACACCGGCACAGTAGTAGGCGGCCGGCGCACAGCC",
#             "10",
#             "2",
#         ]
#     )
#     assert expected == result
