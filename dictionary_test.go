package main

import "testing"


func TestSearch(t *testing.T){
	
	definition := "this is just a test"
	dictionary := Dictionary{"test": definition}

	t.Run("known word",func(t *testing.T){
		got,_ := dictionary.Search("test")
		want := definition

		assertStrings(t, got, want)
	})

	t.Run("unknown word",func(t *testing.T){
		
		_,err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, err, ErrorNotFound)
	})
}

func TestAdd(t *testing.T){
	
	t.Run("new word", func(t *testing.T){
		dictionary := Dictionary{}
		word := "word"
		definition := "simple word"
		dictionary.Add(word, definition)

		err := dictionary.Add(word, definition)
	
		assertErrors(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, word,definition)
	})

	t.Run("existing word", func(t *testing.T){
		word := "word"
		definition := "simple word"
		dictionary := Dictionary{word: definition} 

		err := dictionary.Add(word, "new definition")

		assertErrors(t, err, ErrorWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T){

	t.Run("existing word", func(t *testing.T){
		word := "new"
		definition := "a definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "updated definition"
	
		err := dictionary.Update(word, newDefinition)
	
		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("word that does not exist", func(t *testing.T){
		word := "new"
		definition := "a definition"
		dictionary := Dictionary{}
	
		err := dictionary.Update(word, definition)
	
		assertErrors(t, err, ErrorWordDoesNotExists)
	})
}

func assertErrors(t testing.TB, got, err error){
	t.Helper()

	if got != err {
		t.Errorf("got %q, wanted %q", got , err)
	}
}

func assertStrings(t testing.TB, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string){
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil{
		t.Fatal("should find added word:",err)
	}

	if definition != got{
		t.Errorf("got %q, wanted %q",got,definition)
	}
}
