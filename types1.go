package main

import(
"fmt"
)

func main(){
    compareInts()
    
}

func compareInts(){
    var a uint = 33
    fmt.Println("a: ", a)
    fmt.Println("a string: ", string(a))
    fmt.Println("a byte: ", byte(a))
    fmt.Println("a int: ", int(a))

    fmt.Println()

    var b byte = 33
    fmt.Println("b: ", b)
    fmt.Println("b string: ", string(b))

    fmt.Println()

    var c byte = 90
    fmt.Println("c int: ", int(c))

    fmt.Println()


    bytelist := []byte{33,34,35}
    fmt.Println("bytelist: ", bytelist)
    fmt.Println("bytelist as string: ", string(bytelist))
    fmt.Println("bytelist as string: ", string((bytelist)[0]))

    poodle := "poodle"

    shnoodle := &poodle 
    
    Pointer(shnoodle)

}


func Pointer(str *string){
    fmt.Println(str)
    fmt.Println(&str)
    fmt.Println(str)
}
