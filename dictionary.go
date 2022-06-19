package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrorNotFound = errors.New("No definition found for the word.")

func (d Dictionary) Search(word string) (string,error){
	if d[word] == "" {
		return "", ErrorNotFound
	}
	return d[word], nil
}

//maps are pointer to runtime.hmap
//so we don't need to pass and pointer here
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}