export function countPattern(text: string, pattern: string) {
    let n = 0;
    const lastIndex = text.length - pattern.length;
    for (let i = 0; i <= lastIndex; i++) {
        if (text.slice(i, i + pattern.length) === pattern) n++;
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

export function findPatternApprox(text: string, pattern: string, d: number) {
    const positions: number[] = [];
    const lastIndex = text.length - pattern.length;
    for (let i = 0; i <= lastIndex; i++) {
        if (hammingDistance(text.slice(i, i + pattern.length), pattern) <= d) positions.push(i);
    }
    return positions;
}

export function frequentWords(text: string, k: number, minFreq = 0) {
    const frequencies = new Map<string, number>();
    const lastIndex = text.length - k;
    let highestFreq = 0;
    for (let i = 0; i <= lastIndex; i++) {
        const kmer = text.slice(i, i + k);
        if (!frequencies.has(kmer)) {
            const freq = countPattern(text.slice(i), kmer);
            frequencies.set(kmer, freq);
            highestFreq = Math.max(freq, highestFreq);
        }
    }
    const mostFrequentWords: string[] = [];

    for (const [kmer, freq] of frequencies.entries()) {
        if (freq !== highestFreq || freq < minFreq) continue;
        mostFrequentWords.push(kmer);
    }

    return mostFrequentWords;
}

export function hammingDistance(textA: string, textB: string): number {
    if (textA.length !== textB.length) throw new Error('patterns must be of equal length');
    let distance = 0;
    for (let i = 0; i < textA.length; i++) {
        if (textA[i] !== textB[i]) distance++;
    }
    return distance;
}

function product(iterables: string[], repeat = 1) {
    const repeated = [];
    for (let i = 0; i < repeat; i++) {
        repeated.push(iterables.slice()); // Clone
    }
    return repeated
        .reduce(
            (acc, cur) => {
                const tmp: string[][] = [];
                acc.forEach((a0) => {
                    cur.forEach((a1) => {
                        tmp.push(a0.concat(a1));
                    });
                });
                return tmp;
            },
            [[]] as string[][],
        )
        .map((vs) => vs.join(''));
}
