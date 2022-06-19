package main

import "testing"

func TestSearch(t *testing.T){
	
	phrase := "this is just a test"
	dictionary := Dictionary{"test": phrase}

	t.Run("known word",func(t *testing.T){
		got,_ := dictionary.Search("test")
		want := phrase

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
	dictionary := Dictionary{}
	word := "word"
	definition := "simple word"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word,definition)
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
