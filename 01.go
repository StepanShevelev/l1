package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    uint
	weight float32
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) GetAge() uint {
	return h.age
}

func (h Human) GetWeight() float32 {
	return h.weight
}

type Action struct {
	Human
}

func main() {
	human := Human{name: "Stepan", age: 23, weight: 100.23}
	action := Action{Human: human}

	fmt.Printf("Name: %s \n Age: %d \n Weight: %.2f\n", action.GetName(), action.GetAge(), action.GetWeight())

}
