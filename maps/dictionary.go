package main

import (
	"errors"
)

var ErrNotFound = errors.New("no result from your keyword")

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key] // ok is a bool
	if ok {
		return definition, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
