from functools import partial
from itertools import product, takewhile
from operator import itemgetter


def pattern_count(text: str, pattern: str):
    last_index = len(text) - len(pattern) + 1
    return sum(text[i : i + len(pattern)] == pattern for i in range(last_index))


def pattern_find(text: str, pattern: str):
    last_index = len(text) - len(pattern) + 1
    return [i for i in range(last_index) if text[i : i + len(pattern)] == pattern]


COMPLEMENT = {"A": "T", "C": "G", "T": "A", "G": "C"}


def reverse_complement(pattern: str):
    return "".join(COMPLEMENT[n] for n in reversed(pattern))


def hamming_distance(textA: str, textB: str):
    return sum(textA[i] != textB[i] for i in range(len(textA)))


def pattern_find_approx(text: str, pattern: str, d: int):
    last_index = len(text) - len(pattern) + 1
    return [i for i in range(last_index) if hamming_distance(text[i : i + len(pattern)], pattern) <= d]


def pattern_count_approx(text: str, pattern: str, d: int):
    last_index = len(text) - len(pattern) + 1
    return sum(hamming_distance(text[i : i + len(pattern)], pattern) <= d for i in range(last_index))


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


def frequent_words_with_mismatches(text: str, k: int, d: int):
    all_kmers = ("".join(ns) for ns in product("ACTG", repeat=k))
    frequencies = {kmer: pattern_count_approx(text, kmer, d) for kmer in all_kmers}

    sorted_freqs = sorted(frequencies.items(), key=itemgetter(1), reverse=True)
    highest = sorted_freqs[0][1]

    return {str(kmer) for kmer, _ in takewhile(lambda t: t[1] == highest, sorted_freqs)}


def frequent_words_with_mismatches_reverse_complements(text: str, k: int, d: int):
    all_kmers = ("".join(ns) for ns in product("ACTG", repeat=k))
    frequencies = {kmer: pattern_count_approx(text, kmer, d) + pattern_count_approx(text, reverse_complement(kmer), d) for kmer in all_kmers}

    sorted_freqs = sorted(frequencies.items(), key=itemgetter(1), reverse=True)
    highest = sorted_freqs[0][1]

    return {str(kmer) for kmer, _ in takewhile(lambda t: t[1] == highest, sorted_freqs)}
