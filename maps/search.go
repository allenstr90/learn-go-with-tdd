package maps

import (
	"errors"
)

var (
	WordNotFoundException      = errors.New("Word not found")
	WordAlreadyExistsException = errors.New("Word already exists")
)

func Search(dict map[string]string, k string) string {
	return dict[k]
}

type Dict map[string]string

func (d Dict) Search(k string) (string, error) {
	word, ok := d[k]
	if !ok {
		return "", WordNotFoundException
	}
	return word, nil
}

func (d Dict) Add(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case WordNotFoundException:
		d[k] = v
	case nil:
		return WordAlreadyExistsException
	default:
		return err
	}

	return nil
}
