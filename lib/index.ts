export function countPattern(text: string, pattern: string) {
    let n = 0;
    const lastIndex = text.length - pattern.length;
    for (let i = 0; i <= lastIndex; i++) {
        n += +(text.slice(i, i + pattern.length) === pattern);
    }
    return n;
}

export type Nucleotide = 'A' | 'C' | 'T' | 'G';
