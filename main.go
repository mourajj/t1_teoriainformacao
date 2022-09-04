package main

import (
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

	GolombCoding(2, data)

}
