import { findPatternApprox } from '../../lib';

export default function ba1h(params: string[]) {
    const [pattern, text, d] = params;
    return findPatternApprox(text, pattern, +d);
}
