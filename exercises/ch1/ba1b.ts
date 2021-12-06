import { countPattern } from '../../lib';

function frequentWords(text: string, k: number) {
    // const frequencies = new DefaultDictionary<string, number>(() => 0);
    const frequencies = new Map<string, number>();
    const lastIndex = text.length - k;
    for (let i = 0; i <= lastIndex; i++) {
        const kmer = text.slice(i, i + k);
        frequencies.set(kmer, countPattern(text, kmer));
    }

    let highestFreq = 0;
    let mostFrequentWords: string[] = [];

    for (const [kmer, freq] of Array.from(frequencies.entries()).sort(([_, a], [__, b]) => a - b)) {
        if (freq === highestFreq) {
            mostFrequentWords.push(kmer);
        } else if (freq > highestFreq) {
            highestFreq = freq;
            mostFrequentWords = [kmer];
        }
    }

    return mostFrequentWords;
}

export default function ba1b(params: string[]) {
    const text = params[0];
    const k = +params[1];
    return frequentWords(text, k);
}
