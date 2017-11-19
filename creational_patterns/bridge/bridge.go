package bridge

import (
	"strings"
)

// Impliement show functions of Book
type BookShowImp struct {
	GetAuthor func(string) string
	GetTitle  func(string) string
}

type Book struct {
	Author string
	Title  string
	Imp    *BookShowImp
}

func (b *Book) GetAuthor() string {
	return b.Imp.GetAuthor(b.Author)
}

func (b *Book) GetTitle() string {
	return b.Imp.GetTitle(b.Title)
}

func NewCapsImp() *BookShowImp {
	return &BookShowImp{
		GetAuthor: func(author string) string {
			return strings.ToUpper(author)
		},
		GetTitle: func(title string) string {
			return strings.ToUpper(title)
		},
	}
}

func NewStarsImp() *BookShowImp {
	return &BookShowImp{
		GetAuthor: func(author string) string {
			return strings.Replace(author, " ", "*", -1)
		},
		GetTitle: func(title string) string {
			return strings.Replace(title, " ", "*", -1)
		},
	}
}

func NewBook(author, title, imp_type string) *Book {
	var imp *BookShowImp

	switch imp_type {
	case "STARS":
		imp = NewStarsImp()
	case "CAPS":
		imp = NewCapsImp()
	}

	return &Book{Author: author, Title: title, Imp: imp}
}
