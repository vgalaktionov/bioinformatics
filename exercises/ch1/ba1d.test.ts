import ba1d from './ba1d';

test('example correct', () => {
    expect(ba1d(['ATAT', 'GATATATGCATATACTT'])).toEqual(['1', '3', '9']);
});
