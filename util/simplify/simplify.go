// This file was generated with the following command:
// ["/Users/quarnster/code/go/bin/pegparser", "-peg=../util/simplify/simplify.peg", "-notest", "-ignore=NewLine,EOF", "-testfile=", "-outpath", "../util/simplify/", "-generator=go"]

package simplify

import (
	. "github.com/quarnster/parser"
	"github.com/quarnster/util/text"
)

type SIMPLIFY struct {
	ParserData  Reader
	IgnoreRange text.Region
	Root        Node
	LastError   int
}

func (p *SIMPLIFY) RootNode() *Node {
	return &p.Root
}

func (p *SIMPLIFY) SetData(data string) {
	p.ParserData = NewReader(data)
	p.Root = Node{Name: "SIMPLIFY", P: p}
	p.IgnoreRange = text.Region{}
	p.LastError = 0
}

func (p *SIMPLIFY) Parse(data string) bool {
	p.SetData(data)
	ret := p.realParse()
	p.Root.UpdateRange()
	return ret
}

func (p *SIMPLIFY) Data(start, end int) string {
	return p.ParserData.Substring(start, end)
}

func (p *SIMPLIFY) Error() Error {
	errstr := ""
	line, column := p.ParserData.LineCol(p.LastError)

	if p.LastError == p.ParserData.Len() {
		errstr = "Unexpected EOF"
	} else {
		p.ParserData.Seek(p.LastError)
		if r := p.ParserData.Read(); r == '\r' || r == '\n' {
			errstr = "Unexpected new line"
		} else {
			errstr = "Unexpected " + string(r)
		}
	}
	return NewError(line, column, errstr)
}

