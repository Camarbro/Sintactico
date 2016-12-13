package camargo

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	// If we see a digit then consume as a number.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwis
	// e read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	//case '*':
	//	return ASTERISK, string(ch)
	case ',':
		return COMMA, string(ch)
	case '[':
		return AgruppadorDER, string(ch)
	case ']':
		return AgruppadorISQ, string(ch)
		//digit
	case '0':
		return NUMEROS, string(ch)
	case '1':
		return NUMEROS, string(ch)
	case '2':
		return NUMEROS, string(ch)
	case '3':
		return NUMEROS, string(ch)
	case '4':
		return NUMEROS, string(ch)
	case '5':
		return NUMEROS, string(ch)
	case '6':
		return NUMEROS, string(ch)
	case '7':
		return NUMEROS, string(ch)
	case '8':
		return NUMEROS, string(ch)
	case '9':
		return NUMEROS, string(ch)

	case ':':
		return DosPunt, string(ch)
	case '(':
		return ParDer, string(ch)
	case ')':
		return ParIsq, string(ch)
	case ';':
		return PuntCom, string(ch)
	case '{':
		return LlaveDer, string(ch)
	case '}':
		return LlaveIsq, string(ch)
	case '\n':
		return SaltoDeL, string(ch)
	//case '=':
	//	return Asignaicon, string(ch)
	case '+':
		return OPREL, string(ch)
	case '-':
		return OPREL, string(ch)
	case '/':
		return OPREL, string(ch)
	case '*':
		return OPREL, string(ch)
	case '=':
		return OPREL, string(ch)
	case '>':
		return OPREL, string(ch)
	case '<':
		return OPREL, string(ch)
	case '&':
		return OPREL, string(ch)
	}

	return ILLEGAL, string(ch)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
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

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {
	case "SELECT":
		return SELECT, buf.String()
	case "FROM":
		return FROM, buf.String()
	case "CAMARIF":
		return CAMARIF, buf.String()
	case "CAMARELSE":
		return CAMARELSE, buf.String()
	case "CAMAREND":
		return CAMAREND, buf.String()
	case "CAMARSWITCH":
		return CAMARSWITCH, buf.String()
	case "CAMARMAIN":
		return CAMARMAIN, buf.String()
	case "CAMARFUNC":
		return CAMARFUNC, buf.String()
	case "CAMARFOR":
		return CAMARFOR, buf.String()
	case "CAMARELSEIF":
		return CAMARELSEIF, buf.String()
	case "VAR":
		return VAR, buf.String()
	case "CAMARBREAK":
		return CAMARBREAK, buf.String()
	case "CAMARWHILE":
		return CAMARWHILE, buf.String()
	case "CAMARDO":
		return CAMARDO, buf.String()
	case "CAMARINT":
		return CAMARINT, buf.String()
	case "CAMARTRY":
		return CAMARTRY, buf.String()
	case "CAMARCATCH":
		return CAMARCATCH, buf.String()
	case "CAMARFINALLY":
		return CAMARFINALLY, buf.String()
	case "CAMARPOINT":
		return CAMARPOINT, buf.String()
	case "CAMARNEW":
		return CAMARNEW, buf.String()
	case "CAMARFIXED":
		return CAMARFIXED, buf.String()
	case "CAMARUNCHECKED":
		return CAMARUNCHECKED, buf.String()
	case "CAMARFOREACH":
		return CAMARFOREACH, buf.String()
	case "CAMARIMPORT":
		return CAMARIMPORT, buf.String()
	case "CAMARRETURN":
		return CAMARRETURN, buf.String()

	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' }

// isLetter returns true if the rune is a letter.
func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

// isDigit returns true if the rune is a digit.
func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

// eof represents a marker rune for the end of the reader.
var eof = rune(0)
