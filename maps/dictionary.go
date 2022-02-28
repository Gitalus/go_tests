package maps

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	/*
		Un error es algo que tiene el método Error() que devuelve una string.
		Por esto, se puede usar este nuevo tipo en los métodos que deben retornar un error
		Esta forma de definir los errores es reusable e inmutable, se recomienda.
	*/
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// maps lookup pueden retornar dos valores, el segundo un bool si encuentra la llave o no
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// This work because dictionary is already a pointer, so the copy is a copy of a pointer
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		// In case err is a new one
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		/*
			En este caso se recomienda usar un nuevo error para identificar bien
			cual es el problema. Normalmente, se utilizaría un error existente
		*/
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	// This is the way you delete a word on a map
	delete(d, word)
}
