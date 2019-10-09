package maps

const (
	// ErrNotFound reps an error when a value is not found in a map
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists is thrown when a word exists in the dictionary
	ErrWordExists = DictionaryErr("word already exists")

	// ErrWordDoesNotExist is thrown when a new word is updated in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr implements the error interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary represents a map
type Dictionary map[string]string

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

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update changes the definition of an existing word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
