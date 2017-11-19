package adapter

import "fmt"

type Book struct {
	Author string
	Title  string
}

func (b *Book) GetAuthor() string {
	return b.Author
}

func (b *Book) GetTitle() string {
	return b.Title
}

func NewBook(author, title string) *Book {
	return &Book{Author: author, Title: title}
}

type BookAdapter struct {
	*Book
}

func (ba *BookAdapter) getAuthorAndTitle() string {
	return fmt.Sprintf("%v by %v", ba.Book.GetTitle(), ba.Book.GetAuthor())
}

func NewBookAdapter(book *Book) *BookAdapter {
	return &BookAdapter{Book: book}
}
