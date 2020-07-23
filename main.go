package main

import (
	"Embedding/struct"
	"fmt"
)

func main() {
	display := struct_embedding.Display{
		Height: 1920,
		Width:  1280,
		Color:  true,
	}
	radio := struct_embedding.Radio{
		Frequency: 411.3,
		Amplitude: 10,
	}

	button := struct_embedding.Button{Name: "Call Nikita"}
	mobile := struct_embedding.Mobile{
		Name:    "Iphone 11",
		Display: display,
		Radio:   radio,
		Button:  button,
	}
	CallPerson("Nikita", mobile)
}

func CallPerson(person string, mobile struct_embedding.Mobile) {
	mobile.Tap()
	mobile.Show(fmt.Sprintf("Calling %s...", person))
	mobile.Call(person)
}
