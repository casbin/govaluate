package govaluate

import "sync"

type lexerStream struct {
	sourceString string
	source       []rune
	position     int
	length       int
}

var lexerStreamPool = sync.Pool{
	New: func() interface{} {
		return new(lexerStream)
	},
}

func newLexerStream(source string) *lexerStream {
	ret := lexerStreamPool.Get().(*lexerStream)
	if ret.source == nil {
		ret.source = make([]rune, 0, len(source))
	}
	for _, character := range source {
		ret.source = append(ret.source, character)
	}
	ret.sourceString = source
	ret.position = 0
	ret.length = len(ret.source)
	return ret
}

func (this *lexerStream) readCharacter() rune {
	character := this.source[this.position]
	this.position += 1
	return character
}

func (this *lexerStream) rewind(amount int) {
	this.position -= amount
}

func (this lexerStream) canRead() bool {
	return this.position < this.length
}

func (this *lexerStream) close() {
	this.source = this.source[:0]
	lexerStreamPool.Put(this)
}
