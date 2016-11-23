package marco_test

import (
	"github.com/elephannt/parsercompo"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

// Ensure the parser can parse strings into Statement ASTs.

func TestParser_ParseStatement(t *testing.T) {

	//Erasmo
	f, _ := ioutil.ReadFile("C://Download//Pokemon//grammatic3.txt")
	z := string(f[:])
	//Bonifaz
	//f, _ := ioutil.ReadFile("C://Download//Pokemon//bonifaz.txt")
	//z := string(f[:])
	//z:=strings.Fields(r)
	//var tests1 = []struct {
	//	j    string
	//	stmt *marco.Goifstatement
	//	err  string
	//}{}
	var tests = []struct {
		s    string
		stmt *marco.SelectStatement
		err  string
	}{
		{
			s:    z,
			stmt: &marco.SelectStatement{
			//Fields: []string{""},
			//TableName: "marco",
			},
		},
	}
	for i, tt := range tests {
		stmt, err := marco.NewParser(strings.NewReader(tt.s)).Parse()
		if !reflect.DeepEqual(tt.err, errstring(err)) {
			t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
		} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
			t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
		}
	}

	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseGoif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseGoif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseGoswitch()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q:\ error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseSum()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseSumNom()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseGoFor()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//for i, tt := range tests {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.s)).ParseGoelseif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
	//

	//for i, tt := range tests1 {
	//	stmt, err := marco.NewParser(strings.NewReader(tt.j)).ParseGoif()
	//	if !reflect.DeepEqual(tt.err, errstring(err)) {
	//		t.Errorf("%d. %q: error mismatch:\n  exp=%s\n  got=%s\n\n", i, tt.s, tt.err, err)
	//	} else if tt.err == "" && !reflect.DeepEqual(tt.stmt, stmt) {
	//		t.Errorf("%d. %q\n\nstmt mismatch:\n\nexp=%#v\n\ngot=%#v\n\n", i, tt.s, tt.stmt, stmt)
	//	}
	//}
}

// errstring returns the string representation of an error.
func errstring(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
