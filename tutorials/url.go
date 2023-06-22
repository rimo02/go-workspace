package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com"

func main() {

	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println(response)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println(content)

}
