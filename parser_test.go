package camargo_test

import (
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestParser_ParseStatement(t *testing.T) {

	//Erasmo
	f, _ := ioutil.ReadFile("C://MARCO//marco.txt")
	z := string(f[:])
	println()
	println("Su funcion es: ", z)
	println()

	var tests = []struct {
		s    string
		stmt *camargo.SelectStatement
		err  string
	}{
		{
			s:    z,
			stmt: &camargo.SelectStatement{},
		},
	}
	for i, tt := range tests {
		stmt, err := camargo.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}
}
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
