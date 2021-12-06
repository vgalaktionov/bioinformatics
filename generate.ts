import fetch from 'cross-fetch';
import fs from 'fs';
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
        .split(/\s/);
    return [sampleDataset, sampleOutput];
}

function writeImpl(exercise: string, chapter: string) {
    const dir = `./exercises/ch${chapter}`;
    if (!fs.existsSync(dir)) fs.mkdirSync(dir);

    fs.writeFileSync(
        `${dir}/${exercise}.ts`,
        `
        function foobar(pattern: string): string {
            return pattern
        }

        export default function ${exercise}(params: string[]) {
            const [pattern] = params;
            return foobar(pattern);
        }
        `,
    );
}

function writeTest(exercise: string, chapter: string, sampleDataset?: string[], sampleOutput?: string[]) {
    const dir = `./exercises/ch${chapter}`;
    if (!fs.existsSync(dir)) fs.mkdirSync(dir);

    fs.writeFileSync(
        `${dir}/${exercise}.test.ts`,
        `
        import ${exercise} from './${exercise}';

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

    writeImpl(exercise, chapter);
    writeTest(exercise, chapter, sampleDataset, sampleOutput);
}

generate();
