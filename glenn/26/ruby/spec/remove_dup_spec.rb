require './remove_dup'

describe 'remove_duplicates' do
  [
    { test: [1,1,2], expected: [1,2] },
    { test: [0,0,1,1,1,2,2,3,3,4], expected: [0,1,2,3,4] }
  ].each do |testcase|
    it "should remove duplicates for #{testcase[:test]}" do

      arr = testcase[:test]
      result = remove_duplicates(testcase[:test])

      expect(result).to eq(testcase[:expected].count)
      expect(arr).to eq(testcase[:expected])
    end
  end
end
