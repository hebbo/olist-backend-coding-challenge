package author_test

import (
	"hebbo/stores/author"
	"testing"
)

func TestAdd(t *testing.T) {
	err := author.Add(1, "lorem ipsum")

	if err != nil {
		t.Errorf("unable to add user")
	}

	err = author.Add(1, "")

	if err == nil {
		t.Errorf("author name can't be empty")
	}
}
