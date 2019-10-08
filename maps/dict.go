package maps

import "errors"

// Dictionary represents a map
type Dictionary map[string]string

// ErrNotFound reps an error when a value is not found in a map
var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrAlreadyExists = errors.New("word already exists")

// Search returns a key's value from a dict
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add creates a new word/defintion entry in the dict
func (d Dictionary) Add(word, definition string) error {

	_, ok := d[word]

	if ok {
		return ErrAlreadyExists
	}

	d[word] = definition
	return nil
}
