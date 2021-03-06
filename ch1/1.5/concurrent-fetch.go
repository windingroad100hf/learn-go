package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
    "strings"
)

func main() {
	for _, url := range os.Args[1:] {
        if  !strings.HasPrefix(url, "http://"){
            url = "http://" + url
            fmt.Println("adding prefix http://")
        }
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
        fmt.Println("SHNOODLE")
        fmt.Println(resp.Status)
        fmt.Println("SHNOODLE")
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
