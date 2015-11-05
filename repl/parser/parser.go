package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS

	IDENT

	GET
	SET
)

var eof = rune(0)

type Statement struct {
	Action Token
	Key    string
	Value  string
	Expiry int64
}

type Parser struct {
	scanner *Scanner
	buf     struct {
		tok Token
		lit string
		n   int
	}
}

func NewParser(input string) *Parser {
	r := bufio.NewReader(strings.NewReader(input))

	return &Parser{scanner: NewScanner(r)}
}

func (p *Parser) Parse() (Statement, error) {
	s := Statement{}

	tok, lit := p.scanIgnoreWhitespace()
	if tok != GET && tok != SET {
		return s, fmt.Errorf("found %q expected command", lit)
	}
	s.Action = tok

	tok, lit = p.scanIgnoreWhitespace()
	if tok != IDENT {
		return s, fmt.Errorf("found %q expectend key", lit)
	}
	s.Key = lit

	if s.Action == SET {
		tok, lit = p.scanIgnoreWhitespace()
		if tok != IDENT {
			return s, fmt.Errorf("found %q expectend expiry", lit)
		}

		i, err := strconv.ParseInt(lit, 10, 64)
		if err != nil {
			return s, fmt.Errorf("failed to parse %q as integer", lit)
		}
		s.Expiry = int64(i)
		s.Value = strings.Trim(p.scanToEnd(), " ")
	}

	return s, nil
}

func (p *Parser) scan() (Token, string) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	tok, lit := p.scanner.scan()

	p.buf.tok, p.buf.lit = tok, lit

	return tok, lit
}

func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) scanIgnoreWhitespace() (Token, string) {
	tok, lit := p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}

	return tok, lit
}

func (p *Parser) scanToEnd() string {
	return p.scanner.readToEnd()
}

type Scanner struct {
	reader *bufio.Reader
}

func NewScanner(r *bufio.Reader) *Scanner {
	return &Scanner{r}
}

func (s *Scanner) read() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *Scanner) unread() { _ = s.reader.UnreadRune() }

func (s *Scanner) scan() (Token, string) {
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isText(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) readToEnd() string {
	str := ""

	for {
		ch := s.read()

		if ch == eof {
			return str
		}

		str += string(ch)
	}
}

func (s *Scanner) scanWhitespace() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanIdent() (Token, string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isText(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case "GET":
		return GET, buf.String()
	case "SET":
		return SET, buf.String()
	}

	return IDENT, buf.String()
}

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

func isText(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}
