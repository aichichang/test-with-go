package main

import "errors"

// Somthing about maps
// Attempts to write to a nil map will cause a runtime panic.
// Instead, you can initialize an empty map
// or use the make keyword to create a map for you:
// var dictionary = map[string]string{}
// var dictionary = make(map[string]string)
// Both approaches create an empty hash map and point dictionary at it

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(str string) (string, error) {
	// map lookup -  It can return 2 values
	// The second value is a boolean which indicates if the key was found successfully
	definition, ok := d[str]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// An interesting property of maps is that
// you can modify them without passing as an address to it (e.g &myMap)
func (d Dictionary) Add(str, definition string) error {
	_, err := d.Search(str)

	switch err {
	// if the string cannot be found, add it to the dictionary
	case ErrNotFound:
		d[str] = definition
	// if nil returned, means the string alreay exists in the dictionary
	case nil:
		return ErrWordExists
	// return any other errors from the Search function
	default:
		return err
	}

	return nil
}
