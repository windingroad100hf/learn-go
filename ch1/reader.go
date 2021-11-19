package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nameMap := make(map[string]int)
	for {

		fmt.Println("enter your name or type 'exit' to exit")
		reader := bufio.NewReader(os.Stdin)
		name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("an error occured", err)
		}
		name = strings.TrimSuffix(name, "\n")
		if name == "exit" {
			break
		}
		fmt.Println("It's nice to meet you, " + name)
		nameMap[name]++
		fmt.Println("We have stored your name in a map for safe keeping")
		fmt.Println("The Map looks like this:")
		fmt.Println(nameMap)
		fmt.Printf("There are %d %v %v ", nameMap[name], "other people named ", name)
		fmt.Println("----------------------------------------")
	}

}
