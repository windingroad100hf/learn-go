package main

import(
"fmt"
)

func main(){
    var a int = 0
    var b *int = &a
    fmt.Println(b)
    x := []int{1,2,3,4,5,6}
    var y *[]int = &x
    z := &y
    z1 := *z
    fmt.Println(&x)
    fmt.Println(y)
    fmt.Printf("%T\n",x)
    fmt.Printf("%T\n",y)
    fmt.Printf("%T\n",z1)

    d := make([]string, 4)
    fmt.Println(cap(d))
    fmt.Println(len(d))
    fmt.Println(d)
    var r []string
    copy(r, d)
    fmt.Println(cap(d))
    fmt.Println(cap(r))
    //len doesn't count nil values
    //copy doesnt copy nil values

    printme(1,2,3,4,5,6)
    printme()
}

func printme(x ... int ){
    fmt.Println(x)
    fmt.Println()
}



