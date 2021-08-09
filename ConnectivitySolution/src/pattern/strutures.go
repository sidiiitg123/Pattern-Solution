package pattern

import (
	"fmt"
	"strconv"
)

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

type Pattern struct {
	Name string
}

type Square struct {
	Rectangle
	Pattern
    Point
	SideLength int
}
type Circle struct {
	Pattern
	Point
	Radius int
}
type Rectangle struct {
	Pattern
	Point
	Length  int
	Breadth int
}

type Shape interface {
	draw([][]int) int
}
func Info(z Shape,input [][]int) int{
	return z.draw(input)
}

func(pat Pattern) Print(input [][]int){
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			fmt.Print(strconv.Itoa(input[i][j]) + "      ")
		}
		fmt.Println("\n")
	}
}
func (pat Pattern) CheckInside(p Point, row int, col int) bool {
	if (p.I >= 0 && p.I <= row) && (p.J >= 0 && p.J <= col) {
		return true
	}
	return false
}

func (pat Pattern) MakeOne(l int, r int, t int, b int, arr [][]int) [][]int {
	dir := 0
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

