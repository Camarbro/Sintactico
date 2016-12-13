package camargo

import (
	"fmt"
	"io"
)

type SelectStatement struct{}

type Goifstatement struct{}

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Parse() (*SelectStatement, error) {

	stmt := &SelectStatement{}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARIF {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba operacion", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)

	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)

	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARELSE {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARELSE", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)

	}

	//DO-WHILE EJEMPLO EXTRA SI NO LO BORRO.

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARint", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba asignacion", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digito", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARDO {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARdo", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}

	//TEST
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instruccion", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != COMMA {
		return nil, fmt.Errorf("encontro %q, esperaba ,", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//test1

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba ( ", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARFOR

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digito", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFOR {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba digito", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba digito", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digito", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARSWITCH

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARSWITCH {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARswitch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	return stmt, nil
	//CAMARfuncprint hola

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARfuncmain para numero

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba iden", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != NUMEROS {
		return nil, fmt.Errorf("encontro %q, esperaba digit", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//AQUI TERMINA
	//Suma global.

	if tok, lit := p.scanIgnoreWhitespace(); tok != VAR {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != VAR {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//suma normal

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	// CICLO CAMARFOR

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFOR {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARELSEIF

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARELSEIF {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARelseif", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARswitch

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARSWITCH {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARswitch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARBREAK {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARbreak", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARwhile

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARdo ciclo
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARDO {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARDO", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARwhile", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARTRY
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARTRY {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARtry", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARCATCH {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARcatch", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARfinally

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFINALLY {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfinally", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARpoint

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARPOINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARPOINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARNEW {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARNEW", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARFIXED

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFIXED {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARFIXED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARUNCHECKED

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != OPREL {
		return nil, fmt.Errorf("encontro %q, esperaba oprel", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARUNCHECKED", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	////ARREGLO
	//
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	//ERROR

	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}

	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARNEW {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARnew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARINT {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorDER {
		return nil, fmt.Errorf("encontro %q, esperaba [", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != AgruppadorISQ {
		return nil, fmt.Errorf("encontro %q, esperaba ]", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//
	////CAMARforeach

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFOREACH {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARFOREACH", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//PAQUETES
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARIMPORT {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//funciones

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARfunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//CAMARmain

	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARmain", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMAREND {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	return stmt, nil

}
func (p *Parser) ParseGoIF() (*Goifstatement, error) {
	stmt := &Goifstatement{}
	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != CAMARIF {
		return nil, fmt.Errorf("encontro %q, esperaba CAMARif", lit)
	}
	return stmt, nil
}

func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}
