import b from 'benny';
import fs from 'fs/promises';

const EXERCISES = [
    'ba1a',
    'ba1b',
    'ba1c',
    'ba1d',
    'ba1e',
    'ba1f',
    'ba1g',
    // 'ba1h', 'ba1i'
];

export default b.suite(
    'Chapter 1',

    ...EXERCISES.map((e) =>
        b.add(e, async () => {
            const [exercise, chapter] = e.match(/ba(?<chapter>\d+)[a-z]/) ?? [];
            const file = await fs.readFile(`./data/rosalind_${e}.txt`, 'ascii');
            const params = file.trim().split(/\s+/);
            let impl: (params: string[]) => any;
            try {
                impl = (await import(`./exercises/ch${chapter}/${exercise}.js`)).default;
            } catch (error) {
                impl = (await import(`./exercises/ch${chapter}/${exercise}.ts`)).default;
            }
            if (!impl) throw new Error(`No implementation for ${exercise} found!`);
            return () => {
                impl(params);
            };
        }),
    ),

    b.cycle(),
    b.complete(),
    b.save({ file: 'results', format: 'csv' }),
);
