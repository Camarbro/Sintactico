package camargo

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF
	WS
	// Literals
	IDENT
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
	OPREL         //== ! + - / * |

	// Keywords
	SELECT
	FROM
	CAMARIF
	CAMARSWITCH
	CAMAREND
	CAMARINT
	CAMARFUNC
	CAMARMAIN
	CAMARELSE
	CAMARFOR
	CAMARELSEIF
	VAR
	CAMARBREAK
	CAMARWHILE
	CAMARDO
	CAMARTRY
	CAMARCATCH
	CAMARFINALLY
	CAMARPOINT
	CAMARNEW
	CAMARFIXED
	CAMARFOREACH
	CAMARUNCHECKED
	CAMARIMPORT
	CAMARRETURN
	NUMEROS
)
