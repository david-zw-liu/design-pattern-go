package adapter

import (
	"testing"
)

func TestNewBook(t *testing.T) {
	book := NewBook("David", "Learning Go")

	author := book.GetAuthor()
	if author != "David" {
		t.Errorf("Expected author of book is \"David\", but got \"%v\"", author)
	}

	title := book.GetTitle()
	if title != "Learning Go" {
		t.Errorf("Expected author of book is \"Learning Go\", but got \"%v\"", title)
	}
}

func TestNewBookAdapter(t *testing.T) {
	book := NewBook("David", "Learning Go")
	ba := NewBookAdapter(book)

	ant := ba.getAuthorAndTitle()
	if ant != "Learning Go by David" {
		t.Errorf("Expected AutherAndTitle is \"Learning Go by David\", but got \"%v\"", ant)
	}
}
