package marco

import (
	"fmt"
	"io"
)

// SelectStatement represents a SQL SELECT statement.
type SelectStatement struct{}

//type FromStatement struct{}
//type Statement struct{}
type Goifstatement struct{}

//type Goswitchstatement struct {}
//type GoParseSuma struct {}
//type GoParseSumaNom struct {}
//type GoParsefor struct {}
//type GoParsegoelseif struct {}

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

// Parse parses a SQL SELECT statement.
//func (p *Parser) ParseGoif() (*Goifstatement, error) {
//	stmt := &Goifstatement{}
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goif {
//		return nil, fmt.Errorf("encontro %q, esperaba Goif", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
//		return nil, fmt.Errorf("encontro %q, esperaba ParDer", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
//		return nil, fmt.Errorf("encontro %q, esperaba ParIsq", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goelse {
//		return nil, fmt.Errorf("encontro %q, esperaba Goelse", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, Dospunt :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
//		return nil, fmt.Errorf("encontro %q, LlaveIsq :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
//		return nil, fmt.Errorf("encontro %q, esperaba Llaveder", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goend {
//		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	return stmt, nil
//}
//func (p *Parser) ParseGoswitch() (*Goswitchstatement, error) {
//	stmt := &Goswitchstatement{}
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goswitch {
//		return nil, fmt.Errorf("encontro %q, esperaba Goswitch", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
//		return nil, fmt.Errorf("encontro %q, esperaba ParDer", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
//		return nil, fmt.Errorf("encontro %q, esperaba ParIsq", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goend {
//		return nil, fmt.Errorf("encontro %q, esperaba Goend", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
//		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
//	}
//	return stmt, nil
//}
func (p *Parser) Parse() (*SelectStatement, error) {
	stmt := &SelectStatement{}

	//Bonifaz

	//if tok, lit := p.scanIgnoreWhitespace(); tok != MCCREATE {
	//	return nil, fmt.Errorf("encontro %q, esperaba MCCREATE", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != DATABASE {
	//	return nil, fmt.Errorf("encontro %q, esperaba DATABASE", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
	//	return nil, fmt.Errorf("encontro %q, esperaba Nombre de la tabla", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != MCUSE {
	//	return nil, fmt.Errorf("encontro %q, esperaba Goif", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
	//	return nil, fmt.Errorf("encontro %q, esperaba Nombre de la tabla", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != MCCREATETABLE {
	//	return nil, fmt.Errorf("encontro %q, esperaba MCCREATETABLE", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
	//	return nil, fmt.Errorf("encontro %q, esperaba Nombre de la tabla", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
	//	return nil, fmt.Errorf("encontro %q, esperaba gofor", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
	//	return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	//}
	//if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
	//	return nil, fmt.Errorf("encontro %q, esperaba ident", lit)
	//}


	//Erasmo

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOIF {
		return nil, fmt.Errorf("encontro %q, esperaba Goif", lit)
	}
	return stmt,nil
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

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOELSE {
		return nil, fmt.Errorf("encontro %q, esperaba GOELSE", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba Goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//Suma global.

	if tok, lit := p.scanIgnoreWhitespace(); tok != VAR {
		return nil, fmt.Errorf("encontro %q, esperaba VAR", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba variable", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba tipo de dato", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba gofunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba gomain", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//suma normal

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba gofunc", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba gomain", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	// CICLO GOFOR

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFOR {
		return nil, fmt.Errorf("encontro %q, esperaba gofor", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
		return nil, fmt.Errorf("encontro %q, esperaba gofor", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//GOELSEIF

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOELSEIF {
		return nil, fmt.Errorf("encontro %q, esperaba goelseif", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//goswitch

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOSWITCH {
		return nil, fmt.Errorf("encontro %q, esperaba goswitch", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOBREAK {
		return nil, fmt.Errorf("encontro %q, esperaba gobreak", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//gowhile

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba GOINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba gowhile", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//godo ciclo
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba goint", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba identificador", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GODO {
		return nil, fmt.Errorf("encontro %q, esperaba GODO", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOWHILE {
		return nil, fmt.Errorf("encontro %q, esperaba gowhile", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//GOTRY
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOTRY {
		return nil, fmt.Errorf("encontro %q, esperaba gotry", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOCATCH {
		return nil, fmt.Errorf("encontro %q, esperaba Gocatch", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//gofinally

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFINALLY {
		return nil, fmt.Errorf("encontro %q, esperaba gofinally", lit)
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
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//gopoint

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOPOINT {
		return nil, fmt.Errorf("encontro %q, esperaba GOPOINT", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GONEW {
		return nil, fmt.Errorf("encontro %q, esperaba GONEW", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//GOFIXED

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFIXED {
		return nil, fmt.Errorf("encontro %q, esperaba GOFIXED", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//GOUNCHECKED

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba GOUNCHECKED", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba goint", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba valores", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != Asignaicon {
		return nil, fmt.Errorf("encontro %q, esperaba =", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOUNCHECKED {
		return nil, fmt.Errorf("encontro %q, esperaba GOUNCHECKED", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	////ARREGLO
	//
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba GOINT", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
		return nil, fmt.Errorf("encontro %q, esperaba instrucciones", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GONEW {
		return nil, fmt.Errorf("encontro %q, esperaba gonew", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOINT {
		return nil, fmt.Errorf("encontro %q, esperaba goint", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//
	////goforeach

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFOREACH {
		return nil, fmt.Errorf("encontro %q, esperaba GOFOREACH", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}
	//PAQUETES
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOIMPORT {
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//funciones

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOFUNC {
		return nil, fmt.Errorf("encontro %q, esperaba gofunc", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	//gomain

	if tok, lit := p.scanIgnoreWhitespace(); tok != GOMAIN {
		return nil, fmt.Errorf("encontro %q, esperaba gomain", lit)
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
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOEND {
		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
	}
	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
	}

	return stmt, nil

}
func (p *Parser) ParseGoIF() (*Goifstatement, error) {
	stmt := &Goifstatement{}
	// First token should be a "SELECT" keyword.
	if tok, lit := p.scanIgnoreWhitespace(); tok != GOIF {
		return nil, fmt.Errorf("encontro %q, esperaba Goif", lit)
	}
	return stmt, nil
}

//func (p *Parser) ParseSum() (*GoParseSuma, error) {
//	stmt := &GoParseSuma{}
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goint {
//		return nil, fmt.Errorf("encontro %q, esperaba Goint", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
//		return nil, fmt.Errorf("encontro %q, esperaba Indent", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != COMMA {
//		return nil, fmt.Errorf("encontro %q, esperaba ,", lit)
//	}
//
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goint {
//		return nil, fmt.Errorf("encontro %q, esperaba Goint", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != IDENT {
//		return nil, fmt.Errorf("encontro %q, esperaba Indent", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Gofunc {
//		return nil, fmt.Errorf("encontro %q, esperaba gofunc", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Gomain {
//		return nil, fmt.Errorf("encontro %q, esperaba Gomain", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
//		return nil, fmt.Errorf("encontro %q, esperaba Parder", lit)
//	}
//	return stmt, nil
//}
//
//func (p *Parser) ParseFromStatement() (*FromStatement, error) {
//	stmt := &FromStatement{}
//
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != FROM {
//		return nil, fmt.Errorf("encontro %q, esperaba SELECT", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != SELECT {
//		return nil, fmt.Errorf("encontro %q, esperaba FROM", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != SELECT {
//		return nil, fmt.Errorf("encontro %q, esperaba FROM", lit)
//	}
//
//	return stmt, nil
//}
//
//func (p *Parser) ParseSumNom() (*GoParseSumaNom, error) {
//	stmt := &GoParseSumaNom{}
//
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Gofunc {
//		return nil, fmt.Errorf("encontro %q, esperaba gofunc", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Gomain {
//		return nil, fmt.Errorf("encontro %q, esperaba gomain", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goend {
//		return nil, fmt.Errorf("encontro %q, esperaba goend", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
//		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
//	}
//
//	return stmt, nil
//}
//func (p *Parser) ParseGoFor() (*GoParsefor, error) {
//	stmt := &GoParsefor{}
//
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Gofor {
//		return nil, fmt.Errorf("encontro %q, esperaba gofor", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParDer {
//		return nil, fmt.Errorf("encontro %q, esperaba (", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != ParIsq {
//		return nil, fmt.Errorf("encontro %q, esperaba )", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goend {
//		return nil, fmt.Errorf("encontro %q, esperaba Goend", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
//		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
//	}
//
//	return stmt, nil
//}
//func (p *Parser) ParseGoelseif() (*GoParsegoelseif, error) {
//	stmt := &GoParsegoelseif{}
//
//	// First token should be a "SELECT" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goelseif {
//		return nil, fmt.Errorf("encontro %q, esperaba goelseif", lit)
//	}
//	// Next we should see the "FROM" keyword.
//	if tok, lit := p.scanIgnoreWhitespace(); tok != DosPunt {
//		return nil, fmt.Errorf("encontro %q, esperaba :", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveDer {
//		return nil, fmt.Errorf("encontro %q, esperaba {", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != LlaveIsq {
//		return nil, fmt.Errorf("encontro %q, esperaba }", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != Goend {
//		return nil, fmt.Errorf("encontro %q, esperaba Goend", lit)
//	}
//	if tok, lit := p.scanIgnoreWhitespace(); tok != PuntCom {
//		return nil, fmt.Errorf("encontro %q, esperaba ;", lit)
//	}
//
//	return stmt, nil
//}
//
//
// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
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

/*
// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Message  string
	Found    string
	Expected []string
}


// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// newParseError returns a new instance of ParseError.
func newParseError(found string, expected []string) *ParseError {
	return &ParseError{Found: found, Expected: expected}
}

// Error returns the string representation of the error.
func (e *ParseError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s at line %d, char %d", e.Message)
	}
	return fmt.Sprintf("found %s, expected %s at line %d, char %d", e.Found)
}
*/
