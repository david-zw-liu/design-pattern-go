package prototype

type BookPrototype struct {
	title string
	topic string
}

func (bp *BookPrototype) SetTitle(title string) {
	bp.title = title
}

func (bp *BookPrototype) GetTitle() string {
	return bp.title
}

func (bp *BookPrototype) GetTopic() string {
	return bp.topic
}

type SQLBookPrototype struct {
	BookPrototype
}

func CreateSQLBookPrototype() SQLBookPrototype {
	return SQLBookPrototype{BookPrototype: BookPrototype{topic: "SQL"}}
}
