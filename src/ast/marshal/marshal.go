package marshal

import (
	. "zzc/fall-script/src/ast"
	binarychunck "zzc/fall-script/src/binary_chunck"
	"zzc/fall-script/src/token"
)

const (
	SHORT_STR = 0
	LONG_STR  = 1
)

type AstWriter struct {
	binarychunck.FlWriter
}

func (aw *AstWriter) writeKind(k AstKind) {
	aw.WriteUint8(byte(k))
}

func (aw *AstWriter) writeToken(token token.Token) {
	aw.WriteStr(string(token.Type))
	aw.WriteStr(token.Value)
	aw.WriteUint32(uint32(token.Line))
	aw.WriteUint32(uint32(token.Col))
}

func MarshalAst(node Node) []byte {
	var awr AstWriter
	marshalAst(node, &awr)
	return awr.Bytes()
}

func isNil(node Node) bool {
	// 处理 类型+值都为nil
	if node == nil {
		return true
	}

	// 处理值nil
	switch node := node.(type) {
	case *Program:
		if node == nil {
			return true
		}
	case *BlockStmt:
		if node == nil {
			return true
		}
	case *ExprStmt:

		if node == nil {
			return true
		}
	case *LetStmt:
		if node == nil {
			return true
		}
	case *ReturnStmt:
		if node == nil {
			return true
		}
	case *ForStmt:
		if node == nil {
			return true
		}
	case *WhileStmt:
		if node == nil {
			return true
		}
	case *DoWhileStmt:
		if node == nil {
			return true
		}
	case *NullExpr:
		if node == nil {
			return true
		}
	case *IntExpr:
		if node == nil {
			return true
		}
	case *StrExpr:
		if node == nil {
			return true
		}
	case *BoolExpr:
		if node == nil {
			return true
		}
	case *IdentExpr:
		if node == nil {
			return true
		}
	case *AssignExpr:
		if node == nil {
			return true
		}
	case *IndexExpr:
		if node == nil {
			return true
		}
	case *SliceExpr:
		if node == nil {
			return true
		}
	case *InfixExpr:
		if node == nil {
			return true
		}
	case *PrefixExpr:
		if node == nil {
			return true
		}
	case *CallExpr:
		if node == nil {
			return true
		}
	case *ArrExpr:
		if node == nil {
			return true
		}
	case *HashExpr:
		if node == nil {
			return true
		}
	case *FnExpr:
		if node == nil {
			return true
		}
	case *IfExpr:
		if node == nil {
			return true
		}
	case *StructDeclStmt:
		if node == nil {
			return true
		}
	}

	return false
}

func marshalAst(node Node, awr *AstWriter) {
	if isNil(node) {
		awr.writeKind(NIL_K)
		return
	}

	switch node := node.(type) {

	case *Program:
		marshalProgram(node, awr)
	case *BlockStmt:
		marshalBlock(node, awr)
	case *ExprStmt:
		marshalExpr(node, awr)
	case *LetStmt:
		marshalLet(node, awr)
	case *ReturnStmt:
		marshalReturn(node, awr)
	case *ForStmt:
		marshalFor(node, awr)
	case *WhileStmt:
		marshalWhile(node, awr)
	case *DoWhileStmt:
		marshalDoWhile(node, awr)
	case *NullExpr:
		marshalNull(node, awr)
	case *IntExpr:
		marshalInt(node, awr)
	case *StrExpr:
		marshalStr(node, awr)
	case *BoolExpr:
		marshalBool(node, awr)
	case *IdentExpr:
		marshalIdent(node, awr)
	case *AssignExpr:
		marshalAssign(node, awr)
	case *IndexExpr:
		marshalIndex(node, awr)
	case *SliceExpr:
		marshalSlice(node, awr)
	case *InfixExpr:
		marshalInfix(node, awr)
	case *PrefixExpr:
		marshalPrefix(node, awr)
	case *CallExpr:
		marshalCall(node, awr)
	case *ArrExpr:
		marshalArr(node, awr)
	case *HashExpr:
		marshalHash(node, awr)
	case *FnExpr:
		marshalFn(node, awr)
	case *IfExpr:
		marshalIf(node, awr)
	case *StructDeclStmt:
		marshalStructDeclStmt(node, awr)
	}
}

func marshalProgram(node *Program, w *AstWriter) {
	w.writeKind(PROGRAM_K)
	length := len(node.Stmts)
	w.WriteUint32(uint32(length))
	for _, stmt := range node.Stmts {
		marshalAst(stmt, w)
	}
}

func marshalBlock(node *BlockStmt, w *AstWriter) {
	w.writeKind(BLOCK_K)
	w.writeToken(node.Token)
	length := len(node.Stmts)
	w.WriteUint32(uint32(length))
	for _, stmt := range node.Stmts {
		marshalAst(stmt, w)
	}
}

func marshalExpr(node *ExprStmt, w *AstWriter) {
	w.writeKind(EXPR_K)
	marshalAst(node.Expr, w)
}

func marshalLet(node *LetStmt, w *AstWriter) {
	w.writeKind(LET_K)
	w.writeToken(node.Token)
	marshalAst(node.Name, w)
	marshalAst(node.Value, w)
}

func marshalReturn(node *ReturnStmt, w *AstWriter) {
	w.writeKind(RETURN_K)
	w.writeToken(node.Token)
	marshalAst(node.Value, w)
}

