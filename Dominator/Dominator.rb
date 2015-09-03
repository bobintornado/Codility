def solution(a)
  return -1 if a.empty?

  stack = []
  a.each { |e| stack.empty? ? stack.push(e) : stack.last == e ? stack.push(e) : stack.pop }

  return -1 if stack.empty?

  b = [] + a
  a.delete(candidiate)

  return 0 if a.length == 0

  candidiate = stack.pop
  leaderLength = b.length - a.length

  if leaderLength - a.length > 0
    return b.find_index(candidiate)
  else
    return -1
  end
end