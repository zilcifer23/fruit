package basket

type Fruit struct {
	Name string
	Color String
}

func Fruiter(n, c string) *Fruit {

	f := Fruit{Name: n, Color: c}
	return &f

}
	
	
	
	
