package main

import(
"fmt"
)

type dog struct {
    Name string
    Age int
    Breed string
}


func main() {
    molly := dog{"Molly",2,"poodle"}
    fmt.Println(molly)
//    var breed string = molly.Speak()
    breed := molly.Speak()
    fmt.Println(breed)
}



func (shnood dog) Speak() (z string) {
   switch shnood.Breed{
       case "poodle":
       fmt.Println("WOOF I'm a poodle!!!")
       case "bulldog":
       fmt.Println("WOOF I'm a bulldog!!!")
       default:
       fmt.Println("meow I'm a cat!")
   }
   return shnood.Breed

}
