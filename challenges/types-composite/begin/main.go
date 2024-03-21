// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct{
	name string
}

// define a book type with a title and author
type book struct{
title string
author author
}
// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book){
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) []book {
	return l[authorName]
}

func main() {
	// create a new library
	booklib := library{}
au := author{name: "Aurthor"}
	// add 2 books to the library by the same author
	booklib.addBook(book{title: "The book1", author: au})

	// add 1 book to the library by a different author
	booklib.addBook(book{title: "The book2", author: author{name: "Carol"}})

	// dump the library with spew
	spew.Dump(booklib)

	// lookup books by known author in the library
	books := booklib.lookupByAuthor(au.name)
	fmt.Println(books)

	// print out the first book's title and its author's name
if len(books)>0{
	fmt.Println(books[0].title, books[0].author.name)
}

}
