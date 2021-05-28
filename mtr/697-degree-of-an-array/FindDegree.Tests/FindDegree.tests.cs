using System;
using Xunit;

namespace FindDegree.Tests
{
    public class Examples
    {
        [Fact]
        public void Example1()
        {
            int[] input = new int[] { 1, 2, 2, 3, 1 };
            int output = 2;

            Solution sol = new Solution();
            int actual = sol.FindShortestSubArray(input);

            Assert.Equal(output, actual);
        }

        [Fact]
        public void Example2()
        {
            int[] input = new int[] { 1, 2, 2, 3, 1, 4, 2 };
            int output = 6;

            Solution sol = new Solution();
            int actual = sol.FindShortestSubArray(input);

            Assert.Equal(output, actual);
        }
    }

    public class FailedProblems
    {
        [Fact]
        public void Fail1()
        {
            int[] input = new int[] { 1, 2, 2, 1, 2, 1, 1, 1, 1, 2, 2, 2 };
            int output = 9;

            Solution sol = new Solution();
            int actual = sol.FindShortestSubArray(input);

            Assert.Equal(output, actual);
        }
    }
}
