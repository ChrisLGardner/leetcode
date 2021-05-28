using System;
using System.Collections.Generic;
using System.Linq;

namespace FindDegree
{
    public class Solution
    {
        public int FindShortestSubArray(int[] nums)
        {
            Dictionary<int, int> leftMostPos = new Dictionary<int, int>();
            Dictionary<int, int> rightMostPos = new Dictionary<int, int>();
            Dictionary<int, int> totalCount = new Dictionary<int, int>();

            for (int index = 0; index < nums.Length; index++)
            {
                if (totalCount.TryGetValue(nums[index], out int currCount))
                {
                    rightMostPos[nums[index]] = index;
                    totalCount[nums[index]] = currCount + 1;
                }
                else
                {
                    leftMostPos.Add(nums[index], index);
                    rightMostPos.Add(nums[index], index);
                    totalCount.Add(nums[index], 1);
                }
            }

            int degree = totalCount.Max(x => x.Value);
            int answer = int.MaxValue - 1;

            foreach (KeyValuePair<int, int> kvp in totalCount.Where(x => x.Value == degree))
            {
                answer = Math.Min(answer, rightMostPos[kvp.Key] - leftMostPos[kvp.Key]);
            }

            return Math.Min(nums.Length, answer + 1);
        }
    }
}
