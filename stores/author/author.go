package author

import "errors"

type Store struct {
	Data map[int]Author
}

type Author struct {
	id   int64
	name string
}

func Add(s Store, id int64, name string) (Store, error) {
	if len(name) == 0 {
		return s, errors.New("incorrect name")
	}

	return Store{
		Data: map[int]Author{
			1: {id: id, name: name},
		},
	}, nil
}
