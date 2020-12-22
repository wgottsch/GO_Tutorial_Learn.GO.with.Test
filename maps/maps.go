package maps

import "github.com/pkg/errors"

type Dictionary map[string]string

const (
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	ErrWordExist = DictionaryErr("cannot add word, because it already exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.WithStack(ErrNotFound)
	}
	return definition, nil
}
func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch errors.Cause(err) {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return errors.WithStack(ErrWordExist)
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch errors.Cause(err) {
	case ErrNotFound:
		return errors.WithMessage(ErrNotFound, "Try an d.Update()")
	case nil:
		d[word] = definition
	default:
		return errors.WithMessage(err, "unknown Error: Try an d.Update()")
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch errors.Cause(err) {
	case ErrNotFound:
		return errors.WithMessage(ErrNotFound, "Try an d.Delete()")
	case nil:
		delete(d, word)
	default:
		return errors.WithMessage(err, "unknown Error: Try an d.Delete()")
	}
	return nil
}
