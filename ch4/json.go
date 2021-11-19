package main
import(
"fmt"
//"encoding/json"
"os"
"bufio"
)

type Movie struct{
    Title string
    Year int
    Color bool `json:"color,omitempty"`
}

func main(){
    movielist := make([]Movie, 5)
    for i := 0; i < 5; i++ {
        var Movie = getMovies()
        movielist = append(movielist,Movie)
    }
    fmt.Println(movielist)
}

func getMovies() Movie{
    fmt.Println("Enter A Movie: ")
    reader := bufio.NewReader(os.Stdin)
    title, _ := reader.ReadString(os.Stdin, "\n")
    fmt.Println("Enter A year: ")
    reader := bufio.NewReader(os.Stdin)
    year, _ := reader.ReadString(os.Stdin, "\n")
    fmt.Println("Was it in color: ")
    reader := bufio.NewReader(os.Stdin)
    color, _ := reader.ReadString(os.Stdin, "\n")
    res := movie{title, year, color}
    return res

}
