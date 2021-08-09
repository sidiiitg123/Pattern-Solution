package pattern



func (s Square) draw(arr [][]int) int {
	if s.Point.I < 0 || s.Point.I > len(arr)-1 || s.Point.J < 0 || s.Point.J > len(arr[0])-1 {
		return -1
	}
    r:=s.Rectangle
	lMost, rMost, tMost, bMost := s.Point.J-(s.SideLength), s.Point.J+(s.SideLength), s.Point.I+(s.SideLength), s.Point.I-(s.SideLength)
	x:=Check(lMost,rMost,tMost,bMost,arr,r)
	return x
}
func Check(lMost int,rMost int,tMost int,bMost int,arr [][]int,r Rectangle) int {
	var diagPoint []Point
	bl := Point{I: bMost, J: lMost}
	br := Point{I: bMost, J: rMost}
	tl := Point{I: tMost, J: lMost}
	tr := Point{I: tMost, J: rMost}
	if r.Pattern.CheckInside(bl, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, bl)
	}
	if r.Pattern.CheckInside(br, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, br)
	}
	if r.Pattern.CheckInside(tl, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, tl)
	}
	if r.Pattern.CheckInside(tr, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, tr)
	}
	if len(diagPoint) == 0 {
		return -1
	}

	for i := 0; i < len(diagPoint); i++ {
		temp := diagPoint[i]
		if temp.J < r.Point.J && temp.I < r.Point.I { //bl diag
			r.Pattern.MakeOne(temp.J, r.Point.J, r.Point.I, temp.I, arr)
		} else if temp.J < r.Point.J && temp.I > r.Point.I { // tl diag
			r.Pattern.MakeOne(temp.J, r.Point.J, temp.I, r.Point.I, arr)
		} else if temp.J > r.Point.J && temp.I > r.Point.I { //tr diag
			r.Pattern.MakeOne(r.Point.J, temp.J, temp.I, r.Point.I, arr)
		} else if temp.J > r.Point.J && temp.I < r.Point.I { //br diag
			r.Pattern.MakeOne(r.Point.J, temp.J, r.Point.I, temp.I, arr)
		} else {
			return -1
		}
	}
	return len(diagPoint)
}
