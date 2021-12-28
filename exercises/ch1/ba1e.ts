import { frequentWords } from '../../lib';

function findClumps(genome: string, k: number, L: number, t: number): string[] {
    const lastIndex = genome.length - k - L;
    const words: string[][] = Array(lastIndex).fill([]);
    for (let i = 0; i <= lastIndex; i++) {
        words[i] = frequentWords(genome.slice(i, i + L), k, t);
    }
    return Array.from(new Set(words.flat()));
}

export default function ba1e(params: string[]) {
    const [genome, ...rest] = params;
    const [k, L, t] = rest.map((p) => +p);
    return findClumps(genome, k, L, t);
}
