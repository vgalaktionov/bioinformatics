export function countPattern(text: string, pattern: string) {
    let n = 0;
    const lastIndex = text.length - pattern.length;
    for (let i = 0; i <= lastIndex; i++) {
        n += +(text.slice(i, i + pattern.length) === pattern);
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

export type Nucleotide = 'A' | 'C' | 'T' | 'G';
