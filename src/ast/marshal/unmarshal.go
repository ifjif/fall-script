package marshal

import (
	"encoding/binary"
	"fmt"
	"reflect"

	"zzc/fall-script/src/ast"
	. "zzc/fall-script/src/ast"
	"zzc/fall-script/src/token"
)

type AstReader struct {
	data []byte
}

func (ar *AstReader) readBytes(c int) []byte {
	d := ar.data[:c]
	ar.data = ar.data[c:]
	return d
}

func (ar *AstReader) readByte() uint8 {
	d := ar.data[0]
	ar.data = ar.data[1:]
	return d
}

func (ar *AstReader) readUin16() uint16 {
	d := binary.BigEndian.Uint16(ar.data)
	ar.data = ar.data[2:]
	return d
}

func (ar *AstReader) readUin32() uint32 {
	d := binary.BigEndian.Uint32(ar.data)
	ar.data = ar.data[4:]
	return d
}

func (ar *AstReader) readUin64() uint64 {
	d := binary.BigEndian.Uint64(ar.data)
	ar.data = ar.data[8:]
	return d
}

func (ar *AstReader) readString() string {
	k := ar.readByte()

	length := 0
	switch k {
	case SHORT_STR:
		length = int(ar.readByte())
	case LONG_STR:
		length = int(ar.readUin32())
	default:
		panic("非法字符串")
	}

	d := ar.readBytes(length)
	return string(d)
}

func (ar *AstReader) readKind() AstKind {
	return AstKind(ar.readByte())
}

func (ar *AstReader) readBool() bool {
	d := ar.readByte()
	v := false
	if d == 1 {
		v = true
	}

	return v
}

func (ar *AstReader) readToken() token.Token {
	ttype := ar.readString()
	value := ar.readString()
	line := ar.readUin32()
	col := ar.readUin32()

	return token.NewToken(token.TokenType(ttype), value, int(line), int(col))
}

func Unmarshal(data []byte) (Node, []byte) {
	ar := &AstReader{data: data}
	node := unmarshal(ar)
	return node, ar.data
}

func unmarshal(ar *AstReader) Node {
	kind := AstKind(ar.readKind())

	switch kind {
	case NIL_K:
		return nil
	case NULL_K:
		return unmarshalNull(ar)
	case INTEGER_k:
		return unmarshalInteger(ar)
	case BOOL_K:
		return unmarshalBool(ar)
	case STRING_K:
		return unmarshalString(ar)
	case IDENT_K:
		return unmarshalIdent(ar)
	case INFIX_K:
		return unmarshalInfix(ar)
	case PREFIX_K:
		return unmarshalPrefix(ar)
	case ARRAY_K:
		return unmarshalArray(ar)
	case HASH_K:
		return unmarshalHash(ar)
	case INDEX_K:
		return unmarshalIndex(ar)
	case SLICE_K:
		return unmarshalSlice(ar)
	case CALL_K:
		return unmarshalCall(ar)
	case FUNCTION_K:
		return unmarshalFunction(ar)
	case ASSIGN_K:
		return unmarshalAssign(ar)
	case IF_K:
		return unmarshalIf(ar)
	case LET_K:
		return unmarshalLet(ar)
	case RETURN_K:
		return unmarshalReturn(ar)
	case FOR_K:
		return unmarshalFor(ar)
	case WHILE_K:
		return unmarshalWhile(ar)
	case DO_WHILE_K:
		return unmarshalDoWhile(ar)
	case BLOCK_K:
		return unmarshalBlock(ar)
	case EXPR_K:
		return unmarshalExpr(ar)
	case PROGRAM_K:
		return unmarshalProgram(ar)
	case ATTR_K:
		return unmarshalAttr(ar)
	}
	return nil
}

func unmarshalElems(r *AstReader, count int) []ExprNode {
	elems := make([]ExprNode, count)

	for i := range count {
		elem := unmarshal(r).(ExprNode)
		elems[i] = elem
	}

	return elems
}

