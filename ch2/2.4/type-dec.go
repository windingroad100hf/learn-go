package main

import(
    "fmt"
)

func main(){
    type gmap map[int]string
    type shnoodle string
    var x shnoodle = "12345"
    fmt.Println(x)
    x = "hello"
    fmt.Println(x)
    var mapx gmap = make(map[int]string)
    mapx[1] = "one"
    a, b := mapx[1]
    fmt.Println(a, b, mapx)
    y := 4 
    v, ok := int(y) 
    fmt.Println(y,ok)
    }

