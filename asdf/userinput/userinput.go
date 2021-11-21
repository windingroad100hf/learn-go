package userinput

import(
"os"
"bufio"
"fmt"
)

func GetCountry() (country string, city string) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter a country")
    country, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    fmt.Println("Enter a city")
    city, err = reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    return country, city

}
