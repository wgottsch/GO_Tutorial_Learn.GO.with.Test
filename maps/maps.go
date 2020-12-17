package maps

type Dictionary map[string]string

// Refactor to constant errors
// before:

//var (
//	ErrNotFound = errors.New("could not find the word you were looking for")
//	errWordExist = errors.New("cannot add word, because it already exist")
//)
// after:

const (
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	errWordExist = DictionaryErr("cannot add word, because it already exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return errWordExist
	default:
		return err
	}
	return nil
}
