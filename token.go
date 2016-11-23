package marco

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT // main
	// Misc characters
	ASTERISK      // *
	COMMA         // ,
	DosPunt       // :
	ParDer        // (
	ParIsq        // )
	PuntCom       // ;
	LlaveDer      // {
	LlaveIsq      // }
	SaltoDeL      //n
	Asignaicon    // =
	AgruppadorDER // [
	AgruppadorISQ // ]

	// Keywords
	SELECT
	FROM
	GOIF
	GOSWITCH
	GOEND
	GOINT // Tipodedato
	GOFUNC
	GOMAIN
	GOELSE
	GOFOR
	GOELSEIF
	VAR // variables
	GOBREAK
	GOWHILE
	GODO
	GOTRY
	GOCATCH
	GOFINALLY
	GOPOINT
	GONEW
	GOFIXED
	GOFOREACH
	GOUNCHECKED
	GOIMPORT

	//Bonifaz

	MCCREATE
	DATABASE
	MCUSE
	MCCREATETABLE

)
