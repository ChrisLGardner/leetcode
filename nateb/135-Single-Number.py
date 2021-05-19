class Solution(object):
    def singleNumber(self, nums):
        """
        Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.
        You must implement a solution with a linear runtime complexity and use only constant extra space.
        """
        d = {}
        for var in nums:
            if var in d.keys():
                d[var] += 1
            else:
                d[var] = 1
        
        for k,v in d.items():
            if v == 1:
                return k