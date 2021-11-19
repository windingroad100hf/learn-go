package main

import(
"fmt"
"encoding/json"
"os"
)

type dog struct{
    Name string `json:"shnoodle_name"`
    Age int 
    Breed string
    Adopted bool
}

func main(){
    molly := dog{"Molly", 2, "poodle", true}
    fmt.Println(molly)
    b, _:= json.MarshalIndent(molly,"", "   ")
    fmt.Println(string(b))
    os.Stdout.Write(b)
// why doesn't this print Name
    var unmarshalledMolly struct{Name string; Age int; Breed string; Adopted bool}
    if err := json.Unmarshal(b,&unmarshalledMolly); err != nil {
        panic(err)
    }
    fmt.Println(unmarshalledMolly)
}


//func handleErr(err){
//    if err != nil {
//        panic(err)
//    }
//}
