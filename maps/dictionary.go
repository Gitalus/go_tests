package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	// maps lookup pueden retornar dos valores, el segundo un bool si encuentra la llave o no
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	// This work because dictionary is already a pointer, so the copy is a copy of a pointer
	d[word] = definition
}
