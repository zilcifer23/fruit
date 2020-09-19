package basket

// we restrict it to only certain fruits

var ResFruits = make(map[string]bool)

func init() {
	ResFruits["apple"] = true
	ResFruits["orange"] = true
}

type Fruit struct {
	Name  string
	Color string
}

func Fruiter(n, c string) Fruit {

	f := Fruit{Name: n, Color: c}
	return f

}
