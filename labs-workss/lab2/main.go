package main

import (
	"errors"
	"fmt"
	"strings"
)

func formatIP(ip [4]byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func listEven(a, b int32) ([]int32, error) {
	if b < a {
		return nil, errors.New(fmt.Sprintf("Неверно задан диапазон [%d, %d]", a, b))
	}

	slice := []int32{}

	for i := a; i <= b; i++ {
		if i%2 == 0 {
			slice = append(slice, i)
		}
	}

	return slice, nil
}

func testMap(str string) map[string]int {
	myMap := make(map[string]int)
	chars := strings.Split(str, "")
	for i := 0; i < len(chars); i++ {
		if myMap[chars[i]] != 0 {
			myMap[chars[i]] = myMap[chars[i]] + 1
		} else {
			myMap[chars[i]] = 1
		}
	}

	return myMap
}

func main() {
	// 1
	ip := [4]byte{127, 0, 0, 1}
	formattedIP := formatIP(ip)
	fmt.Println(formattedIP)

	// 2
	var slice, error1 = listEven(0, 4)

	if error1 != nil {
		fmt.Println(error1.Error())
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println("Четное число ", slice[i])
	}

	var _, error2 = listEven(4, 0)
	if error2 != nil {
		fmt.Println(error2.Error())
	}

	// 3
	fmt.Println(testMap("SIIIUUUUUUUU"))

	// Опрeделите структуру «точка»
	point := Point{X: 1, Y: 12}
	fmt.Println("Point", point)
	triangle := Triangle{
		A: Point{X: 1, Y: 1},
		B: Point{X: 2, Y: 1},
		C: Point{X: 1, Y: 5}}

	circle := Circle{Center: Point{X: 1, Y: 1}, Radius: 4}

	fmt.Print("Triangle: ")
	printArea(triangle)
	fmt.Print("Circle: ")
	printArea(circle)

	increment := func(a float64) float64 {
		return a + 1
	}

	slice2 := []float64{1, 3, 4, 5, 8}
	fmt.Println("срез до Map()", slice2)
	slice2 = Map(slice2, increment)
	fmt.Println("срез после Map()", slice2)
}
