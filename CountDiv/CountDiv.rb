def solution(a, b, k)
  n1 = (a - a % k)/k
  n2 = (b - b % k)/k
  (a % k) == 0 ? (n2 - n1 + 1) : n2 - n1
end
