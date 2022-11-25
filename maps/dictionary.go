package main

const (
	ErrNotFound      = DictionaryErr("no result from your keyword")
	ErrWordExists    = DictionaryErr("key already existed, cannot add word")
	errWordNotExists = DictionaryErr("the key does not exisit, cannot update")
)

type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key] // ok is a bool
	if ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return ErrWordExists
	case ErrNotFound:
		d[key] = value
		return nil
	default:
		return err
	}
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
		return nil
	case ErrNotFound:
		return errWordNotExists
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
