package struct_embedding

type Wheel struct {
	Form string
}

type Wheels []*Wheel

type RegularCar struct {
	Brand string
	*Wheels
}

func CreateCars(brands []string) []RegularCar {
	var cars []RegularCar
	wheel := Wheel{Form: "circle"}
	wheels := Wheels{&wheel, &wheel, &wheel, &wheel}
	for _, brand := range brands {
		cars = append(cars, RegularCar{
			Brand:  brand,
			Wheels: &wheels,
		})
	}
	return cars
}
