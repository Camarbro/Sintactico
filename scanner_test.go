package camargo_test

import (
	"strings"
	"testing"

	"github.com/Camarbro/Sintactico"
)

// Ensure the scanner can scan tokens correctly.
func TestScanner_Scan(t *testing.T) {
	var tests = []struct {
		s   string
		tok camargo.Token
		lit string
	}{
		// Special tokens (EOF, ILLEGAL, WS)
		{s: ``, tok: camargo.EOF},
		{s: `#`, tok: camargo.ILLEGAL, lit: `#`},
		{s: ` `, tok: camargo.WS, lit: " "},
		{s: "\t", tok: camargo.WS, lit: "\t"},
		//{s: "\n", tok: camargo.WS, lit: "\n"},
		{s: "\r", tok: camargo.WS, lit: "\r"},

		// Misc characters
		//{s: `*`, tok: camargo.ASTERISK, lit: "*"},
		{s: `:`, tok: camargo.DosPunt, lit: ":"},
		{s: `(`, tok: camargo.ParDer, lit: "("},
		{s: `)`, tok: camargo.ParIsq, lit: ")"},
		{s: `;`, tok: camargo.PuntCom, lit: ";"},

		// Identifiers
		{s: `foo`, tok: camargo.IDENT, lit: `foo`},
		{s: `Zx12_3U_-`, tok: camargo.IDENT, lit: `Zx12_3U_`},
		//{s: `Zx12_3U_-`, tok: camargo.IDENT, lit: `Zx12_3U_`},

		//Keywords
		{s: `CAMARIF`, tok: camargo.CAMARIF, lit: "CAMARIF"},
		{s: `CAMARELSEIF`, tok: camargo.CAMARELSEIF, lit: "CAMARELSEIF"},
		{s: `CAMARMAIN`, tok: camargo.CAMARMAIN, lit: "CAMARMAIN"},
		{s: `CAMARELSE`, tok: camargo.CAMARELSE, lit: "CAMARELSE"},
		{s: `CAMARSWITCH`, tok: camargo.CAMARSWITCH, lit: "CAMARSWITCH"},
		{s: `CAMARFUNC`, tok: camargo.CAMARFUNC, lit: "CAMARFUNC"},
		{s: `CAMARINT`, tok: camargo.CAMARINT, lit: "CAMARINT"},
		{s: `CAMAREND`, tok: camargo.CAMAREND, lit: "CAMAREND"},
		{s: `VAR`, tok: camargo.VAR, lit: "VAR"},
		{s: `CAMARWHILE`, tok: camargo.CAMARWHILE, lit: "CAMARWHILE"},
		{s: `CAMARDO`, tok: camargo.CAMARDO, lit: "CAMARDO"},
		{s: `CAMARTRY`, tok: camargo.CAMARTRY, lit: "CAMARTRY"},
		{s: `CAMARCATCH`, tok: camargo.CAMARCATCH, lit: "CAMARCATCH"},
		{s: `CAMARFINALLY`, tok: camargo.CAMARFINALLY, lit: "CAMARFINALLY"},
		{s: `CAMARPOINT`, tok: camargo.CAMARPOINT, lit: "CAMARPOINT"},
		{s: `CAMARFIXED`, tok: camargo.CAMARFIXED, lit: "CAMARFIXED"},
		{s: `CAMARUNCHECKED`, tok: camargo.CAMARUNCHECKED, lit: "CAMARUNCHECKED"},
		{s: `CAMARNEW`, tok: camargo.CAMARNEW, lit: "CAMARNEW"},
		{s: `CAMARFOREACH`, tok: camargo.CAMARFOREACH, lit: "CAMARFOREACH"},
		{s: `CAMARIMPORT`, tok: camargo.CAMARIMPORT, lit: "CAMARIMPORT"},
		{s: `CAMARRETURN`, tok: camargo.CAMARRETURN, lit: "CAMARRETURN"},
	}

	for i, tt := range tests {
		s := camargo.NewScanner(strings.NewReader(tt.s))
		tok, lit := s.Scan()
		if tt.tok != tok {
			t.Errorf("%d. %q token mismatch: exp=%q got=%q <%q>", i, tt.s, tt.tok, tok, lit)
		} else if tt.lit != lit {
			t.Errorf("%d. %q literal mismatch: exp=%q got=%q", i, tt.s, tt.lit, lit)
		}
	}
}
