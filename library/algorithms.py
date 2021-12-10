from itertools import takewhile
from operator import itemgetter


def pattern_count(text: str, pattern: str):
    last_index = len(text) - len(pattern) + 1
    return sum(text[i : i + len(pattern)] == pattern for i in range(last_index))


def pattern_find(text: str, pattern: str):
    last_index = len(text) - len(pattern) + 1
    return [i for i in range(last_index) if text[i : i + len(pattern)] == pattern]


def reverse_complement(pattern: str):
    return "".join({"A": "T", "C": "G", "T": "A", "G": "C"}[n] for n in reversed(pattern))


def frequent_words(text: str, k: int, min_freq: int = 0):
    frequencies = {}
    last_index = len(text) - k + 1
    for i in range(last_index):
        kmer = text[i : i + k]
        if kmer not in frequencies:
            frequencies[kmer] = pattern_count(text, kmer)

    sorted_freqs = sorted(frequencies.items(), key=itemgetter(1), reverse=True)
    highest = sorted_freqs[0][1]

    return {str(kmer) for kmer, _ in takewhile(lambda t: t[1] == highest and t[1] >= min_freq, sorted_freqs)}
