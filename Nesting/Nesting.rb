def solution(s)
    s = s.split(//)
    
    return 1 if s.empty?
    return 0 if (s.length % 2) != 0 
    
    stack = []
    s.each { |e| e == ")" && stack.last == "(" ? stack.pop : stack.push(e) }
    stack.empty? ? 1 : 0
end