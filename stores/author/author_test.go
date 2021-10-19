package author_test

import (
	"hebbo/stores/author"
	"testing"
)

func TestAdd(t *testing.T) {
	empty_store := author.Store{}

	new_store, err := author.Add(empty_store, 1, "lorem ipsum")

	if err != nil {
		t.Errorf("unable to add user")
	}

	if new_store.Data == nil {
		t.Errorf("incorrect store returned")
	}

	_, err = author.Add(empty_store, 1, "")

	if err == nil {
		t.Errorf("author name can't be empty")
	}
}
