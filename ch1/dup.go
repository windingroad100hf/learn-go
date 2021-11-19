package main

//list will never be longer than 10
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	names := make(map[int]string)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(input.Text())
		names[len(input.Text())%10] = input.Text()
		fmt.Println(names)
	}

}
