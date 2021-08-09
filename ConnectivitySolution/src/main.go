package main

import (
	"./pattern"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {




	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("enter the number of rows  for a matrix")
	scanner.Scan()
	row, _ := strconv.Atoi(scanner.Text())
	fmt.Println("enter the number of columns  for a matrix")
	scanner.Scan()
	col, _ := strconv.Atoi(scanner.Text())
	//	input := make([][]int, row)
	//	for i := 0; i < len(input); i++ {
	//		input[i] = make([]int,col)
	//	}



	//				pt := pattern.Point{
	//					I: 0,
	//					J: 4,
	//				}
	//
	//				parentRectangle:=pattern.Pattern{
	//					Name: "rectangle",
	//				}
	//				rectObj:=pattern.Rectangle{
	//					Pattern: parentRectangle,
	//					Point:   pt,
	//					Length:  2,
	//					Breadth: 1,
	//				}
	//				pattern.Info(rectObj,input)
	//rectObj.Pattern.Print(input)


	for {
		input := make([][]int, row)
		for i := 0; i < len(input); i++ {
			input[i] = make([]int,col)
		}
		fmt.Print("Enter 0 to quit\nEnter 1 to print a square \nEnter 2 to print a circle \nEnter 3 to print a rectangle\n")
		scanner.Scan()
		number, _ := strconv.Atoi(scanner.Text())

			if number == 1 {
				for{
					fmt.Print("give the x and y coordinate of point from where you want to draw a square\n")

					scanner.Scan()
					x, _ := strconv.Atoi(scanner.Text())
					scanner.Scan()
					y, _ := strconv.Atoi(scanner.Text())

					pt := pattern.Point{
						I: x,
						J: y,
					}
					fmt.Println("provide the length of the square")
					scanner.Scan()
					sqLength, _ := strconv.Atoi(scanner.Text())

					parentSquare:=pattern.Pattern{
						Name: "square",
					}
					r:=pattern.Rectangle{
						Pattern: parentSquare,
						Point:   pt,
						Length:  sqLength,
						Breadth: sqLength,
					}
					sqObj:=pattern.Square{
						Rectangle: r,
						Pattern:    parentSquare,
						Point:      pt,
						SideLength: sqLength,
					}
					numberSquare:=pattern.Info(sqObj,input)


					if numberSquare!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberSquare)
						sqObj.Pattern.Print(input)
					}else{
						fmt.Println("cant draw a square")
					}


					fmt.Println("press q if u want to quit or any other key to continue drawing ")
					scanner.Scan()
					text := scanner.Text()

					if strings.ToLower(text) =="q"{
						break
					}
				}
			} else if number == 2 {
				for{
					fmt.Print("give the x and y coordinate of circle center\n")

					scanner.Scan()
					x, _ := strconv.Atoi(scanner.Text())
					scanner.Scan()
					y, _ := strconv.Atoi(scanner.Text())

					pt := pattern.Point{
						I: x,
						J: y,
					}
					fmt.Println("provide the radius of circle")
					scanner.Scan()
					radius, _ := strconv.Atoi(scanner.Text())

					parentSquare:=pattern.Pattern{
						Name: "circle",
					}

					circleObj:=pattern.Circle{
						Pattern: parentSquare,
						Point:   pt,
						Radius: radius,
					}
					numberCircle:=pattern.Info(circleObj,input)
					if numberCircle!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberCircle)
						circleObj.Pattern.Print(input)
					}else{
						fmt.Println("cant draw a circle")
					}
					fmt.Println("press q if u want to quit or any other key to continue drawing")
					scanner.Scan()
					text := scanner.Text()

					if strings.ToLower(text) =="q"{
						break
					}
				}

			} else if number == 3 {
				for{
					fmt.Print("give the x and y coordinate of point from where you want to draw a rectangle\n")

					scanner.Scan()
					x, _ := strconv.Atoi(scanner.Text())
					scanner.Scan()
					y, _ := strconv.Atoi(scanner.Text())

					pt := pattern.Point{
						I: x,
						J: y,
					}
					fmt.Println("provide the width of rectangle")
					scanner.Scan()
					width, _ := strconv.Atoi(scanner.Text())

					fmt.Println("provide the length of rectangle")
					scanner.Scan()
					length, _ := strconv.Atoi(scanner.Text())

					parentRectangle:=pattern.Pattern{
						Name: "rectangle",
					}
					rectObj:=pattern.Rectangle{
						Pattern: parentRectangle,
						Point:   pt,
						Length:  length,
						Breadth: width,
					}
					numberRectangle:=pattern.Info(rectObj,input)

					if numberRectangle!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberRectangle)
						rectObj.Pattern.Print(input)
					}else{
						fmt.Println("cant draw a rectangle")
					}
					fmt.Println("press q if u want to quit or any other key to continue drawing")
					scanner.Scan()
					text := scanner.Text()

					if strings.ToLower(text) =="q"{
						break
					}
				}
			}else if number==0{
				return
			}else {
				fmt.Println("try again")
			}


		}

	}


