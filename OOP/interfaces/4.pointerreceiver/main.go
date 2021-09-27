package main

import "fmt"

// An animal is defined as anything that can speak.
type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c *Cat) Speak() string {
	return "Meoow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "I hate Java."
}

// You can pass suply that function with any value
// what would be the value of "v" ?
func DoSomething(v interface{}) {
	// ...
}

func main() {
	animals := []Animal{&Dog{}, new(Cat), Llama{}, JavaProgrammer{}} // could as well be &Cat{}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
