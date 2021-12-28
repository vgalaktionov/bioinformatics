function minimumSkew(genome: string): number[] {
    const skew: number[] = Array(genome.length + 1).fill(0);
    let minSkew = 0;
    for (let i = 0; i < genome.length; i++) {
        const nucleotide = genome[i];
        const newSkew = nucleotide == 'G' ? skew[i] + 1 : nucleotide == 'C' ? skew[i] - 1 : skew[i];
        skew[i + 1] = newSkew;
        if (newSkew < minSkew) minSkew = newSkew;
    }

    return skew.map((s, i) => (s === minSkew ? i : undefined)).filter(Boolean) as number[];
}

export default function ba1f(params: string[]) {
    const [pattern] = params;
    return minimumSkew(pattern);
}
