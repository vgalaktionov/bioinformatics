import ba1g from './ba1g';

test('example correct', () => {
    expect(ba1g(['GGGCCGTTGGT', 'GGACCGTTGAC'])).toEqual(3);
});
