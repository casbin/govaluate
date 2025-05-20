package govaluate

import (
	"sync"
	"unicode/utf8"
)

type lexerStream struct {
	sourceString string
	source       []rune
	strPosition  int
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
	ret.strPosition = 0
	ret.length = len(ret.source)
	return ret
}

func (this *lexerStream) readCharacter() rune {
	character := this.source[this.position]
	this.position += 1
	this.strPosition += utf8.RuneLen(character)
	return character
}

func (this *lexerStream) rewind(amount int) {
	if amount < 0 {
		this.position -= amount
		this.strPosition -= amount
	}
	strAmount := 0
	for i := 0; i < amount; i++ {
		if this.position >= this.length {
			strAmount += 1
			this.position -= 1
			continue
		}
		strAmount += utf8.RuneLen(this.source[this.position])
		this.position -= 1
	}
	this.strPosition -= strAmount
}

func (this lexerStream) canRead() bool {
	return this.position < this.length
}

func (this *lexerStream) close() {
	this.source = this.source[:0]
	lexerStreamPool.Put(this)
}
