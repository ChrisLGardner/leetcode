class Solution(object):
    def findShortestSubArray(self, nums):
        """
        the degree of this array is defined as the maximum frequency of any one of its elements.
        Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.
        source: https://leetcode.com/problems/degree-of-an-array/
        """
        d = {}
        maxval = 0

        for var in nums:
            if var in d.keys():
                d[var] += 1
            else:
                d[var] = 1
            # capture max value
            if maxval < d[var] :
                maxval =d[var]
         
        # positions of maxvar
        shorts = []

        for k,v in d.items():
            # skip when value is not highest
            if v != maxval: continue
            pos = []
            # get shortsub for each k
            for x in range(len(nums)):
                if (nums[x] == k):
                    pos.append(x)

            shortsub = pos[len(pos)-1] - pos[0] + 1
            shorts.append(shortsub)
        return min(shorts)