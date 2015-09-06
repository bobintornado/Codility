def solution(t)
    return -1 if t.x.nil?

    $max_depth = current_depth = 0 
    recur_traverse(t, current_depth)

    return $max_depth
end

def recur_traverse(t, current_depth)
	if !t.l.nil?
		$max_depth = [current_depth+1,$max_depth].max
		recur_traverse(t.l,current_depth+1)
	end

	if !t.r.nil?
		$max_depth = [current_depth+1,$max_depth].max
		recur_traverse(t.r,current_depth+1)
	end
end