package main

import (
	interfaceEmbedding "Embedding/interface"
	"Embedding/struct"
	"fmt"
	_ "net/http/pprof"
	"strconv"
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

	mobile := struct_embedding.Mobile{
		Name:    "Iphone 11",
		Display: display,
		Radio:   radio,
		Button:  struct_embedding.Button{Name: "Call Nikita"},
	}
	CallPerson("Nikita", mobile)

	// Вызов методов анонимных структур
	IPad := struct {
		Year int
		struct_embedding.Display
	}{
		Year:    2019,
		Display: display,
	}
	IPad.Show(strconv.Itoa(IPad.Year))

	// Встраивание интерфейсов
	truck := interfaceEmbedding.Truck{}
	machine := interfaceEmbedding.NewMachine(truck)
	machine.Beep()

}

func CallPerson(person string, mobile struct_embedding.Mobile) {
	mobile.Tap()
	mobile.Show(fmt.Sprintf("Calling %s...", person))
	mobile.Call(person)
}
