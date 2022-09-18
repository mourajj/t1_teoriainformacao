package main

import (
	"fmt"
	"teoriaInformacao/encoding"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	//data, err := os.ReadFile("files/test.txt")
	//check(err)

	//fmt.Println(data)

	testData := []byte{3, 4, 1, 2, 2, 7, 3, 5, 5, 5, 2, 4, 6, 3}
	//testData2 := []byte{40, 11}

	//fmt.Println(encoding.EliasGamma(testData))
	//fmt.Println(encoding.Unario(testData))
	fmt.Println(encoding.Fibonacci(testData))
}
