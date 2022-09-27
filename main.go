package main

import (
	"fmt"

	"github.com/uacademy/bigint/bigint"
)

func main() {
	a, err := bigint.NewInt("-50")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("100")
	if err != nil {
		panic(err)
	}

	fmt.Println("Value a:", a)
	fmt.Println("Value b:", b)
	fmt.Println("Add:", bigint.Add(a, b))
	fmt.Println("Sub:", bigint.Sub(a, b))
	fmt.Println("Mult:", bigint.Mult(a, b))
	fmt.Println("Mod:", bigint.Mod(a,b))
	fmt.Println("Abs:", a.Abs())
}
