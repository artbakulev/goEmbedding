package interfaceEmbedding

type Beeper interface {
	Beep()
}

type Machine struct {
	Beeper
}

type Car struct{}

func (c Car) Beep() {
	println("Машина делает beep!")
}

type Truck struct{}

func (t Truck) Beep() {
	println("Грузовик делает beep!")
}


func NewMachine(mechanism Beeper) Machine {
	return Machine{mechanism}
}