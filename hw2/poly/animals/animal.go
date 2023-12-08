package animals

import "fmt"

type Animal struct {
    Name string
}

func (a *Animal) Speak() {
    fmt.Printf("%s makes sound\n", a.Name)
}
