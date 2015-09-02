def solution(a, b, k)
    # write your code in Ruby 2.2
    n1 = (a - a % k)/k
    n2 = (b - b % k)/k
    r = n2 - n1
    if  (a % k) == 0
        (r + 1)
    else 
        r
    end
end