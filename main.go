package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	data, err := os.ReadFile("files/test.txt")
	check(err)

	fmt.Println(string(data))

	fmt.Println(data)

	//Converting byte to bit
	fmt.Printf("%08b", data)

}
