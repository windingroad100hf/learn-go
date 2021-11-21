package timeteller

import(
"time"
"fmt"
"strings"
)




func GetTime(country string, city string){
    countrycity := strings.TrimSuffix(country, "\n") + "//" + strings.TrimSuffix(city, "\n")
    location, err := time.LoadLocation(countrycity)
    if err != nil{
        panic(err)
    }
    now := time.Now()

    strtime := now.Format(time.ANSIC)
    latime, err:= time.ParseInLocation(time.ANSIC, strtime, location)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(latime)
}
