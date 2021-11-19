package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " " ))
	fmt.Println(time.Since(start).Seconds())
}
