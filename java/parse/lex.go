package parse

import (
	"unicode/utf8"
)

const eof = -1

type stateFn func(*lexer) stateFn

type lexer struct {
	src    []byte
	start  int
	pos    int
	r      rune
	state  stateFn
	tokens chan token
}

func lex(src []byte) *lexer {
	// TODO some sort of analysis should take place to determine the start state.
	// At minimum for now, this should start with a lexPackage for imports, etc
	return &lexer{
		src:    src,
		state:  lexClass,
		tokens: make(chan token),
	}
}

func (l *lexer) run() {
	for l.state != nil {
		l.state = l.state(l)
	}
	close(l.tokens)
}

func (l *lexer) emit(typ tokenType) {
	l.tokens <- token{typ, l.start, l.pos}
}

func (l *lexer) next() {
	if l.pos >= len(l.src) {
		l.r = eof
		return
	}
	r, width := utf8.DecodeRune(l.src[l.pos:])
	l.r = r
	l.pos += width
}

func (l *lexer) reset() {
	l.start = l.pos
}

func (l *lexer) discard() {
	l.next()
	l.reset()
}

func lexClass(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			// emit what we have and let parser decide if it's enough
			l.emit(tokenClass)
			return nil
		case '{':
			l.emit(tokenClass)
			l.discard()
			// TODO should allow reuse of this method for inner and outer classes
			return lexClassBody
		default:
			l.next()
		}
	}
}

func lexClassBody(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			return nil
		case '{':
			// an inner class, interface, or class initializer
			return lexClass
		case '}':
			// end of class body
			l.discard()
			l.emit(tokenClassEnd)
		case '=', ';':
			// '=' assignment to class member which may lead to an anonymous class
			// ';' terminates the field
			return lexField
		case '(':
			return lexMethod
		default:
			l.next()
		}
	}
}

func lexField(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			l.emit(tokenField)
			return nil
		case '=':
			l.emit(tokenField)
			l.discard()
			return lexFieldAssignment
		case ';':
			l.emit(tokenField)
			l.discard()
			return lexClassBody
		default:
			l.next()
		}
	}
}

func lexFieldAssignment(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			return nil
		case ';':
			l.discard()
			return lexClassBody
		case '(':
			// found anonymous inner class
			// TODO handle
			l.next()
		default:
			l.next()
		}
	}
}

func lexMethod(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			l.emit(tokenMethod)
			return nil
		case ';':
			// for abstract and interface method defs
			l.emit(tokenMethod)
			l.discard()
			return lexClassBody
		case '{':
			l.emit(tokenMethod)
			l.discard()
			return lexMethodBody
		default:
			l.next()
		}
	}
}

func lexMethodBody(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			return nil
		case '}':
			l.discard()
			return lexClassBody
		default:
			l.next()
		}
	}
}
