import { frequentWords } from '../../lib';

export default function ba1b(params: string[]) {
    const text = params[0];
    const k = +params[1];
    return frequentWords(text, k);
}
