
        function foobar(pattern: string): string {
            return pattern
        }

        export default function ba1d(params: string[]) {
            const [pattern] = params;
            return foobar(pattern);
        }
        