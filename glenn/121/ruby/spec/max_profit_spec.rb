require './max_profit'

describe 'max_profit' do
  [
    { test: [1], expected: 0 },
    { test: [7,1,5,3,6,4], expected: 5 },
    { test: [7,6,4,3,1], expected: 0 },
    { test: [17,12,15,13,16,14], expected: 4 },
    { test: [17,12,15,13,16,14,11], expected: 4 },
    { test: [17,12,15,13,16,14,11,18], expected: 7 },
  ].each do |testcase|
    it "should get maximum profit for #{testcase[:test]}" do

      arr = testcase[:test]
      result = max_profit(testcase[:test])

      expect(result).to eq(testcase[:expected])
    end
  end
end



# [7,2,5,3,6,4] = 4
# [7,2,5,3,6,4,1] = 4
# [7,2,5,3,6,4,1,8] = 7
