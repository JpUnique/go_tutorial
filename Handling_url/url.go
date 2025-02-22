package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://example.com:2000/leanrning/go/hello"

func main() {
	fmt.Println(myurl)

	//PARSING URL
	parsedURL, err := url.Parse(myurl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Protocol:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)
	fmt.Println("Port:", parsedURL.Port())

}
