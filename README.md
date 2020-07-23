# Встраивание в Go

В Go можно объявить анонимные поля у структур:

```go
type Display struct {
    Height int 
    Width int
}


func (d Display) Show() {}

type Mobile struct {
    Brand string
    Display
}

```

В данном случае мы можем использовать методы Display через объекты Mobile:

Неявно - данный метод сначала ищется у вызывающего, а потом у остальных встроенных объектов, удобно переопределять:

`mobile.Show()`

Явно:

`mobile.Display.Show()`

Важно заметить, что данный метод у вызывающего полностью перекрывает метод у встроенных структур и не по сигнатуре, 
а просто по имени, например метод:

```go
func (m Mobile) Show(smth string) string {}
```

перекроет `Show()` у `Display`

Создание таких структур:

```go
mobile := Mobile{Display: Display{Width: 1920, Height: 1280}, Brand: "IPhone"}
```

## Интерфейсы и встроенные структуры

Встраивание структур позволяет родительским структурам соответствовать интерфейсам. Например:


```go
type Shower interface {
    Show(string)
}
```

Теперь Display и Mobile реализуют интерфейс Shower.


## Анонимные структуры и интерфейсы

У анонимных структур нельзя объявить методы, но благодаря встраиванию их можно добавить:

```go
	IPad := struct {
		Year int
		struct_embedding.Display
	}{
		Year: 2019,
		Display: display,
	}
	IPad.Show(strconv.Itoa(IPad.Year))
```


##  Встраивание указателя

Встраивание указателя на структуру позволяет переиспользовать общую часть элементов, например:

```go
type Wheels []Wheel

type RegularCar struct {
    *Wheels // колеса у "обычных" машин обычно одинаковые
} 
```


## Встраивание интерфейсов

Встраивание интерфейсов позволяет явно указать, что структура должна соответствовать интерфейсу, 
но при этом скрыть детали реализации.

```go
type Beeper interface {
	Beep()
}

type Machine struct {
	Beeper
}

type Truck struct{}

func (t Truck) Beep() {
	println("Грузовик делает beep!")
}
``` 

Теперь можно создать структуру таким образом:

```go
truck := interfaceEmbedding.Truck{}
machine := interfaceEmbedding.NewMachine(truck)
machine.Beep()
```

Мы получили доступ к методы `Beep()` у `Truck`, но при этом не имеем доступа к остальным его полям и методам.

Это используется, например, в функции Sort.
