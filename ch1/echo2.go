package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	start := time.Now()
	sep := ""
	for i, ele := range os.Args[1:]{
		sep += ele + " "
		fmt.Println(i , sep)
	}
	fmt.Println(time.Since(start).Seconds())
}
