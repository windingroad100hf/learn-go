package main

import (
	"bufio"
	"fmt"
	"os"
)

type spot struct {
	title string
}

func main() {
	b := spot{title: "hello"}
	b.printString()
	var x = []int{1, 2, 3, 4, 5, 6}
	y := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(equal(x, y))
	getUserArray()

}

func (b spot) printString() {

	fmt.Println("guess my name")
	reader := bufio.NewReader(os.Stdin)
	for {
		val, err := reader.ReadString('\n')
		fmt.Println(val)
		if err != nil {
			fmt.Println("error")
		}
		if val == b.title {
			fmt.Println("that's right")
		} else {
			fmt.Println("try again")
		}

	}

}

func getUserArray() {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 5; i++ {
		fmt.Println("enter an integer")
		val, _ := reader.ReadString('\n')
		fmt.Println(val)
	}
}

func equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for p, q := range a {
		if b[p] != q {
			return false
		}

	}
	return true

}
