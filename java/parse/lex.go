package parse

import (
	"bytes"
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
		state:  lexPackage,
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

func (l *lexer) read() []byte {
	return l.src[l.start:l.pos]
}

func (l *lexer) readahead(n int) []byte {
	pos := l.pos
	for i := 0; i < n; i++ {
		l.next()
	}
	lexeme := l.src[pos:l.pos]
	l.pos = pos
	return lexeme
}

func lexPackage(l *lexer) stateFn {
	if !bytes.Equal(l.readahead(7), []byte("package")) {
		return lexImports
	}

	for {
		switch l.r {
		case eof:
			return nil
		case ';':
			l.emit(tokenPackage)
			l.discard()
			return lexImports
		default:
			l.next()
		}
	}
}

func lexImports(l *lexer) stateFn {
	for {
		switch l.r {
		case eof:
			return nil
		case ';':
			if l.start == l.pos {
				l.next()
				continue
			}
			l.emit(tokenImport)
			l.reset()
		case '{':
			return lexClass
		default:
			l.next()
		}
	}
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
			// check if static class initializer
			s := bytes.TrimSpace(l.src[l.start : l.pos-1])
			if bytes.Equal(s, []byte("static")) {
				l.emit(tokenInit)
				l.discard()
				return lexMethodBody(l)
			}

			// an inner class or interface
			return lexClass
		case '}':
			// end of class body
			l.discard()
			l.emit(tokenClassEnd)
		case '@':
			if bytes.Equal(l.readahead(9), []byte("interface")) {
				// declared annotation
				continue
			}
			l.reset()
			lexAnnotation(l)
			continue
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

func lexInit(l *lexer) stateFn {
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
	l.emit(tokenMethod)
	l.reset()
	lexParameters(l)

	for {
		switch l.r {
		case eof:
			// l.emit(tokenMethod)
			return nil
		case ';':
			// for abstract and interface method defs
			// l.emit(tokenMethod)
			l.discard()
			return lexClassBody
		case '{':
			// l.emit(tokenMethod)
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
		case '\'':
			lexChar(l)
		case '"':
			lexString(l)
		case '}':
			l.discard()
			return lexClassBody
		default:
			l.next()
		}
	}
}

func lexChar(l *lexer) {
	l.next()
	for {
		switch l.r {
		case eof:
			return
		case '\\':
			l.next()
			l.next()
		case '\'':
			l.next()
			return
		default:
			l.next()
		}
	}
}

func lexString(l *lexer) {
	l.next()
	for {
		switch l.r {
		case eof:
			return
		case '\\':
			l.next()
			l.next()
		case '"':
			l.next()
			return
		default:
			l.next()
		}
	}
}

func lexAnnotation(l *lexer) {
	for {
		switch l.r {
		case eof:
			l.emit(tokenAnnotation)
			return
		case '(':
			l.emit(tokenAnnotation)
			l.reset()
			lexParameters(l)
		case ')':
			l.discard()
			return
		case '\n': // TODO(d) more of a regress on whats allowed in java
			l.emit(tokenAnnotation)
			l.reset()
			return
		default:
			l.next()
		}
	}
}

func lexParameters(l *lexer) {
	for {
		switch l.r {
		case eof:
			l.emit(tokenParameter)
			return
		case '"':
			lexString(l)
		case '\'':
			lexChar(l)
		case ',':
			l.emit(tokenParameter)
			l.discard()
		case ')':
			if !bytes.Equal(l.read(), []byte(")")) {
				l.emit(tokenParameter)
			}
			l.reset()
			return
		default:
			l.next()
		}
	}
}
