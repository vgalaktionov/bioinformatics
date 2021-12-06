import ba1e from './ba1e';

test('example correct', () => {
    expect(
        new Set(
            ba1e([
                'CGGACTCGACAGATGTGAAGAAATGTGAAGACTGAGTGAAGAGAAGAGGAAACACGACACGACATTGCGACATAATGTACGAATGTAATGTGCCTATGGC',
                '5',
                '75',
                '4',
            ]),
        ),
    ).toEqual(new Set(['CGACA', 'GAAGA', 'AATGT']));
});
