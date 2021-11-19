package main

import(
"fmt"
)

func main() {
    var x int = 5
    var y int = 2
    fmt.Println(x/y)
    fmt.Println(float64(x/y))
    fmt.Println(float64(x)/float64(y))
    fmt.Println(x<<4)
    var z float64 = 1.9999999999
    fmt.Println(int(z))
}
