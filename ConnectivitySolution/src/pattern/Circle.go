package pattern

import (
	"math"
)

func DrawCircle(radius int,input [][]int,p Point) int{
	val:=-1

	tMost,bMost,lMost,rMost:=p.I+radius,p.I-radius,p.J-radius,p.J+radius
	bl := Point{I: bMost, J: lMost}
	br := Point{I: bMost, J: rMost}
	tl := Point{I: tMost, J: lMost}
	tr := Point{I: tMost, J: rMost}
	//check if all point are inside a matrix
	bool1:=checkInside(bl,len(input),len(input[0]))
	bool2:=checkInside(br,len(input),len(input[0]))
	bool3:=checkInside(tl,len(input),len(input[0]))
	bool4:=checkInside(tr,len(input),len(input[0]))

	if bool1 && bool2 && bool3 && bool4{
		val=1
		for i:=bl.I;i<=tr.I;i++{
			for j:=bl.J;j<=tr.J;j++{
				pt:=Point{
					I:i,
					J:j,

				}
				distance:=math.Abs(float64(pt.J - p.J))+math.Abs(float64(pt.I-p.I))
				if int(distance)<=radius{
					input[pt.I][pt.J]=1
				}else {
					continue
				}
			}
		}
	} else{
		val=-1
	}
	return val
}