func (p *SIMPLIFY) realParse() bool {
	return p.Simplified()
}
func (p *SIMPLIFY) Simplified() bool {
	// Simplified		<-	(Line / Unknown)+ EOF
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			{
				save := p.ParserData.Pos()
				accept = p.Line()
				if !accept {
					accept = p.Unknown()
					if !accept {
					}
				}
				if !accept {
					p.ParserData.Seek(save)
				}
			}
			if !accept {
				p.ParserData.Seek(save)
			} else {
				for accept {
					{
						save := p.ParserData.Pos()
						accept = p.Line()
						if !accept {
							accept = p.Unknown()
							if !accept {
							}
						}
						if !accept {
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
		}
		if accept {
			accept = p.EOF()
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Simplified"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *SIMPLIFY) Unknown() bool {
	// Unknown			<-	(!'\n' .)+ '\n'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			{
				save := p.ParserData.Pos()
				s := p.ParserData.Pos()
				if p.ParserData.Read() != '\n' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				p.ParserData.Seek(s)
				p.Root.Discard(s)
				accept = !accept
				if accept {
					if p.ParserData.Pos() >= p.ParserData.Len() {
						accept = false
					} else {
						p.ParserData.Read()
						accept = true
					}
					if accept {
					}
				}
				if !accept {
					if p.LastError < p.ParserData.Pos() {
						p.LastError = p.ParserData.Pos()
					}
					p.ParserData.Seek(save)
				}
			}
			if !accept {
				p.ParserData.Seek(save)
			} else {
				for accept {
					{
						save := p.ParserData.Pos()
						s := p.ParserData.Pos()
						if p.ParserData.Read() != '\n' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						p.ParserData.Seek(s)
						p.Root.Discard(s)
						accept = !accept
						if accept {
							if p.ParserData.Pos() >= p.ParserData.Len() {
								accept = false
							} else {
								p.ParserData.Read()
								accept = true
							}
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
				}
				accept = true
			}
		}
		if accept {
			if p.ParserData.Read() != '\n' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Unknown"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *SIMPLIFY) Line() bool {
	// Line			<-	(Word ' '?)* (Scope / (';' NewLine))
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				{
					save := p.ParserData.Pos()
					accept = p.Word()
					if accept {
						if p.ParserData.Read() != ' ' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						accept = true
						if accept {
						}
					}
					if !accept {
						if p.LastError < p.ParserData.Pos() {
							p.LastError = p.ParserData.Pos()
						}
						p.ParserData.Seek(save)
					}
				}
			}
			accept = true
		}
		if accept {
			{
				save := p.ParserData.Pos()
				accept = p.Scope()
				if !accept {
					{
						save := p.ParserData.Pos()
						if p.ParserData.Read() != ';' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						if accept {
							accept = p.NewLine()
							if accept {
							}
						}
						if !accept {
							if p.LastError < p.ParserData.Pos() {
								p.LastError = p.ParserData.Pos()
							}
							p.ParserData.Seek(save)
						}
					}
					if !accept {
					}
				}
				if !accept {
					p.ParserData.Seek(save)
				}
			}
			if accept {
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Line"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *SIMPLIFY) Scope() bool {
	// Scope			<-	'{' NewLine Line+ '}' NewLine
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		if p.ParserData.Read() != '{' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		if accept {
			accept = p.NewLine()
			if accept {
				{
					save := p.ParserData.Pos()
					accept = p.Line()
					if !accept {
						p.ParserData.Seek(save)
					} else {
						for accept {
							accept = p.Line()
						}
						accept = true
					}
				}
				if accept {
					if p.ParserData.Read() != '}' {
						p.ParserData.UnRead()
						accept = false
					} else {
						accept = true
					}
					if accept {
						accept = p.NewLine()
						if accept {
						}
					}
				}
			}
		}
		if !accept {
			if p.LastError < p.ParserData.Pos() {
				p.LastError = p.ParserData.Pos()
			}
			p.ParserData.Seek(save)
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Scope"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *SIMPLIFY) Word() bool {
	// Word			<-	(![{}; \n] .)+
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			s := p.ParserData.Pos()
			{
				accept = false
				c := p.ParserData.Read()
				if c == '{' || c == '}' || c == ';' || c == ' ' || c == '\n' {
					accept = true
				} else {
					p.ParserData.UnRead()
				}
			}
			p.ParserData.Seek(s)
			p.Root.Discard(s)
			accept = !accept
			if accept {
				if p.ParserData.Pos() >= p.ParserData.Len() {
					accept = false
				} else {
					p.ParserData.Read()
					accept = true
				}
				if accept {
				}
			}
			if !accept {
				if p.LastError < p.ParserData.Pos() {
					p.LastError = p.ParserData.Pos()
				}
				p.ParserData.Seek(save)
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		} else {
			for accept {
				{
					save := p.ParserData.Pos()
					s := p.ParserData.Pos()
					{
						accept = false
						c := p.ParserData.Read()
						if c == '{' || c == '}' || c == ';' || c == ' ' || c == '\n' {
							accept = true
						} else {
							p.ParserData.UnRead()
						}
					}
					p.ParserData.Seek(s)
					p.Root.Discard(s)
					accept = !accept
					if accept {
						if p.ParserData.Pos() >= p.ParserData.Len() {
							accept = false
						} else {
							p.ParserData.Read()
							accept = true
						}
						if accept {
						}
					}
					if !accept {
						if p.LastError < p.ParserData.Pos() {
							p.LastError = p.ParserData.Pos()
						}
						p.ParserData.Seek(save)
					}
				}
			}
			accept = true
		}
	}
	end := p.ParserData.Pos()
	if accept {
		node := p.Root.Cleanup(start, end)
		node.Name = "Word"
		node.P = p
		node.Range = node.Range.Clip(p.IgnoreRange)
		p.Root.Append(node)
	} else {
		p.Root.Discard(start)
	}
	if p.IgnoreRange.A >= end || p.IgnoreRange.B <= start {
		p.IgnoreRange = text.Region{}
	}
	return accept
}

func (p *SIMPLIFY) NewLine() bool {
	// NewLine			<-	'\n'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	if p.ParserData.Read() != '\n' {
		p.ParserData.UnRead()
		accept = false
	} else {
		accept = true
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *SIMPLIFY) EOF() bool {
	// EOF				<-  !.
	accept := false
	accept = true
	start := p.ParserData.Pos()
	s := p.ParserData.Pos()
	if p.ParserData.Pos() >= p.ParserData.Len() {
		accept = false
	} else {
		p.ParserData.Read()
		accept = true
	}
	p.ParserData.Seek(s)
	p.Root.Discard(s)
	accept = !accept
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}
