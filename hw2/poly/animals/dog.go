package animals

import "fmt"

type Dog struct {
	*Animal
}

func (d *Dog) Run() {
	fmt.Printf("%s RUNNING \n", d.Name)
}

func (d *Dog) MakeSound() {
	fmt.Printf("%s BARKS \n", d.Name)
}

func NewDog(name string) *Dog {
	return &Dog{&Animal{Name: name}}
}
