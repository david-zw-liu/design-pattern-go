package prototype

import "testing"

func TestBookPrototype(t *testing.T) {
	sqlBookPrototype := CreateSQLBookPrototype()
	book1 := sqlBookPrototype // Golang assign create copy
	book2 := sqlBookPrototype
	book1.SetTitle("SQL For Cats")
	book2.SetTitle("OReilly Learning SQL")

	if title := book1.GetTitle(); title != "SQL For Cats" {
		t.Errorf("Expected that title is \"SQL For Cats\", but got \"%v\"", title)
	}

	if title := book2.GetTitle(); title != "OReilly Learning SQL" {
		t.Errorf("Expected that title is \"OReilly Learning SQL\", but got \"%v\"", title)
	}

	if topic := book2.GetTopic(); topic != "SQL" {
		t.Errorf("Expected that topic of book1 is \"SQL\", but got \"%v\"", topic)
	}

	if topic := book2.GetTopic(); topic != "SQL" {
		t.Errorf("Expected that topic of book2 is \"SQL\", but got \"%v\"", topic)
	}
}
