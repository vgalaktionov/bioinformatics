import ba1c from './ba1c';

test('example correct', () => {
    expect(ba1c(['AAAACCCGGT'])).toEqual('ACCGGGTTTT');
});
