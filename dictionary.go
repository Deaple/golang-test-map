package main

// import (
// 	"errors"
// )

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrorNotFound = DictionaryErr("No definition found for the word.")
 	ErrorWordExists = DictionaryErr("This word already exists in dictionary.")
	ErrorWordDoesNotExists = DictionaryErr("Cannot update word because it does note exist.")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Update(word, newDefinition string) error{
	_, err := d.Search(word)

	switch err{
	case ErrorNotFound:
		return ErrorWordDoesNotExists
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Search(word string) (string,error){
	if d[word] == "" {
		return "", ErrorNotFound
	}
	return d[word], nil
}

//maps are pointer to runtime.hmap
//so we don't need to pass and pointer here
func (d Dictionary) Add(word, definition string) error{
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}
	
	return nil
}