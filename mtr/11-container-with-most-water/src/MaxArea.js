"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.maxArea = void 0;
const maxArea = function maxArea(height) {
    var max = 0;
    var left = 0;
    var right = height.length - 1;
    while (right > left) {
        max = Math.max(max, (right - left) * Math.min(height[right], height[left]));
        if (height[right] > height[left])
            left++;
        else
            right--;
    }
    return max;
};
exports.maxArea = maxArea;
