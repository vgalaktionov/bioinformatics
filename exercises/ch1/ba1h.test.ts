
import ba1h from './ba1h';

test('example correct', () => {
    expect(ba1h(["ATTCTGGA","CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAATGCCTAGCGGCTTGTGGTTTCTCCTACGCTCC","3"])).toEqual([6,7,26,27,78]);
});
        