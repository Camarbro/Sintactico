package marco_test

import (
	"github.com/elephannt/parsercompo"
	"strings"
	"testing"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok marco.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: marco.EOF},
		{s: `#`, tok: marco.ILLEGAL, lit: `#`},
		{s: ` `, tok: marco.WS, lit: " "},
		{s: "\t", tok: marco.WS, lit: "\t"},
		//{s: "\n", tok: marco.WS, lit: "\n"},
		{s: "\r", tok: marco.WS, lit: "\r"},

		// Misc characters
		{s: `*`, tok: marco.ASTERISK, lit: "*"},
		{s: `:`, tok: marco.DosPunt, lit: ":"},
		{s: `(`, tok: marco.ParDer, lit: "("},
		{s: `)`, tok: marco.ParIsq, lit: ")"},
		{s: `;`, tok: marco.PuntCom, lit: ";"},
		{s: `=`, tok: marco.Asignaicon, lit: "="},
		//{s: `[`, tok: marco.AgruppadorDER, lit: "]"},
		//{s: `[`, tok: marco.AgruppadorISQ, lit: "]"},
		//{s: `\n`, tok: marco.SaltoDeL, lit: "\n"},

		// Identifiers
		{s: `foo`, tok: marco.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: marco.IDENT, lit: `Zx12_3U_`},

		//Keywords
		{s: `GOIF`, tok: marco.GOIF, lit: "GOIF"},
		{s: `GOELSEIF`, tok: marco.GOELSEIF, lit: "GOELSEIF"},
		{s: `GOMAIN`, tok: marco.GOMAIN, lit: "GOMAIN"},
		{s: `GOELSE`, tok: marco.GOELSE, lit: "GOELSE"},
		{s: `GOSWITCH`, tok: marco.GOSWITCH, lit: "GOSWITCH"},
		{s: `GOFUNC`, tok: marco.GOFUNC, lit: "GOFUNC"},
		{s: `GOINT`, tok: marco.GOINT, lit: "GOINT"},
		{s: `GOEND`, tok: marco.GOEND, lit: "GOEND"},
		{s: `VAR`, tok: marco.VAR, lit: "VAR"},
		{s: `GOWHILE`, tok: marco.GOWHILE, lit: "GOWHILE"},
		{s: `GODO`, tok: marco.GODO, lit: "GODO"},
		{s: `GOTRY`, tok: marco.GOTRY, lit: "GOTRY"},
		{s: `GOCATCH`, tok: marco.GOCATCH, lit: "GOCATCH"},
		{s: `GOFINALLY`, tok: marco.GOFINALLY, lit: "GOFINALLY"},
		{s: `GOPOINT`, tok: marco.GOPOINT, lit: "GOPOINT"},
		{s: `GOFIXED`, tok: marco.GOFIXED, lit: "GOFIXED"},
		{s: `GOUNCHECKED`, tok: marco.GOUNCHECKED, lit: "GOUNCHECKED"},
		{s: `GONEW`, tok: marco.GONEW, lit: "GONEW"},
		{s: `GOFOREACH`, tok: marco.GOFOREACH, lit: "GOFOREACH"},
		{s: `GOIMPORT`, tok: marco.GOIMPORT, lit: "GOIMPORT"},


		//Bonifaz
		//{s: `MCCREATE`, tok: marco.MCCREATE, lit: "MCCREATE"},
		//
		//{s: `MCUSE`, tok: marco.MCUSE, lit: "MCUSE"},
		//
		//{s: `MCCREATETABLE`, tok: marco.MCCREATETABLE, lit: "MCCREATETABLE"},
		//
		//{s: `DATABASE`, tok: marco.DATABASE, lit: "DATABASE"},




	}

	for i, tt := range tests {
		s := marco.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
