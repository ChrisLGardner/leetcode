"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const deleteDuplicates_1 = require("../src/deleteDuplicates");
const chai_1 = require("chai");
require("mocha");
describe('builtin examples', () => {
    it('should return 2', () => {
        const result = deleteDuplicates_1.removeDuplicates([1, 1, 2]);
        chai_1.expect(result).to.equal(2);
    });
    it('should return 5', () => {
        const result = deleteDuplicates_1.removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]);
        chai_1.expect(result).to.equal(5);
    });
});
