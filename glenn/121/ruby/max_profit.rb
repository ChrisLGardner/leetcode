# @param {Integer[]} prices
# @return {Integer}
def max_profit(prices)
  max_profit = 0

  begin_val = prices[0]
  end_val = prices[0]

  for idx in 1..prices.count - 1
    val = prices[idx]
    if val < begin_val
      # We've encountered a new possible profit span
      profit = end_val - begin_val
      max_profit = profit if profit > max_profit
      begin_val = val
      end_val = val
    elsif val > end_val
      end_val = val
    end
  end
  profit = end_val - begin_val
  max_profit = profit if profit > max_profit

  max_profit
end

# Runtime: 136 ms, faster than 92.66% of Ruby online submissions for Best Time to Buy and Sell Stock.
# Memory Usage: 219.5 MB, less than 62.08% of Ruby online submissions for Best Time to Buy and Sell Stock.