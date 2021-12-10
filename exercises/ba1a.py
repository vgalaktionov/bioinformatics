def pattern_count(text: str, pattern: str):
    last_index = len(text) - len(pattern) + 1
    return sum(text[i : i + len(pattern)] == pattern for i in range(last_index))


def ba1a(params: list[str]):
    text, pattern = params
    return pattern_count(text, pattern)


def test_ba1a():
    expected = 2
    result = ba1a(["GCGCG", "GCG"])
    assert expected == result