func unmarshalStmts(r *AstReader) []StmtNode {
	count := int(r.readUin32())

	stmts := make([]StmtNode, count)
	for i := range count {
		stmt := unmarshal(r).(StmtNode)
		stmts[i] = stmt
	}

	return stmts
}

func getTypedItems[T any](elems []ExprNode) []T {
	output := make([]T, len(elems))

	for i, elem := range elems {
		v, ok := elem.(T)
		if !ok {
			typeName := reflect.TypeOf((*T)(nil)).Elem().String()
			msg := fmt.Sprintf("类型转换失败, expected: %T, got: %T", typeName, elem)
			panic(msg)
		}
		output[i] = v
	}

	return output
}

func unmarshalPairs(r *AstReader) []*Pair {
	count := int(r.readUin16())
	pairs := make([]*Pair, count)

	for i := range count {
		key := unmarshal(r).(ExprNode)
		value := unmarshal(r).(ExprNode)
		pair := &Pair{Key: key, Value: value}
		pairs[i] = pair
	}

	return pairs
}

func checkKind(t, a AstKind, text string) {
	if t != a {
		msg := fmt.Sprintf("expected %d(%s) get: %d", t, text, a)
		panic(msg)
	}
}

func unmarshalNull(r *AstReader) Node {
	tok := r.readToken()

	return &NullExpr{Token: tok}
}

func unmarshalInteger(r *AstReader) Node {
	tok := r.readToken()

	v := r.readUin64()
	return &IntExpr{Token: tok, Value: int64(v)}
}

func unmarshalBool(r *AstReader) Node {
	tok := r.readToken()

	v := r.readBool()
	return &BoolExpr{Token: tok, Value: v}
}

func unmarshalString(r *AstReader) Node {
	tok := r.readToken()

	str := r.readString()
	return &StrExpr{Token: tok, Value: str}
}

func unmarshalIdent(r *AstReader) Node {
	tok := r.readToken()

	str := r.readString()
	return &IdentExpr{Token: tok, Value: str}
}

func unmarshalInfix(r *AstReader) Node {
	tok := r.readToken()

	left := unmarshal(r).(ExprNode)
	op := r.readString()
	right := unmarshal(r).(ExprNode)

	return &InfixExpr{Token: tok, Left: left, Op: op, Right: right}
}

func unmarshalPrefix(r *AstReader) Node {
	tok := r.readToken()

	op := r.readString()
	right := unmarshal(r).(ExprNode)
	return &PrefixExpr{Token: tok, Op: op, Right: right}
}

func unmarshalArray(r *AstReader) Node {
	tok := r.readToken()
	es := r.readUin16()
	elems := unmarshalElems(r, int(es))

	return &ArrExpr{Token: tok, Elements: elems}
}

func unmarshalHash(r *AstReader) Node {
	tok := r.readToken()

	pairs := unmarshalPairs(r)

	return &HashExpr{Token: tok, Pairs: pairs}
}

func unmarshalIndex(r *AstReader) Node {
	tok := r.readToken()

	left := unmarshal(r).(ExprNode)
	index := unmarshal(r).(ExprNode)

	return &IndexExpr{Token: tok, Left: left, Index: index}
}

func unmarshalSlice(r *AstReader) Node {
	tok := r.readToken()
	se := &SliceExpr{Token: tok}
	start := unmarshal(r)
	if start != nil {
		se.Start = start.(ExprNode)
	}
	end := unmarshal(r)
	if end != nil {
		se.Start = start.(ExprNode)
	}
	step := unmarshal(r)
	if step != nil {
		se.Start = start.(ExprNode)
	}
	capc := unmarshal(r)
	if capc != nil {
		se.Start = start.(ExprNode)
	}

	return se
}