func marshalFor(node *ForStmt, w *AstWriter) {
	w.writeKind(FOR_K)
	w.writeToken(node.Token)
	marshalAst(node.Start, w)
	marshalAst(node.Condition, w)
	marshalAst(node.Update, w)
	marshalAst(node.Body, w)
}

func marshalWhile(node *WhileStmt, w *AstWriter) {
	w.writeKind(WHILE_K)
	w.writeToken(node.Token)
	marshalAst(node.Condition, w)
	marshalAst(node.Body, w)
}

func marshalDoWhile(node *DoWhileStmt, w *AstWriter) {
	w.writeKind(DO_WHILE_K)
	w.writeToken(node.Token)
	marshalAst(node.Body, w)
	marshalAst(node.Condition, w)
}

func marshalNull(node *NullExpr, w *AstWriter) {
	w.writeKind(NULL_K)
	w.writeToken(node.Token)
}

func marshalInt(node *IntExpr, w *AstWriter) {
	w.writeKind(INTEGER_k)
	w.writeToken(node.Token)
	w.WriteUint64(uint64(node.Value))
}

func marshalStr(node *StrExpr, w *AstWriter) {
	w.writeKind(STRING_K)
	w.writeToken(node.Token)
	w.WriteStr(node.Value)
}

func marshalBool(node *BoolExpr, w *AstWriter) {
	w.writeKind(BOOL_K)
	w.writeToken(node.Token)
	w.WriteBool(node.Value)
}

func marshalIdent(node *IdentExpr, w *AstWriter) {
	w.writeKind(IDENT_K)
	w.writeToken(node.Token)
	w.WriteStr(node.Value)
}

func marshalAssign(node *AssignExpr, w *AstWriter) {
	w.writeKind(ASSIGN_K)
	w.writeToken(node.Token)
	marshalAst(node.Left, w)
	marshalAst(node.Value, w)
}

func marshalIndex(node *IndexExpr, w *AstWriter) {
	w.writeKind(INDEX_K)
	w.writeToken(node.Token)
	marshalAst(node.Left, w)
	marshalAst(node.Index, w)
}

func marshalSlice(node *SliceExpr, w *AstWriter) {
	w.writeKind(SLICE_K)
	w.writeToken(node.Token)
	marshalAst(node.Start, w)
	marshalAst(node.End, w)
	marshalAst(node.Step, w)
	marshalAst(node.Cap, w)
}

func marshalInfix(node *InfixExpr, w *AstWriter) {
	w.writeKind(INFIX_K)
	w.writeToken(node.Token)
	marshalAst(node.Left, w)
	w.WriteStr(node.Op)
	marshalAst(node.Right, w)
}

func marshalPrefix(node *PrefixExpr, w *AstWriter) {
	w.writeKind(PREFIX_K)
	w.writeToken(node.Token)
	w.WriteStr(node.Op)
	marshalAst(node.Right, w)
}

func marshalCall(node *CallExpr, w *AstWriter) {
	w.writeKind(CALL_K)
	w.writeToken(node.Token)
	marshalAst(node.Fn, w)
	w.WriteUint8(byte(len(node.Args)))
	for _, arg := range node.Args {
		marshalAst(arg, w)
	}
}

func marshalArr(node *ArrExpr, w *AstWriter) {
	w.writeKind(ARRAY_K)
	w.writeToken(node.Token)
	w.WriteUint16(uint16(len(node.Elements)))
	for _, ele := range node.Elements {
		marshalAst(ele, w)
	}
}

func marshalHash(node *HashExpr, w *AstWriter) {
	w.writeKind(HASH_K)
	w.writeToken(node.Token)

	w.WriteUint16(uint16(len(node.Pairs)))
	for _, pair := range node.Pairs {
		marshalAst(pair.Key, w)
		marshalAst(pair.Value, w)
	}
}

func marshalFn(node *FnExpr, w *AstWriter) {
	w.writeKind(FUNCTION_K)
	w.writeToken(node.Token)
	w.WriteStr(node.Name)
	marshalAst(node.Ident, w)
	w.WriteUint8(uint8(len(node.Params)))
	for _, param := range node.Params {
		marshalAst(param, w)
	}
	marshalAst(node.Body, w)
	w.WriteBool(node.UnName)
	w.WriteUint8(uint8(len(node.Attrs)))
	for _, attr := range node.Attrs {
		marshalAttr(attr, w)
	}
}

func marshalAttr(node *AttributeExpr, w *AstWriter) {
	w.writeKind(ATTR_K)
	w.writeToken(node.Token)
	w.WriteStr(node.Name)
	w.WriteUint8(uint8(len(node.Args)))
	for _, arg := range node.Args {
		marshalAst(arg, w)
	}
}

func marshalIf(node *IfExpr, w *AstWriter) {
	w.writeKind(IF_K)
	w.writeToken(node.Token)
	marshalAst(node.Condition, w)
	marshalAst(node.Consequence, w)
	marshalAst(node.Alternative, w)
}

func marshalStructDeclStmt(node *StructDeclStmt, w *AstWriter) {
	w.writeKind(STRUCT_K)
	w.writeToken(node.Token)
	marshalIdent(node.Name, w)
	w.WriteUint16(uint16(len(node.Fields)))
	for _, f := range node.Fields {
		marshalFieldDeclExpr(f, w)
	}
}

func marshalFieldDeclExpr(node *FieldDeclExpr, w *AstWriter) {
	w.writeToken(node.Token)
	marshalIdent(node.Name, w)
	w.WriteBool(node.IsEmbed)
}
