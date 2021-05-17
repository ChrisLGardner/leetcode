using System;
using System.Collections;
using System.Collections.Generic;
using System.Linq;

namespace singlenumber
{
    public class Solution
    {
        public int SingleNumber(int[] nums)
        {
            Dictionary<int, int> map = new Dictionary<int, int>();

            foreach (int i in nums)
            {
                if (map.TryGetValue(i, out int count))
                {
                    map[i] = count + 1;
                    continue;
                }
                map.Add(i, 1);
            }

            return map.FirstOrDefault(x => x.Value == 1).Key;
        }
    }
}
