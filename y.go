//line grammar.y:5
package stop

import __yyfmt__ "fmt"

//line grammar.y:7
import (
	"fmt"

	"github.com/patdhlk/stop/ast"
)

//line grammar.y:16
type parserSymType struct {
	yys      int
	node     ast.Node
	nodeList []ast.Node
	str      string
	token    *parserToken
}

const PROGRAM_BRACKET_LEFT = 57346
const PROGRAM_BRACKET_RIGHT = 57347
const PROGRAM_STRING_START = 57348
const PROGRAM_STRING_END = 57349
const PAREN_LEFT = 57350
const PAREN_RIGHT = 57351
const COMMA = 57352
const SQUARE_BRACKET_LEFT = 57353
const SQUARE_BRACKET_RIGHT = 57354
const ARITH_OP = 57355
const IDENTIFIER = 57356
const INTEGER = 57357
const FLOAT = 57358
const BOOL = 57359
const STRING = 57360

var parserToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"PROGRAM_BRACKET_LEFT",
	"PROGRAM_BRACKET_RIGHT",
	"PROGRAM_STRING_START",
	"PROGRAM_STRING_END",
	"PAREN_LEFT",
	"PAREN_RIGHT",
	"COMMA",
	"SQUARE_BRACKET_LEFT",
	"SQUARE_BRACKET_RIGHT",
	"ARITH_OP",
	"IDENTIFIER",
	"INTEGER",
	"FLOAT",
	"BOOL",
	"STRING",
}
var parserStatenames = [...]string{}

const parserEofCode = 1
const parserErrCode = 2
const parserInitialStackSize = 16

//line grammar.y:208

//line yacctab:1
var parserExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const parserNprod = 22
const parserPrivate = 57344

var parserTokenNames []string
var parserStates []string

const parserLast = 39

var parserAct = [...]int{

	9, 7, 30, 18, 24, 17, 18, 21, 18, 3,
	22, 19, 8, 18, 1, 6, 20, 28, 29, 23,
	25, 8, 26, 27, 7, 11, 2, 4, 10, 5,
	31, 0, 0, 15, 16, 12, 13, 14, 6,
}
var parserPact = [...]int{

	-3, -1000, -3, -1000, -1000, -1000, -1000, 20, -1000, 0,
	20, -3, -1000, -1000, -1000, 20, -1, -1000, 20, -5,
	-1000, 20, 20, -1000, -1000, 8, -7, -10, -1000, 20,
	-1000, -7,
}
var parserPgo = [...]int{

	0, 0, 29, 27, 25, 9, 20, 14,
}
var parserR1 = [...]int{

	0, 7, 7, 4, 4, 5, 5, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 6, 6,
	6, 3,
}
var parserR2 = [...]int{

	0, 0, 1, 1, 2, 1, 1, 3, 3, 1,
	1, 1, 1, 2, 3, 1, 4, 4, 0, 3,
	1, 1,
}
var parserChk = [...]int{

	-1000, -7, -4, -5, -3, -2, 18, 4, -5, -1,
	8, -4, 15, 16, 17, 13, 14, 5, 13, -1,
	-1, 8, 11, -1, 9, -6, -1, -1, 9, 10,
	12, -1,
}
var parserDef = [...]int{

	1, -2, 2, 3, 5, 6, 21, 0, 4, 0,
	0, 9, 10, 11, 12, 0, 15, 7, 0, 0,
	13, 18, 0, 14, 8, 0, 20, 0, 16, 0,
	17, 19,
}
var parserTok1 = [...]int{

	1,
}
var parserTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18,
}
var parserTok3 = [...]int{
	0,
}

var parserErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	parserDebug        = 0
	parserErrorVerbose = false
)

type parserLexer interface {
	Lex(lval *parserSymType) int
	Error(s string)
}

type parserParser interface {
	Parse(parserLexer) int
	Lookahead() int
}

type parserParserImpl struct {
	lval  parserSymType
	stack [parserInitialStackSize]parserSymType
	char  int
}

func (p *parserParserImpl) Lookahead() int {
	return p.char
}

func parserNewParser() parserParser {
	return &parserParserImpl{}
}

const parserFlag = -1000

