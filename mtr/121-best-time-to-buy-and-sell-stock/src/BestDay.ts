var maxProfit = function maxProfit(prices: number[]): number {
    let lowest = Number.MAX_VALUE;
    let highest = 0;

    for (let i = 0; i < prices.length; i++) {
        if (prices[i] < lowest) {
            lowest = prices[i];
        }
        else {
            if (prices[i] - lowest > highest) {
                highest = prices[i] - lowest;
            }
        }
    }

    return highest;
};
export { maxProfit }