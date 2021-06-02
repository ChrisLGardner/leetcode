class Solution(object):
    def maxProfit(self, prices):
        """
        :type prices: List[int]
        :rtype: int
        """
        min_in = prices[0]
        max_out = 0
        for i in range(1,len(prices)):
            if prices[i] < min_in:
                min_in = prices[i]
            elif prices[i] - min_in > max_out:
                max_out = prices[i] - min_in
        
        return max_out