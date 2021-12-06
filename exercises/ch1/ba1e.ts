import { frequentWords } from '../../lib';

function findClumps(genome: string, k: number, L: number, t: number): string[] {
    const words: string[] = [];
    const lastIndex = genome.length - k - L;
    for (let i = 0; i <= lastIndex; i++) {
        words.push(...frequentWords(genome.slice(i, i + L), k, t));
    }
    return Array.from(new Set(words));
}

export default function ba1e(params: string[]) {
    const [genome, ...rest] = params;
    const [k, L, t] = rest.map((p) => +p);
    return findClumps(genome, k, L, t);
}
