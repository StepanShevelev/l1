package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1000000000)
	b := big.NewInt(200000000)
	c := big.NewInt(0)

	fmt.Println("Результат умножения: ", c.Mul(a, b))
	fmt.Println("Результат деления: ", c.Div(a, b))
	fmt.Println("Результат сложения: ", c.Add(a, b))
	fmt.Println("Результат вычитания: ", c.Sub(a, b))
}
