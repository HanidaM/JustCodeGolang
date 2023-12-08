package animals

import "fmt"

type Cat struct {
	*Animal
}

func (c *Cat) Run() {
	fmt.Printf("%s RUNNING\n", c.Name)
}

func (c *Cat) MakeSound() {
	fmt.Printf("%s MEAW\n", c.Name)
}

func NewCat(name string) *Cat {
	return &Cat{&Animal{Name: name}}
}
