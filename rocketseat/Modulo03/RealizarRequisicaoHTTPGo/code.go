package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// resp, err := http.Post("https://minhaapi.com", "application/json", nil)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	// data, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))
	// return

	resp, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
