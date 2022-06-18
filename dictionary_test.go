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
		want := "No definition found for the word."

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(t testing.TB, got, want string){
	t.Helper()

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}
}