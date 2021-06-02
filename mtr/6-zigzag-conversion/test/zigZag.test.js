"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const zigZag_1 = require("../src/zigZag");
const chai_1 = require("chai");
require("mocha");
describe('Example tests', () => {
    it('Example 1', () => {
        // Input: s = "PAYPALISHIRING", numRows = 3
        // Output: "PAHNAPLSIIGYIR"
        const input = "PAYPALISHIRING";
        const numRows = 3;
        const expected = "PAHNAPLSIIGYIR";
        const result = zigZag_1.convert(input, numRows);
        chai_1.assert.equal(result, expected);
    });
    it('Example 2', () => {
        // Input: s = "PAYPALISHIRING", numRows = 4
        // Output: "PINALSIGYAHRPI"
        const input = "PAYPALISHIRING";
        const numRows = 4;
        const expected = "PINALSIGYAHRPI";
        const result = zigZag_1.convert(input, numRows);
        chai_1.assert.equal(result, expected);
    });
});
