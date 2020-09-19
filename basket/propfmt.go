package basket

import "fmt"

func (s Fruit) Fruitfmt() {
	fmt.Printf("%s is in %s color\n", s.Name, s.Color)
}
