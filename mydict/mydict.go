package mydict

import "errors"



type Dictionary map[string]string

var errNoWord = errors.New("no word")
var errWordExists = errors.New("word already exists")
var errCantUpdate = errors.New("can't update non-existing word")

// we can add methods to types

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", errNoWord
	} 
	return value, nil
	
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	if err == errNoWord {
		d[word] = definition
		return nil
	} 

	return errWordExists
}

func (d Dictionary) Update(word string, newDef string) error {
	_, err := d.Search(word)
	if err == errNoWord {
		return errCantUpdate
	} 
	d[word] = newDef
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}