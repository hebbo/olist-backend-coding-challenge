package author

import "errors"

type author struct {
	id   int64
	name string
}

func Add(id int64, name string) error {
	if len(name) == 0 {
		return errors.New("incorrect name")
	}

	return nil
}
