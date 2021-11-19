package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string =  "Hello"
    fmt.Println(len(str))
    fmt.Println(str[0])
    fmt.Print("Now we convert 72 to its UTF string which is 'H': ")
    fmt.Println(string(str[0]))
    x := str[0]
    fmt.Println(x)
    a := str[0:4]
    fmt.Println(a)
    b := `\\\\\///////$$$$$$\n\n\n\n\n\n\n\n\`
    fmt.Println(b)
    c := "SHNOODLE"
    c1 := []byte(c)
    fmt.Println(c1)
    c1 = append(c1, 109)
    fmt.Println(c1)
    c2 := string(c1)
    fmt.Println(c2)


    var g string = "Hello there shnoodle"
    if strings.Contains(g, "lsdjfklsdjfklds") {
        fmt.Println("contains shnoodle")
    } else {
        fmt.Println("doodle")
    }
}
