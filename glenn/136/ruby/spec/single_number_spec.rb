require './single_number'

describe 'single_number' do
  [
    { test: [2,2,1], expected: 1 },
    { test: [4,1,2,1,2], expected: 4 },
    { test: [4], expected: 4 },
  ].each do |testcase|
    it "should find the single number for #{testcase[:test]}" do

      arr = testcase[:test]
      result = single_number(testcase[:test])

      expect(result).to eq(testcase[:expected])
    end
  end
end
