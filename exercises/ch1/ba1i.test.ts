
import ba1i from './ba1i';

test('example correct', () => {
    expect(ba1i(["ACGTTGCATGTCGCATGATGCATGAGAGCT","4","1"])).toEqual(["GATG","ATGC","ATGT"]);
});
        