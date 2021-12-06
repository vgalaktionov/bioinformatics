import { countPattern } from '../../lib';

function frequentWords(text: string, k: number) {
    const frequencies = new Map<string, number>();
    const lastIndex = text.length - k;
    for (let i = 0; i <= lastIndex; i++) {
        const kmer = text.slice(i, i + k);
        if (!frequencies.has(kmer)) frequencies.set(kmer, countPattern(text.slice(i), kmer));
    }
    const sorted = Array.from(frequencies.entries()).sort(([_, a], [__, b]) => b - a);
    const highestFreq = sorted[0][1];
    const mostFrequentWords: string[] = [];

    for (const [kmer, freq] of sorted) {
        if (freq !== highestFreq) break;
        mostFrequentWords.push(kmer);
    }

    return mostFrequentWords;
}

export default function ba1b(params: string[]) {
    const text = params[0];
    const k = +params[1];
    return frequentWords(text, k);
}
