package token

type IToken interface {
	Valid() bool
}

type Token struct {
	Tag    string
	Id     string
	Level  int64
	Status int64
	Data   interface{}
}

func (t *Token) Valid() bool {
	return len(t.Id) != 0
}
