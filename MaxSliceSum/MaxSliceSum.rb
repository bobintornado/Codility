def solution(a)
  max_slice = max_ending = a[0]
  a.slice(1,a.length-1).each do |e|
    max_ending = [e,max_ending+e].max
    max_slice = [max_slice, max_ending].max
  end
  return max_slice
end

solution(  [3, 2, -6, 4, 0] )
