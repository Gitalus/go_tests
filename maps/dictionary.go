package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// maps lookup pueden retornar dos valores, el segundo un bool si encuentra la llave o no
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