func parserTokname(c int) string {
	if c >= 1 && c-1 < len(parserToknames) {
		if parserToknames[c-1] != "" {
			return parserToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func parserStatname(s int) string {
	if s >= 0 && s < len(parserStatenames) {
		if parserStatenames[s] != "" {
			return parserStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func parserErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !parserErrorVerbose {
		return "syntax error"
	}

	for _, e := range parserErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + parserTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := parserPact[state]
	for tok := TOKSTART; tok-1 < len(parserToknames); tok++ {
		if n := base + tok; n >= 0 && n < parserLast && parserChk[parserAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if parserDef[state] == -2 {
		i := 0
		for parserExca[i] != -1 || parserExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; parserExca[i] >= 0; i += 2 {
			tok := parserExca[i]
			if tok < TOKSTART || parserExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if parserExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += parserTokname(tok)
	}
	return res
}

func parserlex1(lex parserLexer, lval *parserSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = parserTok1[0]
		goto out
	}
	if char < len(parserTok1) {
		token = parserTok1[char]
		goto out
	}
	if char >= parserPrivate {
		if char < parserPrivate+len(parserTok2) {
			token = parserTok2[char-parserPrivate]
			goto out
		}
	}
	for i := 0; i < len(parserTok3); i += 2 {
		token = parserTok3[i+0]
		if token == char {
			token = parserTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = parserTok2[1] /* unknown char */
	}
	if parserDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", parserTokname(token), uint(char))
	}
	return char, token
}

func parserParse(parserlex parserLexer) int {
	return parserNewParser().Parse(parserlex)
}

func (parserrcvr *parserParserImpl) Parse(parserlex parserLexer) int {
	var parsern int
	var parserVAL parserSymType
	var parserDollar []parserSymType
	_ = parserDollar // silence set and not used
	parserS := parserrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	parserstate := 0
	parserrcvr.char = -1
	parsertoken := -1 // parserrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		parserstate = -1
		parserrcvr.char = -1
		parsertoken = -1
	}()
	parserp := -1
	goto parserstack

ret0:
	return 0

ret1:
	return 1

parserstack:
	/* put a state and value onto the stack */
	if parserDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", parserTokname(parsertoken), parserStatname(parserstate))
	}

	parserp++
	if parserp >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserS[parserp] = parserVAL
	parserS[parserp].yys = parserstate

parsernewstate:
	parsern = parserPact[parserstate]
	if parsern <= parserFlag {
		goto parserdefault /* simple state */
	}
	if parserrcvr.char < 0 {
		parserrcvr.char, parsertoken = parserlex1(parserlex, &parserrcvr.lval)
	}
	parsern += parsertoken
	if parsern < 0 || parsern >= parserLast {
		goto parserdefault
	}
	parsern = parserAct[parsern]
	if parserChk[parsern] == parsertoken { /* valid shift */
		parserrcvr.char = -1
		parsertoken = -1
		parserVAL = parserrcvr.lval
		parserstate = parsern
		if Errflag > 0 {
			Errflag--
		}
		goto parserstack
	}

parserdefault:
	/* default state action */
	parsern = parserDef[parserstate]
	if parsern == -2 {
		if parserrcvr.char < 0 {
			parserrcvr.char, parsertoken = parserlex1(parserlex, &parserrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if parserExca[xi+0] == -1 && parserExca[xi+1] == parserstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			parsern = parserExca[xi+0]
			if parsern < 0 || parsern == parsertoken {
				break
			}
		}
		parsern = parserExca[xi+1]
		if parsern < 0 {
			goto ret0
		}
	}
	if parsern == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			parserlex.Error(parserErrorMessage(parserstate, parsertoken))
			Nerrs++
			if parserDebug >= 1 {
				__yyfmt__.Printf("%s", parserStatname(parserstate))
				__yyfmt__.Printf(" saw %s\n", parserTokname(parsertoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for parserp >= 0 {
				parsern = parserPact[parserS[parserp].yys] + parserErrCode
				if parsern >= 0 && parsern < parserLast {
					parserstate = parserAct[parsern] /* simulate a shift of "error" */
					if parserChk[parserstate] == parserErrCode {
						goto parserstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if parserDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", parserS[parserp].yys)
				}
				parserp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if parserDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", parserTokname(parsertoken))
			}
			if parsertoken == parserEofCode {
				goto ret1
			}
			parserrcvr.char = -1
			parsertoken = -1
			goto parsernewstate /* try again in the same state */
		}
	}

	/* reduction by production parsern */
	if parserDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", parsern, parserStatname(parserstate))
	}

	parsernt := parsern
	parserpt := parserp
	_ = parserpt // guard against "declared and not used"

	parserp -= parserR2[parsern]
	// parserp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if parserp+1 >= len(parserS) {
		nyys := make([]parserSymType, len(parserS)*2)
		copy(nyys, parserS)
		parserS = nyys
	}
	parserVAL = parserS[parserp+1]

	/* consult goto table to find next state */
	parsern = parserR1[parsern]
	parserg := parserPgo[parsern]
	parserj := parserg + parserS[parserp].yys + 1

	if parserj >= parserLast {
		parserstate = parserAct[parserg]
	} else {
		parserstate = parserAct[parserj]
		if parserChk[parserstate] != -parsern {
			parserstate = parserAct[parserg]
		}
	}
	// dummy call; replaced with literal code
	switch parsernt {

	case 1:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line grammar.y:38
		{
			parserResult = &ast.LiteralNode{
				Value: "",
				Typex: ast.TString,
				Posx:  ast.Pos{Column: 1, Line: 1},
			}
		}
	case 2:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:46
		{
			parserResult = parserDollar[1].node

			// We want to make sure that the top value is always an Output
			// so that the return value is always a string, list of map from an
			// interpolation.
			//
			// The logic for checking for a LiteralNode is a little annoying
			// because functionally the AST is the same, but we do that because
			// it makes for an easy literal check later (to check if a string
			// has any interpolations).
			if _, ok := parserDollar[1].node.(*ast.Output); !ok {
				if n, ok := parserDollar[1].node.(*ast.LiteralNode); !ok || n.Typex != ast.TString {
					parserResult = &ast.Output{
						Exprs: []ast.Node{parserDollar[1].node},
						Posx:  parserDollar[1].node.Pos(),
					}
				}
			}
		}
	case 3:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:69
		{
			parserVAL.node = parserDollar[1].node
		}
	case 4:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line grammar.y:73
		{
			var result []ast.Node
			if c, ok := parserDollar[1].node.(*ast.Output); ok {
				result = append(c.Exprs, parserDollar[2].node)
			} else {
				result = []ast.Node{parserDollar[1].node, parserDollar[2].node}
			}

			parserVAL.node = &ast.Output{
				Exprs: result,
				Posx:  result[0].Pos(),
			}
		}
	case 5:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:89
		{
			parserVAL.node = parserDollar[1].node
		}
	case 6:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:93
		{
			parserVAL.node = parserDollar[1].node
		}
	case 7:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line grammar.y:99
		{
			parserVAL.node = parserDollar[2].node
		}
	case 8:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line grammar.y:105
		{
			parserVAL.node = parserDollar[2].node
		}
	case 9:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:109
		{
			parserVAL.node = parserDollar[1].node
		}
	case 10:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:113
		{
			parserVAL.node = &ast.LiteralNode{
				Value: parserDollar[1].token.Value.(int),
				Typex: ast.TInt,
				Posx:  parserDollar[1].token.Pos,
			}
		}
	case 11:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:121
		{
			parserVAL.node = &ast.LiteralNode{
				Value: parserDollar[1].token.Value.(float64),
				Typex: ast.TFloat,
				Posx:  parserDollar[1].token.Pos,
			}
		}
	case 12:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:129
		{
			parserVAL.node = &ast.LiteralNode{
				Value: parserDollar[1].token.Value.(bool),
				Typex: ast.TBool,
				Posx:  parserDollar[1].token.Pos,
			}
		}
	case 13:
		parserDollar = parserS[parserpt-2 : parserpt+1]
		//line grammar.y:137
		{
			// This is REALLY jank. We assume that a singular ARITH_OP
			// means 0 ARITH_OP expr, which... is weird. We don't want to
			// support *, /, etc., only -. We should fix this later with a pure
			// Go scanner/parser.
			if parserDollar[1].token.Value.(ast.ArithmeticOp) != ast.ArithmeticOpSub {
				if parserErr == nil {
					parserErr = fmt.Errorf("Invalid unary operation: %v", parserDollar[1].token.Value)
				}
			}

			parserVAL.node = &ast.Arithmetic{
				Op: parserDollar[1].token.Value.(ast.ArithmeticOp),
				Exprs: []ast.Node{
					&ast.LiteralNode{Value: 0, Typex: ast.TInt},
					parserDollar[2].node,
				},
				Posx: parserDollar[2].node.Pos(),
			}
		}
	case 14:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line grammar.y:158
		{
			parserVAL.node = &ast.Arithmetic{
				Op:    parserDollar[2].token.Value.(ast.ArithmeticOp),
				Exprs: []ast.Node{parserDollar[1].node, parserDollar[3].node},
				Posx:  parserDollar[1].node.Pos(),
			}
		}
	case 15:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:166
		{
			parserVAL.node = &ast.VariableAccess{Name: parserDollar[1].token.Value.(string), Posx: parserDollar[1].token.Pos}
		}
	case 16:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line grammar.y:170
		{
			parserVAL.node = &ast.Call{Func: parserDollar[1].token.Value.(string), Args: parserDollar[3].nodeList, Posx: parserDollar[1].token.Pos}
		}
	case 17:
		parserDollar = parserS[parserpt-4 : parserpt+1]
		//line grammar.y:174
		{
			parserVAL.node = &ast.Index{
				Target: &ast.VariableAccess{
					Name: parserDollar[1].token.Value.(string),
					Posx: parserDollar[1].token.Pos,
				},
				Key:  parserDollar[3].node,
				Posx: parserDollar[1].token.Pos,
			}
		}
	case 18:
		parserDollar = parserS[parserpt-0 : parserpt+1]
		//line grammar.y:186
		{
			parserVAL.nodeList = nil
		}
	case 19:
		parserDollar = parserS[parserpt-3 : parserpt+1]
		//line grammar.y:190
		{
			parserVAL.nodeList = append(parserDollar[1].nodeList, parserDollar[3].node)
		}
	case 20:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:194
		{
			parserVAL.nodeList = append(parserVAL.nodeList, parserDollar[1].node)
		}
	case 21:
		parserDollar = parserS[parserpt-1 : parserpt+1]
		//line grammar.y:200
		{
			parserVAL.node = &ast.LiteralNode{
				Value: parserDollar[1].token.Value.(string),
				Typex: ast.TString,
				Posx:  parserDollar[1].token.Pos,
			}
		}
	}
	goto parserstack /* stack new state and value */
}
