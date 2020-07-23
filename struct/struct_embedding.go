package struct_embedding

import "fmt"

type Display struct {
	Height int
	Width  int
	Color  bool
}

func (d Display) Show(item string) {
	fmt.Printf("There is '%s' on the screen\n", item)
}

type Radio struct {
	Frequency float32
	Amplitude float32
}

func (r Radio) Call(person string) {
	fmt.Printf("Calling %s! beep, beep, beep...\n", person)
}

type Button struct {
	Name string
}

func (b Button) Tap() {
	fmt.Printf("You have tap the '%s' button\n", b.Name)
}

type Mobile struct {
	Name string
	Display
	Radio
	Button
}

type Shower interface {
	Show(string)
}
