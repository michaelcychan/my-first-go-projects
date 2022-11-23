package main

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if ok {
		return definition, nil
	}
	return "", errors.New("no result from your keyword")
}
