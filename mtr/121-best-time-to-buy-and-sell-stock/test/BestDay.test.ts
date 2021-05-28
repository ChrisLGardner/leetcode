import { maxProfit } from '../src/BestDay';
import { expect } from 'chai';
import 'mocha';
describe('Example tests',
    () => {
        it('Example 1', () => {
            const input = [7, 1, 5, 3, 6, 4];
            const expected = 5;

            const result = maxProfit(input);
            expect(result).to.equal(expected);
        });

        it('Example 2', () => {
            const input = [7, 6, 4, 3, 1];
            const expected = 0;

            const result = maxProfit(input);
            expect(result).to.equal(expected);
        });
    });

