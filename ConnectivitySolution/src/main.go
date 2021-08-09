package main

import (
	"./pattern"
	"bufio"
	"os"
	"strings"

	//"errors"
	"fmt"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("enter the number of rows  for a matrix")
	scanner.Scan()
	row, _ := strconv.Atoi(scanner.Text())
	fmt.Println("enter the number of columns  for a matrix")
	scanner.Scan()
	col, _ := strconv.Atoi(scanner.Text())


	//pt := pattern.Point{
	//				I: 5,
	//				J: 2,
	//			}
	////x:=pattern.DrawSquare(input,pt,1)
	//x:=pattern.DrawRectangle(input,pt,4,11)
	//if x!=-1{
	//	fmt.Println("Number of square which can be drawn from that point is %d",x)
	//	print(input)
	//}else{
	//	fmt.Println("cant draw a square")
	//}





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

					numberSquare:=pattern.DrawSquare(input, pt, sqLength)
					if numberSquare!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberSquare)
						print(input)
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

					numberCircle:=pattern.DrawCircle(radius,input,pt)
					if numberCircle!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberCircle)
						print(input)
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
					height, _ := strconv.Atoi(scanner.Text())

					fmt.Println("provide the length of rectangle")
					scanner.Scan()
					breadth, _ := strconv.Atoi(scanner.Text())


					numberRectangle:=pattern.DrawRectangle(input, pt, height,breadth)
					if numberRectangle!=-1{
						fmt.Println("Number of square which can be drawn from that point is ",numberRectangle)
						print(input)
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

	func print(input[][]int) {
		for i := 0; i < len(input); i++ {
			for j := 0; j < len(input[0]); j++ {
				fmt.Print(strconv.Itoa(input[i][j]) + "      ")
			}
			fmt.Println("\n")
		}
	}