func unmarshalCall(r *AstReader) Node {
	tok := r.readToken()

	fn := unmarshal(r).(ExprNode)
	as := r.readByte()
	args := unmarshalElems(r, int(as))
	return &CallExpr{Token: tok, Fn: fn, Args: args}
}

func unmarshalFunction(r *AstReader) Node {
	tok := r.readToken()

	name := r.readString()
	ident := unmarshal(r)
	ps := r.readByte()
	params := unmarshalElems(r, int(ps))
	np := getTypedItems[*IdentExpr](params)
	body := unmarshal(r)
	unName := r.readBool()
	as := r.readByte()
	attrs := unmarshalElems(r, int(as))
	na := getTypedItems[*AttributeExpr](attrs)

	fn := &FnExpr{
		Token:  tok,
		Name:   name,
		Params: np,
		UnName: unName,
		Attrs:  na,
	}

	if ident == nil {
		fn.Ident = nil
	} else {
		ni := ident.(*ast.IdentExpr)
		fn.Ident = ni
	}

	if body == nil {
		fn.Body = nil
	} else {
		nb := body.(*ast.BlockStmt)
		fn.Body = nb
	}

	return fn
}

func unmarshalAssign(r *AstReader) Node {
	tok := r.readToken()

	left := unmarshal(r).(ExprNode)
	value := unmarshal(r).(ExprNode)

	return &AssignExpr{Token: tok, Left: left, Value: value}
}

func unmarshalIf(r *AstReader) Node {
	tok := r.readToken()

	condition := unmarshal(r).(ExprNode)
	consequence := unmarshal(r).(*BlockStmt)
	alternative := unmarshal(r).(*BlockStmt)
	return &IfExpr{
		Token:       tok,
		Condition:   condition,
		Consequence: consequence,
		Alternative: alternative,
	}
}

func unmarshalLet(r *AstReader) Node {
	tok := r.readToken()

	name := unmarshal(r).(*IdentExpr)
	value := unmarshal(r).(ExprNode)
	return &LetStmt{Token: tok, Name: name, Value: value}
}

func unmarshalReturn(r *AstReader) Node {
	tok := r.readToken()

	value := unmarshal(r).(ExprNode)
	return &ReturnStmt{Token: tok, Value: value}
}

func unmarshalFor(r *AstReader) Node {
	tok := r.readToken()

	start := unmarshal(r).(StmtNode)
	condition := unmarshal(r).(ExprNode)
	update := unmarshal(r).(ExprNode)
	body := unmarshal(r).(*BlockStmt)
	return &ForStmt{
		Token:     tok,
		Start:     start,
		Condition: condition,
		Update:    update,
		Body:      body,
	}
}

func unmarshalWhile(r *AstReader) Node {
	tok := r.readToken()

	condition := unmarshal(r).(ExprNode)
	body := unmarshal(r).(*BlockStmt)

	return &WhileStmt{Token: tok, Condition: condition, Body: body}
}

func unmarshalDoWhile(r *AstReader) Node {
	tok := r.readToken()

	body := unmarshal(r).(*BlockStmt)
	condition := unmarshal(r).(ExprNode)
	return &DoWhileStmt{
		Token:     tok,
		Body:      body,
		Condition: condition,
	}
}

func unmarshalBlock(r *AstReader) Node {
	tok := r.readToken()

	stmts := unmarshalStmts(r)

	return &BlockStmt{Token: tok, Stmts: stmts}
}

func unmarshalExpr(r *AstReader) Node {
	expr := unmarshal(r).(ExprNode)
	return &ExprStmt{Token: expr.GetToken(), Expr: expr}
}

func unmarshalProgram(r *AstReader) Node {
	stmts := unmarshalStmts(r)

	return &Program{Stmts: stmts}
}

func unmarshalAttr(r *AstReader) Node {
	tok := r.readToken()

	name := r.readString()
	count := r.readByte()
	args := unmarshalElems(r, int(count))
	return &AttributeExpr{Token: tok, Name: name, Args: args}
}
