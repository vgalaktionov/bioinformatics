import fs from 'fs';

async function loader() {
    if (process.argv.length !== 3) throw new Error('Must provide an exercise!');

    const [exercise, chapter] = process.argv[2].toLowerCase().match(/ba(?<chapter>\d+)[a-z]/) ?? [];
    if (!exercise || !chapter) throw new Error('Must provide a valid exercise!');

    console.info(`Running exercise ${exercise} from chapter ${chapter}...`);

    const impl = (await import(`./exercises/ch${chapter}/${exercise}.ts`)).default;
    if (!impl) throw new Error(`No implementation for ${exercise} found!`);

    const input = fs.readFileSync(0, 'utf-8');
    if (!input || !input.length) throw new Error(`No input passed!`);

    console.info('Running with input:\n', input);

    impl(input.trim().split(/\s+/));
}

loader();
