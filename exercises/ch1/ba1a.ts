import { countPattern } from '../../lib';

export default function ba1a(input: string[]) {
    const [text, pattern] = input;
    return countPattern(text, pattern);
}
