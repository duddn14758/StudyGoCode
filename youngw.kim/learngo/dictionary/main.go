package main

import (
	"fmt"
	"github.com/youngw.kim/learngo/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "second"
	definition := "Second word"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("word '", word, "' added")
	}
	/*
		definition, err := dictionary.Search("second")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/
	second, _ := dictionary.Search(word)
	err = dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(second)
	}

	err = dictionary.Update(word, "updated Second")
	if err != nil {
		fmt.Println(err)
	}
	uWord, _ := dictionary.Search(word)
	fmt.Println(uWord)

	dictionary.Delete(word)
	uWord, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(uWord)
	}
}
