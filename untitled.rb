def solution(a)
    # write your code in Ruby 2.2
    hash = {}
    a.each_index do |index|
    	hash[index] = a[index]
    end
    three = []
    hash.sort_by { |k,v| v.abs }.last(3).map { |k,v| three << v }
    puts three
end

solution([0,3,2,5])
