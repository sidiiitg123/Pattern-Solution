package pattern

type Point struct {
	I int
	J int
}

func New(m int, n int) Point {
	p := Point{
		I: m,
		J: n,
	}
	return p
}

func checkInside(p Point, row int, col int) bool {
	if (p.I >= 0 && p.I <= row) && (p.J >= 0 && p.J <= col) {
		return true
	}
	return false
}

func makeOne(l int, r int, t int, b int, arr [][]int) [][]int {
	dir:=0
	for l <= r && b <= t {
		if dir == 0 {
			for i := l; i <= r; i++ {
				arr[b][i] = 1
			}
			b++
		} else if dir == 1 {
			for i := b; i <= t; i++ {
				arr[i][r] = 1
			}
			r--
		} else if dir == 2 {
			for i := r; i >= l; i-- {
				arr[t][i] = 1
			}
			t--
		} else if dir == 3 {
			for i := t; i >= b; i-- {
				arr[i][l] = 1
			}
			l++
		}
		dir = (dir + 1) % 4
	}
	return arr
}

func DrawSquare(arr [][]int, p Point, lSquare int) int {

	if p.I<0 || p.I>len(arr)-1 ||p.J<0 || p.J>len(arr[0])-1{
		return -1
	}

	lMost, rMost, tMost, bMost := p.J-(lSquare), p.J+(lSquare), p.I+(lSquare), p.I-(lSquare)

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
