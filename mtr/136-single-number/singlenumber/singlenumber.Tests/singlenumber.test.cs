using System;
using Xunit;

namespace singlenumber.Tests
{
    public class singlenumbertests
    {
        [Fact]
        public void Example1()
        {
            // Example 1:
            // Input: nums = [2, 2, 1]
            // Output: 1

            Solution sol = new();
            int[] input = new int[] {2, 2, 1};
            int expected = 1;

            int output = sol.SingleNumber(input);

            Assert.Equal(expected, output);
        }

        [Fact]
        public void Example2()
        {
            // Example 2:
            // Input: nums = [4, 1, 2, 1, 2]
            // Output: 4

            Solution sol = new();
            int[] input = new int[] { 4, 1, 2, 1, 2 };
            int expected = 4;

            int output = sol.SingleNumber(input);

            Assert.Equal(expected, output);
        }

        [Fact]
        public void Example3()
        {
            // Example 3:
            // Input: nums = [1]
            // Output: 1

            Solution sol = new();
            int[] input = new int[] { 1 };
            int expected = 1;

            int output = sol.SingleNumber(input);

            Assert.Equal(expected, output);
        }
    }
}
