package main


import(
 "github.com/windingroad100hf/timeteller/timeteller"
 "github.com/windingroad100hf/timeteller/userinput"
)

func main(){
    city, country := userinput.GetCountry()
    timeteller.GetTime(city, country)
}
