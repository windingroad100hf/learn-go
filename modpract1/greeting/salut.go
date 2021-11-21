package greeting

import("fmt")

type Sal struct{
    str string
}

type Soomla string

func Salut(){
    a := Sal{str: "Salut"}
    fmt.Println(a.str)
}

func Aurevoir(){
    fmt.Println("Au'Revoir!")
}
