# @param {Integer[]} heights
# @return {Integer}
def max_area(heights)
  left_index = 0
  right_index = heights.length - 1
  biggest_area = 0

  begin
    if heights[right_index] > heights[left_index]
      area = (right_index - left_index) * heights[left_index]
      left_index = left_index + 1
    else
      area = (right_index - left_index) * heights[right_index]
      right_index = right_index - 1
    end
    biggest_area = area if area > biggest_area
  end until left_index >= right_index

  biggest_area
end

# Runtime: 116 ms, faster than 86.52% of Ruby online submissions for Container With Most Water.
# Memory Usage: 216.5 MB, less than 32.02% of Ruby online submissions for Container With Most Water.