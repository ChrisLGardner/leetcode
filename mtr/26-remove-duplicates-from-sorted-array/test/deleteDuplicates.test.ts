import { removeDuplicates } from '../src/deleteDuplicates';
import { expect } from 'chai';
import 'mocha';
describe('builtin examples',
    () => {
        it('should return 2', () => {
            const result = removeDuplicates([1, 1, 2]);
            expect(result).to.equal(2);
        });
        it('should return 5', () => {
            const result = removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]);
            expect(result).to.equal(5);
        });
    });