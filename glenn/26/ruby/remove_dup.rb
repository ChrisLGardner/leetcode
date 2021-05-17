# @param {Integer[]} nums
# @return {Integer}
def remove_duplicates(nums)
  lastnum = nil
  nums.reject! do |num|
    if num == lastnum
      true
    else
      lastnum = num
      false
    end
  end
  return nums.count
end
