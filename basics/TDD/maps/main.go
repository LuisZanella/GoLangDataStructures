package main

import (
	"errors"
)

var (
	notFoundError   = errors.New("couldn't find the word you were looking for")
	wordExistsError = errors.New("can't add word because it already exists")
)

type Dictionary map[string]string

// ! You should never initialize an empty map variable! i.e: var m map[string]string
// ! instead initialize an empty map
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
func (d Dictionary) search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", notFoundError
	}
	return value, nil
}

func (d Dictionary) add(key string, value string) error {
	_, err := d.search(key)
	if errors.Is(err, notFoundError) {
		d[key] = value
		return nil
	}
	if err == nil {
		return wordExistsError
	}
	return err
}