import ba1b from './ba1b';

test('example correct', () => {
    expect(new Set(ba1b(['ACGTTGCATGTCGCATGATGCATGAGAGCT', '4']))).toEqual(new Set(['CATG', 'GCAT']));
});
