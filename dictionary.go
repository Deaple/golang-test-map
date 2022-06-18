package main

import (
	"errors"
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string,error){
	if d[word] == "" {
		return "", errors.New("No definition found for the word.")
	}
	return d[word], nil
}