
def foobar(text: str, pattern: str):
    return 1


def ba1b(params: list[str]):
    text, pattern = params
    return foobar(text, pattern)


def test_ba1b():
    expected = 'CATG GCAT'
    result = ba1b(['ACGTTGCATGTCGCATGATGCATGAGAGCT', '4'])
    assert expected == result
