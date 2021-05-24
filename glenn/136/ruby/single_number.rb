# @param {Integer[]} nums
# @return {Integer}
def single_number(nums)
  instance_hash = {}
  nums.each do |num|
    if instance_hash[num].nil?
      instance_hash[num] = 1
    else
      instance_hash.delete(num)
    end
  end

  return nil unless instance_hash.keys.count == 1
  instance_hash.keys[0]
end


# Runtime: 48 ms, faster than 98.72% of Ruby online submissions for Single Number.
# Memory Usage: 212.8 MB, less than 28.63% of Ruby online submissions for Single Number.