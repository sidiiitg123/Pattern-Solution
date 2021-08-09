package pattern

func (r Rectangle) draw(input [][]int) int {
	if r.Point.I < 0 || r.Point.I > len(input)-1 || r.Point.J < 0 || r.Point.J > len(input[0])-1 {
		return -1
	}

	lMost, rMost, tMost, bMost := r.Point.J-(r.Length), r.Point.J+(r.Length), r.Point.I+(r.Breadth), r.Point.I-(r.Breadth)
	x:=Check(lMost,rMost,tMost,bMost,input,r)


	return x
}
