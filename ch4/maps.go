package main

import(
"fmt"
)

func main(){
    m := map[string]int{
        "hello":5,
        "billy":20,
        }
    fmt.Println(m)
    val,ok := m["hello"]
    fmt.Println(val)
    fmt.Println(ok)
    _, inmap := m["hello"]
    delete(m,"hello")
    if inmap {
        fmt.Println("DAMN")
    }
    _, inmap = m["hello"]
    if !inmap {
        fmt.Println("NOPE")
    }

}
