package basket

import (
	"errors"
	"fmt"
)

func (s Fruit) Fruitfmt() (string, error) {

	_, present := ResFruits[s.Name]

	if present {
		return fmt.Sprintf("%s is in %s color\n", s.Name, s.Color), nil
	} else {
		return "", errors.New("Not in list")
	}

}
