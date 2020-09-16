package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cannot update non-exist word")
	errWordExist  = errors.New("Word already exist")
)

// Search for word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add to Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	/*
		if err == errNotFound {
			d[word] = def
		} else if err == nil {
			return errWordExist
		}
	*/
	return nil
}

// Update to Dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete from Dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
