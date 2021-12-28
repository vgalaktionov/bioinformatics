import fetch from 'cross-fetch';
import fs from 'fs/promises';
import { JSDOM } from 'jsdom';

async function fetchProblem(exercise: string) {
    const result = await (await fetch(`https://rosalind.info/problems/${exercise}`)).text();
    const dom = new JSDOM(result);
    const sampleDataset = dom.window.document
        .querySelector('#sample-dataset + .codehilite>pre')
        ?.innerHTML.trim()
        .split(/\s/);
    const sampleOutput = dom.window.document
        .querySelector('#sample-output + .codehilite>pre')
        ?.innerHTML.trim()
        .split(/\s/)
        .map((d) => (d.match(/^\d+$/) ? +d : d));
    return [sampleDataset, sampleOutput];
}

async function writeImpl(exercise: string, chapter: string) {
    const dir = `./exercises/ch${chapter}`;
    await fs.mkdir(dir, { recursive: true });

    await fs.writeFile(
        `${dir}/${exercise}.ts`,
        `function foobar(pattern: string): string {
    return pattern
}

export default function ${exercise}(params: string[]) {
    const [pattern] = params;
    return foobar(pattern);
}
        `,
    );
}

async function writeTest(
    exercise: string,
    chapter: string,
    sampleDataset?: (string | number)[],
    sampleOutput?: (string | number)[],
) {
    const dir = `./exercises/ch${chapter}`;
    await fs.mkdir(dir, { recursive: true });

    await fs.writeFile(
        `${dir}/${exercise}.test.ts`,
        `import ${exercise} from './${exercise}';

test('example correct', () => {
    expect(${exercise}(${JSON.stringify(sampleDataset)})).toEqual(${JSON.stringify(sampleOutput)});
});
        `,
    );
}

async function generate() {
    if (process.argv.length !== 3) throw new Error('Must provide an exercise!');

    const [exercise, chapter] = process.argv[2].toLowerCase().match(/ba(?<chapter>\d+)[a-z]/) ?? [];
    if (!exercise || !chapter) throw new Error('Must provide a valid exercise!');

    console.info(`Generating exercise ${exercise} from chapter ${chapter}...`);

    const [sampleDataset, sampleOutput] = await fetchProblem(exercise);

    await Promise.all([writeImpl(exercise, chapter), writeTest(exercise, chapter, sampleDataset, sampleOutput)]);
}

generate();
