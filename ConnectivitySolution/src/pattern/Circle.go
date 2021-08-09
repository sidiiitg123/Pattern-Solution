package pattern

import "math"

func (c Circle) draw(input [][]int) int{
		val:=-1

		tMost,bMost,lMost,rMost:=c.Point.I+c.Radius,c.Point.I-c.Radius,c.Point.J-c.Radius,c.Point.J+c.Radius
		bl := Point{I: bMost, J: lMost}
		br := Point{I: bMost, J: rMost}
		tl := Point{I: tMost, J: lMost}
		tr := Point{I: tMost, J: rMost}
		//check if all point are inside a matrix
		bool1:=c.Pattern.CheckInside(bl,len(input),len(input[0]))
		bool2:=c.Pattern.CheckInside(br,len(input),len(input[0]))
		bool3:=c.Pattern.CheckInside(tl,len(input),len(input[0]))
		bool4:=c.Pattern.CheckInside(tr,len(input),len(input[0]))

		if bool1 && bool2 && bool3 && bool4{
			val=1
			for i:=bl.I;i<=tr.I;i++{
				for j:=bl.J;j<=tr.J;j++{
					pt:=Point{
						I:i,
						J:j,

					}
					distance:=math.Abs(float64(pt.J - c.Point.J))+math.Abs(float64(pt.I-c.Point.I))
					if int(distance)<=c.Radius{
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

