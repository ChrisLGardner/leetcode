const maxArea = function maxArea(height: number[]): number {
    var max = 0;
    var left = 0;
    var right = height.length - 1;

    while (right > left) {
        max = Math.max(max, (right - left) * Math.min(height[right], height[left]));

        if (height[right] > height[left]) left++;
        else right--;
    }

    return max;
};
export { maxArea };