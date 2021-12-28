import fs from 'fs/promises';

const verbose = process.env.VERBOSE != null;

async function run() {
    if (process.argv.length !== 3) throw new Error('Must provide an exercise!');

    const [exercise, chapter] = process.argv[2].toLowerCase().match(/ba(?<chapter>\d+)[a-z]/) ?? [];
    if (!exercise || !chapter) throw new Error('Must provide a valid exercise!');

    if (verbose) console.info(`Running exercise ${exercise} from chapter ${chapter}...`);

    let impl: (params: string[]) => any;
    try {
        impl = (await import(`./exercises/ch${chapter}/${exercise}.js`)).default;
    } catch (error) {
        impl = (await import(`./exercises/ch${chapter}/${exercise}.ts`)).default;
    }
    if (!impl) throw new Error(`No implementation for ${exercise} found!`);

    const input = (await fs.readFile('/dev/stdin', 'ascii')).trim().split(/\s+/);
    if (!input || !input.length) throw new Error(`No input passed!`);

    if (verbose) console.info('Running with input:\n', input);

    await fs.writeFile('/dev/stdout', impl(input).toString());
}

run();
