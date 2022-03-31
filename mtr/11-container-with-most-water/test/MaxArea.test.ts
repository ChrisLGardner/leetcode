import { maxArea } from '../src/MaxArea';
import { assert } from 'chai';
import 'mocha';
describe('Example tests',
    () => {
        it('Example 1', () => {
            const input = [1, 8, 6, 2, 5, 4, 8, 3, 7];
            const expected = 49;
            const result = maxArea(input);
            assert.equal(result, expected);
        });
        it('Example 2', () => {
            const input = [1, 1];
            const expected = 1;
            const result = maxArea(input);
            assert.equal(result, expected);
        });
        it('Example 3', () => {
            const input = [4, 3, 2, 1, 4];
            const expected = 16;
            const result = maxArea(input);
            assert.equal(result, expected);
        });
        it('Example 4', () => {
            const input = [1, 2, 1];
            const expected = 2;
            const result = maxArea(input);
            assert.equal(result, expected);
        });
    });
