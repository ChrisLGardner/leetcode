"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const BestDay_1 = require("../src/BestDay");
const chai_1 = require("chai");
require("mocha");
describe('Example tests', () => {
    it('Example 1', () => {
        const input = [7, 1, 5, 3, 6, 4];
        const expected = 5;
        const result = BestDay_1.maxProfit(input);
        chai_1.expect(result).to.equal(expected);
    });
    it('Example 2', () => {
        const input = [7, 6, 4, 3, 1];
        const expected = 0;
        const result = BestDay_1.maxProfit(input);
        chai_1.expect(result).to.equal(expected);
    });
});
