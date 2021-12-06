export function* slideWindow(text: string, length: number) {
    const lastIndex = text.length - length;
    for (let i = 0; i <= lastIndex; i++) {
        yield text.slice(i, i + length);
    }
}

export function countPattern(text: string, pattern: string) {
    let n = 0;
    for (const window of slideWindow(text, pattern.length)) {
        n += +(window === pattern);
    }
    return n;
}

export function findPattern(text: string, pattern: string) {
    const positions: number[] = [];
    const lastIndex = text.length - pattern.length;
    for (let i = 0; i <= lastIndex; i++) {
        if (text.slice(i, i + pattern.length) === pattern) positions.push(i);
    }
    return positions;
}

export function frequentWords(text: string, k: number, minFreq = 0) {
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
        if (freq !== highestFreq || freq < minFreq) break;
        mostFrequentWords.push(kmer);
    }

    return mostFrequentWords;
}

export type Nucleotide = 'A' | 'C' | 'T' | 'G';
