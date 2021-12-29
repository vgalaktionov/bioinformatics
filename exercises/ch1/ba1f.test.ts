import ba1f from './ba1f';

test('example correct', () => {
    expect(
        ba1f(['CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG']),
    ).toEqual([53, 97]);
});
