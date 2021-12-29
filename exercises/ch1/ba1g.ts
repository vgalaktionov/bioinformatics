import { hammingDistance } from '../../lib';

export default function ba1g(params: string[]) {
    const [textA, textB] = params;
    return hammingDistance(textA, textB);
}
