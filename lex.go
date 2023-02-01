package pic

import (
	"errors"
	"log"
)

type lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func newLexer(input string) *lexer {
	l := &lexer{input: input}
	l.readChar()
	return l
}

func (l *lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// Parser for a cobol picture string

type result struct {
	picType    string // N for number, A for string
	strLen     int
	intPartLen int
	decPartLen int
	sign       bool
	signLeft   bool
}

type parser struct {
	l   *lexer
	res result
}

func newParser(l *lexer) *parser {
	return &parser{l: l}
}

func (p *parser) parse() (result, error) {
	for {
		switch p.l.ch {
		case '9', 'S', 'V':
			p.parseNumber()
		case 'X', 'A':
			p.parseString()
		case 0:
			return p.res, nil
		default:
			return p.res, errors.New("INVALID_PIC")
		}
	}
}

func (p *parser) parseNumber() {
	var onDecimalPart bool
	p.res.picType = "N"

	if p.l.ch == 'S' {
		p.res.sign = true
		p.res.signLeft = true
		p.l.readChar()
	}

	if p.l.ch == '9' {
		p.res.intPartLen = 1
	}

	p.l.readChar()
	for p.l.ch == '9' || p.l.ch == 'V' || p.l.ch == '(' || p.l.ch == 'S' {
		if p.l.ch == '(' {
			var num int
			p.l.readChar()
			for p.l.ch != ')' {
				num = num*10 + int(p.l.ch-'0')
				p.l.readChar()
			}
			log.Println(num)
			if onDecimalPart {
				p.res.decPartLen += num - 1
			} else {
				p.res.intPartLen += num - 1
			}
			p.l.readChar()
		}
		if p.l.ch == 'V' {
			p.l.readChar()
			onDecimalPart = true
		}
		if p.l.ch == '9' {
			if onDecimalPart {
				p.res.decPartLen += 1
			} else {
				p.res.intPartLen += 1
			}
		}
		if p.l.ch == 'S' {
			p.res.sign = true
			p.res.signLeft = false
		}
		p.l.readChar()
	}
}

func (p *parser) parseString() {
	p.res.picType = "A"
	p.l.readChar()
	for p.l.ch == 'X' || p.l.ch == 'A' || p.l.ch == '(' {
		if p.l.ch == '(' {
			var num int
			p.l.readChar()
			for p.l.ch != ')' {
				num = num*10 + int(p.l.ch-'0')
				p.l.readChar()
			}
			p.res.strLen += num
			p.l.readChar()
		}
		if p.l.ch == 'X' || p.l.ch == 'A' {
			p.res.strLen += 1
			p.l.readChar()
		}
	}
}
