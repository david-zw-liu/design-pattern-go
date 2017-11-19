package bridge

import "testing"

func TestNewBook(t *testing.T) {
	capsBook := NewBook("Larry Truett", "Go for Cats", "CAPS")
	if author := capsBook.GetAuthor(); author != "LARRY TRUETT" {
		t.Errorf("Expected author of book is \"LARRY TRUETT\", but got \"%v\"", author)
	}
	if title := capsBook.GetTitle(); title != "GO FOR CATS" {
		t.Errorf("Expected author of book is \"Go for Cats\", but got \"%v\"", title)
	}

	starsBook := NewBook("Larry Truett", "Go for Cats", "STARS")
	if author := starsBook.GetAuthor(); author != "Larry*Truett" {
		t.Errorf("Expected author of book is \"Larry*Truett\", but got \"%v\"", author)
	}
	if title := starsBook.GetTitle(); title != "Go*for*Cats" {
		t.Errorf("Expected author of book is \"Go*for*Cats\", but got \"%v\"", title)
	}

}
