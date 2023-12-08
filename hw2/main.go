package main

import (
	"NewGoProject/hw2/poly/actions"
	"NewGoProject/hw2/poly/animals"
)

func main() {
	dog := animals.NewDog("Rex")
	cat := animals.NewCat("Mia")

	var sound actions.Sound
	sound = dog
	sound.MakeSound()

	sound = cat
	sound.MakeSound()

	var runnable actions.Run
	runnable = dog
	runnable.Run()

	runnable = cat
	runnable.Run()
}
