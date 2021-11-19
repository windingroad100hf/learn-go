package main 

import (
"fmt"
)

func main(){
    x := "Hello"
//    y := "World"
    xp := &x
 //   yp := &y
    fmt.Println(xp)
    fmt.Println(*xp)

    fmt.Println("/////")
    *xp = "poodle"
   // x = "hola"
    fmt.Println(xp)
    fmt.Println(*xp)

    //fmt.Println(yp)
    //fmt.Println(*yp)
}
