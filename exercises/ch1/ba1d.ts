import { findPattern } from '../../lib';

export default function ba1d(params: string[]) {
    const [pattern, genome] = params;
    return findPattern(genome, pattern);
}
