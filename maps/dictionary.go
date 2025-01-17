/**
We have covered:
- A full CRUD (Create, Read, Update and Delete) API for our dictionary
- Create errors that are constants (custom error type)
*/

package main

const (
	ErrWordExist = DictionaryErr("the word already exist")
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

// In Go, any type that implements a method with the signature Error() string 
// satisfies the error interface, which is a built-in interface in Go.
func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	// Having a switch like this provides an extra safety net, in case Search 
	// returns an error other than ErrNotFound
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return nil
}
