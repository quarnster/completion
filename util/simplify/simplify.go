// This file was generated with the following command:
// ["/Users/quarnster/code/go/bin/pegparser", "-peg=../util/simplify/simplify.peg", "-notest", "-ignore=NewLine,EOF,Spacing,End,Impl,Contents,Simplified", "-testfile=", "-outpath", "../util/simplify/", "-generator=go"]

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
	// Simplified		<-	Contents+ EOF
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			accept = p.Contents()
			if !accept {
				p.ParserData.Seek(save)
			} else {
				for accept {
					accept = p.Contents()
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
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *SIMPLIFY) Contents() bool {
	// Contents		<-	(Class / Method / Variable / Unknown)
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		accept = p.Class()
		if !accept {
			accept = p.Method()
			if !accept {
				accept = p.Variable()
				if !accept {
					accept = p.Unknown()
					if !accept {
					}
				}
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *SIMPLIFY) Access() bool {
	// Access			<-	("public" / "protected" / "private" / "internal" / "static") Spacing
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != 'p' || p.ParserData.Read() != 'u' || p.ParserData.Read() != 'b' || p.ParserData.Read() != 'l' || p.ParserData.Read() != 'i' || p.ParserData.Read() != 'c' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if !accept {
				{
					accept = true
					s := p.ParserData.Pos()
					if p.ParserData.Read() != 'p' || p.ParserData.Read() != 'r' || p.ParserData.Read() != 'o' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'e' || p.ParserData.Read() != 'c' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'e' || p.ParserData.Read() != 'd' {
						p.ParserData.Seek(s)
						accept = false
					}
				}
				if !accept {
					{
						accept = true
						s := p.ParserData.Pos()
						if p.ParserData.Read() != 'p' || p.ParserData.Read() != 'r' || p.ParserData.Read() != 'i' || p.ParserData.Read() != 'v' || p.ParserData.Read() != 'a' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'e' {
							p.ParserData.Seek(s)
							accept = false
						}
					}
					if !accept {
						{
							accept = true
							s := p.ParserData.Pos()
							if p.ParserData.Read() != 'i' || p.ParserData.Read() != 'n' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'e' || p.ParserData.Read() != 'r' || p.ParserData.Read() != 'n' || p.ParserData.Read() != 'a' || p.ParserData.Read() != 'l' {
								p.ParserData.Seek(s)
								accept = false
							}
						}
						if !accept {
							{
								accept = true
								s := p.ParserData.Pos()
								if p.ParserData.Read() != 's' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'a' || p.ParserData.Read() != 't' || p.ParserData.Read() != 'i' || p.ParserData.Read() != 'c' {
									p.ParserData.Seek(s)
									accept = false
								}
							}
							if !accept {
							}
						}
					}
				}
			}
			if !accept {
				p.ParserData.Seek(save)
			}
		}
		if accept {
			accept = p.Spacing()
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
		node.Name = "Access"
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

func (p *SIMPLIFY) Class() bool {
	// Class			<-	Access* "class " (Word ' '?)+  Impl
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				accept = p.Access()
			}
			accept = true
		}
		if accept {
			{
				accept = true
				s := p.ParserData.Pos()
				if p.ParserData.Read() != 'c' || p.ParserData.Read() != 'l' || p.ParserData.Read() != 'a' || p.ParserData.Read() != 's' || p.ParserData.Read() != 's' || p.ParserData.Read() != ' ' {
					p.ParserData.Seek(s)
					accept = false
				}
			}
			if accept {
				{
					save := p.ParserData.Pos()
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
					if !accept {
						p.ParserData.Seek(save)
					} else {
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
				}
				if accept {
					accept = p.Impl()
					if accept {
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
		node.Name = "Class"
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

func (p *SIMPLIFY) Method() bool {
	// Method			<-	Access* (Word ' '?)+ Arguments Impl
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				accept = p.Access()
			}
			accept = true
		}
		if accept {
			{
				save := p.ParserData.Pos()
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
				if !accept {
					p.ParserData.Seek(save)
				} else {
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
			}
			if accept {
				accept = p.Arguments()
				if accept {
					accept = p.Impl()
					if accept {
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
		node.Name = "Method"
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

func (p *SIMPLIFY) Impl() bool {
	// Impl			<-	(';' End) / (End Body?)
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			save := p.ParserData.Pos()
			if p.ParserData.Read() != ';' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.End()
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
			{
				save := p.ParserData.Pos()
				accept = p.End()
				if accept {
					accept = p.Body()
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
			if !accept {
			}
		}
		if !accept {
			p.ParserData.Seek(save)
		}
	}
	if accept && start != p.ParserData.Pos() {
		if start < p.IgnoreRange.A || p.IgnoreRange.A == 0 {
			p.IgnoreRange.A = start
		}
		p.IgnoreRange.B = p.ParserData.Pos()
	}
	return accept
}

func (p *SIMPLIFY) Variable() bool {
	// Variable		<-	Access* (Word ' '?)+ ';' End
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				accept = p.Access()
			}
			accept = true
		}
		if accept {
			{
				save := p.ParserData.Pos()
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
				if !accept {
					p.ParserData.Seek(save)
				} else {
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
			}
			if accept {
				if p.ParserData.Read() != ';' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				if accept {
					accept = p.End()
					if accept {
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
		node.Name = "Variable"
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

func (p *SIMPLIFY) Arguments() bool {
	// Arguments		<-	'(' (Word ([ ,]*))* ')'
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		if p.ParserData.Read() != '(' {
			p.ParserData.UnRead()
			accept = false
		} else {
			accept = true
		}
		if accept {
			{
				accept = true
				for accept {
					{
						save := p.ParserData.Pos()
						accept = p.Word()
						if accept {
							{
								accept = true
								for accept {
									{
										accept = false
										c := p.ParserData.Read()
										if c == ' ' || c == ',' {
											accept = true
										} else {
											p.ParserData.UnRead()
										}
									}
								}
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
			if accept {
				if p.ParserData.Read() != ')' {
					p.ParserData.UnRead()
					accept = false
				} else {
					accept = true
				}
				if accept {
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
		node.Name = "Arguments"
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
	// Unknown			<-	(!'\n' .)+ End
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
			accept = p.End()
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

func (p *SIMPLIFY) Body() bool {
	// Body			<-	Spacing* '{' End (!'}' Contents)* '}' End
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				accept = p.Spacing()
			}
			accept = true
		}
		if accept {
			if p.ParserData.Read() != '{' {
				p.ParserData.UnRead()
				accept = false
			} else {
				accept = true
			}
			if accept {
				accept = p.End()
				if accept {
					{
						accept = true
						for accept {
							{
								save := p.ParserData.Pos()
								s := p.ParserData.Pos()
								if p.ParserData.Read() != '}' {
									p.ParserData.UnRead()
									accept = false
								} else {
									accept = true
								}
								p.ParserData.Seek(s)
								p.Root.Discard(s)
								accept = !accept
								if accept {
									accept = p.Contents()
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
						if p.ParserData.Read() != '}' {
							p.ParserData.UnRead()
							accept = false
						} else {
							accept = true
						}
						if accept {
							accept = p.End()
							if accept {
							}
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
		node.Name = "Body"
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

func (p *SIMPLIFY) Spacing() bool {
	// Spacing			<-	' '
	accept := false
	accept = true
	start := p.ParserData.Pos()
	if p.ParserData.Read() != ' ' {
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

func (p *SIMPLIFY) Word() bool {
	// Word			<-	(![(){}; \n] .)+
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
				if c == '(' || c == ')' || c == '{' || c == '}' || c == ';' || c == ' ' || c == '\n' {
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
						if c == '(' || c == ')' || c == '{' || c == '}' || c == ';' || c == ' ' || c == '\n' {
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

func (p *SIMPLIFY) End() bool {
	// End				<-	Spacing* (NewLine / EOF)
	accept := false
	accept = true
	start := p.ParserData.Pos()
	{
		save := p.ParserData.Pos()
		{
			accept = true
			for accept {
				accept = p.Spacing()
			}
			accept = true
		}
		if accept {
			{
				save := p.ParserData.Pos()
				accept = p.NewLine()
				if !accept {
					accept = p.EOF()
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
