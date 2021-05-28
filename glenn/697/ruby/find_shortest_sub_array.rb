# @param {Integer[]} nums
# @return {Integer}
def find_shortest_sub_array(nums)
  first_index = {}
  degree = {}
  max_degree = -1
  min_span = -1

  nums.each_with_index do |val, idx|
    if first_index[val].nil?
      first_index[val] = idx
      degree[val] = 0
    end
    degree[val] = degree[val] + 1
    span = idx - first_index[val] + 1

    if degree[val] > max_degree || (degree[val] == max_degree && span < min_span)
      max_degree = degree[val]
      min_span = span
    end
  end

  min_span
end

# Runtime: 88 ms, faster than 66.67% of Ruby online submissions for Degree of an Array.
# Memory Usage: 212.6 MB, less than 44.44% of Ruby online submissions for Degree of an Array.