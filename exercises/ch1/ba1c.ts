function reverseComplement(pattern: string): string {
    return pattern
        .split('')
        .map((n: string) => {
            switch (n) {
                case 'A':
                    return 'T';
                case 'C':
                    return 'G';
                case 'T':
                    return 'A';
                case 'G':
                    return 'C';
                default:
                    throw new Error(`Invalid nucleotide ${n}`);
            }
        })
        .reverse()
        .join('');
}

export default function ba1c(params: string[]) {
    const [pattern] = params;
    return reverseComplement(pattern);
}
