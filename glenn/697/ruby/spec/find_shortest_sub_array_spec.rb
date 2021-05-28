require './find_shortest_sub_array'

describe 'find_shortest_sub_array' do
  [
    { test: [1], expected: 1 },
    { test: [1,2,2,3,1], expected: 2 },
    { test: [1,2,2,3,1,4,2], expected: 6 }
  ].each do |testcase|
    it "should find shortest sub array for #{testcase[:test]}" do

      arr = testcase[:test]
      result = find_shortest_sub_array(testcase[:test])

      expect(result).to eq(testcase[:expected])
    end
  end
end
