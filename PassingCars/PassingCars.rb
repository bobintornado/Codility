def solution(a)
  # write your code in Ruby 2.2
  sum = 0
  b = Array.new(a)
  b.delete(0)
  sum = b.length
  pairs = 0
  last_index = 0
  a.each_index do |index|
    if a[index] == 0
      if index == a.length-1
        break
      end
      if index == 0
        pairs += sum
      else
        sum = sum - (index - last_index - 1)
        pairs += sum
        last_index = index
      end
    end
    if pairs > 1000000000
      return -1
    end
  end
  pairs
end
