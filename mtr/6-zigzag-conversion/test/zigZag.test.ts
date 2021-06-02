import { convert } from '../src/zigZag';
import { assert } from 'chai';
import 'mocha';
describe('Example tests',
    () => {
        it('Example 1', () => {
            // Input: s = "PAYPALISHIRING", numRows = 3
            // Output: "PAHNAPLSIIGYIR"

            const input = "PAYPALISHIRING";
            const numRows = 3;
            const expected = "PAHNAPLSIIGYIR";

            const result = convert(input, numRows);

            assert.equal(result, expected);
        });
        it('Example 2', () => {
            // Input: s = "PAYPALISHIRING", numRows = 4
            // Output: "PINALSIGYAHRPI"

            const input = "PAYPALISHIRING";
            const numRows = 4;
            const expected = "PINALSIGYAHRPI";

            const result = convert(input, numRows);

            assert.equal(result, expected);
        });
    });
