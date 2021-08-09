package pattern

func DrawRectangle(arr [][]int, p Point, width int,length int) int {

	if p.I<0 || p.I>len(arr)-1 ||p.J<0 || p.J>len(arr[0])-1{
		return -1
	}
	rect:=Point{
		I: width,
		J: length,
	}


	lMost, rMost, tMost, bMost := p.J-(rect.J), p.J+(rect.J), p.I+(rect.I), p.I-(rect.I)

	var diagPoint []Point
	bl := Point{I: bMost, J: lMost}
	br := Point{I: bMost, J: rMost}
	tl := Point{I: tMost, J: lMost}
	tr := Point{I: tMost, J: rMost}
	if checkInside(bl, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, bl)
	}
	if checkInside(br, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, br)
	}
	if checkInside(tl, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, tl)
	}
	if checkInside(tr, len(arr)-1, len(arr[0])-1) {
		diagPoint = append(diagPoint, tr)
	}
	if len(diagPoint)==0{
		return -1
	}

	for i:=0;i<len(diagPoint);i++{
		temp:=diagPoint[i]
		if temp.J<p.J &&temp.I<p.I{//bl diag
			makeOne(temp.J,p.J,p.I,temp.I,arr)
		}else if temp.J<p.J &&temp.I>p.I{// tl diag
			makeOne(temp.J,p.J,temp.I,p.I,arr)
		}else if temp.J>p.J &&temp.I>p.I{//tr diag
			makeOne(p.J,temp.J,temp.I,p.I,arr)
		}else if temp.J>p.J &&temp.I<p.I{//br diag
			makeOne(p.J,temp.J,p.I,temp.I,arr)
		}else{
			return -1
		}
	}

	//fmt.Print(len(diagPoint))
	return len(diagPoint)
}