package main
import(
"fmt"
"os"
"bufio"
"strconv"
"strings"
)

func main(){ 
   dec := GetUserDec()
   binary := fmt.Sprintf("%08b",dec )
   fmt.Println("DECIMAL  |  BINARY")
   fmt.Println(dec, "------>", binary)
}

func GetUserDec() int {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("ENTER A DECIMAL NUMBER: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSuffix(input,"\n")
    dec, err := strconv.Atoi(input)
    if err != nil{
        panic(err)
    }
    return dec 
}

