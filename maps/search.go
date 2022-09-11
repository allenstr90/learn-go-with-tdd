package maps

const (
	WordNotFoundException      = DictErr("Word not found")
	WordAlreadyExistsException = DictErr("Word already exists")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

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

func (d Dict) Update(key, value string) {
	d[key] = value
}

func (d Dict) Delete(key string) {
	delete(d, key)
}
