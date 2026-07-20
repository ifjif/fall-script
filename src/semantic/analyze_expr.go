package semantic

import (
	"zzc/fall-script/src/ast"
	"zzc/fall-script/src/ir"
	"zzc/fall-script/src/utils"
)

func (a *Analyzer) analyzeExpr(expr ast.ExprNode) ir.Expr {
	switch expr := expr.(type) {
	case *ast.NullExpr:
		return a.analyzeNullExpr(expr)
	case *ast.IntExpr:
		return a.analyzeIntExpr(expr)
	case *ast.BoolExpr:
		return a.analyzeBoolExpr(expr)
	case *ast.StrExpr:
		return a.analyzeStrExpr(expr)
	case *ast.IdentExpr:
		return a.analyzeIdentExpr(expr)
	case *ast.ArrExpr:
		return a.analyzeArrExpr(expr)
	case *ast.HashExpr:
		return a.analyzeHashExpr(expr)
	case *ast.InfixExpr:
		return a.analyzeInfixExpr(expr)
	case *ast.PrefixExpr:
		return a.analyzePrefixExpr(expr)
	case *ast.SliceExpr:
		return a.analyzeSliceExpr(expr)
	case *ast.IndexExpr:
		return a.analyzeIndexExpr(expr)
	case *ast.MemberExpr:
		return a.analyzeMemberExpr(expr)
	case *ast.AssignExpr:
		return a.analyzeAssignExpr(expr)
	case *ast.IfExpr:
		return a.analyzeIfExpr(expr)
	case *ast.CallExpr:
		return a.analyzeCallExpr(expr)
	case *ast.FnExpr:
		return a.analyzeFnExpr(expr)
	case *ast.StructLiteralExpr:
		return a.analyzeStructLiteralExpr(expr)
	case *ast.MatchExpr:
		return a.analyzeMatchExpr(expr)
	}

	return nil
}

func (a *Analyzer) analyzeIntExpr(expr *ast.IntExpr) *ir.IntegerLiteral {
	return &ir.IntegerLiteral{Token: expr.Token, Value: expr.Value}
}

func (a *Analyzer) analyzeBoolExpr(expr *ast.BoolExpr) *ir.BoolLiteral {
	return &ir.BoolLiteral{Token: expr.Token, Value: expr.Value}
}

func (a *Analyzer) analyzeNullExpr(expr *ast.NullExpr) *ir.NullLiteral {
	return &ir.NullLiteral{Token: expr.Token}
}

func (a *Analyzer) analyzeStrExpr(expr *ast.StrExpr) *ir.StringLiteral {
	return &ir.StringLiteral{Token: expr.Token, Value: expr.Value}
}

func (a *Analyzer) analyzeIdentExpr(expr *ast.IdentExpr) *ir.Ident {
	value := expr.Value

	symbol, ok := a.SymbolTable.Resolve(value)

	if !ok {
		err := utils.IdentifierNotFoundErr(value)
		a.errors = append(a.errors, appendPosInfo(expr.Token, err))
		return nil
	}

	return &ir.Ident{Token: expr.Token, Symbol: symbol}
}

func (a *Analyzer) analyzeArrExpr(expr *ast.ArrExpr) *ir.ArrayLiteral {
	elems := make([]ir.Expr, len(expr.Elements))
	for i, elem := range expr.Elements {
		ielem := a.analyzeExpr(elem)
		if ielem != nil {
			elems[i] = ielem
		}
	}

	return &ir.ArrayLiteral{Token: expr.Token, Elements: elems}
}

func (a *Analyzer) analyzeHashExpr(expr *ast.HashExpr) *ir.HashLiteral {
	pairs := make([]*ir.Pair, len(expr.Pairs))

	for i, pair := range expr.Pairs {
		key := a.analyzeExpr(pair.Key)
		if key == nil {
			continue
		}
		value := a.analyzeExpr(pair.Value)
		if value == nil {
			continue
		}

		ipair := &ir.Pair{Key: key, Value: value}
		pairs[i] = ipair
	}

	return &ir.HashLiteral{Token: expr.Token, Pairs: pairs}
}

func (a *Analyzer) analyzeInfixExpr(expr *ast.InfixExpr) *ir.InfixExpr {
	op := expr.Op
	left := a.analyzeExpr(expr.Left)
	if left == nil {
		return nil
	}

	right := a.analyzeExpr(expr.Right)
	if right == nil {
		return nil
	}

	return &ir.InfixExpr{Token: expr.Token, Op: op, Left: left, Right: right}
}

func (a *Analyzer) analyzePrefixExpr(expr *ast.PrefixExpr) *ir.PrefixExpr {
	op := expr.Op
	right := a.analyzeExpr(expr.Right)
	if right == nil {
		return nil
	}

	return &ir.PrefixExpr{Token: expr.Token, Op: op, Right: right}
}

func (a *Analyzer) analyzeSliceExpr(expr *ast.SliceExpr) *ir.SliceExpr {
	var start ir.Expr = &ir.NullLiteral{}
	var end ir.Expr = &ir.NullLiteral{}
	var step ir.Expr = &ir.NullLiteral{}
	var cap_ ir.Expr = &ir.NullLiteral{}

	if expr.Start != nil {
		start = a.analyzeExpr(expr.Start)
	}

	if expr.End != nil {
		end = a.analyzeExpr(expr.End)
	}

	if expr.Step != nil {
		step = a.analyzeExpr(expr.Step)
	}

	if expr.Cap != nil {
		cap_ = a.analyzeExpr(expr.Cap)
	}

	return &ir.SliceExpr{Token: expr.Token, Start: start, End: end, Step: step, Cap: cap_}
}

func (a *Analyzer) analyzeIndexExpr(expr *ast.IndexExpr) *ir.IndexExpr {
	left := a.analyzeExpr(expr.Left)
	index := a.analyzeExpr(expr.Index)

	return &ir.IndexExpr{Token: expr.Token, Left: left, Index: index}
}

func (a *Analyzer) analyzeMemberExpr(expr *ast.MemberExpr) *ir.MemberExpr {
	visitor := a.analyzeExpr(expr.Visitor)
	var member ir.Expr
	memb := expr.Member
	switch memb := memb.(type) {
	case *ast.IdentExpr:
		member = &ir.StringLiteral{Token: memb.Token, Value: memb.Value}
	case *ast.IntExpr:
		member = &ir.IntegerLiteral{Token: memb.Token, Value: memb.Value}
	}

	return &ir.MemberExpr{Token: expr.Token, Visitor: visitor, Member: member}
}

func (a *Analyzer) analyzeAssignExpr(expr *ast.AssignExpr) *ir.AssignExpr {
	left := a.analyzeExpr(expr.Left)
	value := a.analyzeExpr(expr.Value)

	return &ir.AssignExpr{Token: expr.Token, Left: left, Value: value}
}

func (a *Analyzer) analyzeIfExpr(expr *ast.IfExpr) *ir.IfExpr {
	condition := a.analyzeExpr(expr.Condition)
	consequence := a.analyzeBlockStmt(expr.Consequence)
	var alternative *ir.BlockStmt
	if expr.Alternative != nil {
		alternative = a.analyzeBlockStmt(expr.Alternative)
	}

	return &ir.IfExpr{Token: expr.Token, Condition: condition, Consequence: consequence, Alternative: alternative}
}

func (a *Analyzer) analyzeCallExpr(expr *ast.CallExpr) *ir.CallExpr {
	callee := a.analyzeExpr(expr.Fn)
	args := make([]ir.Expr, len(expr.Args))
	for i, arg := range expr.Args {
		args[i] = a.analyzeExpr(arg)
	}

	return &ir.CallExpr{Token: expr.Token, Callee: callee, Args: args}
}

func (a *Analyzer) analyzeFnExpr(expr *ast.FnExpr) *ir.FnExpr {
	var name *ir.Ident
	if !expr.UnName {
		// 顶部 函数被提升
		nameSym, ok := a.SymbolTable.ResolveSelf(expr.Name)
		if !ok {
			nameSym = a.SymbolTable.Define(expr.Ident.Value)
		}
		name = &ir.Ident{Token: expr.Ident.Token, Symbol: nameSym}
	}

	a.enterFunction()
	if expr.Name != "" {
		a.SymbolTable.DefineFunction(expr.Name)
	}
	params := make([]*ir.Ident, len(expr.Params))
	for i, param := range expr.Params {
		symbol := a.SymbolTable.Define(param.Value)
		params[i] = &ir.Ident{Token: param.Token, Symbol: symbol}
	}

	body := a.analyzeBlockStmt(expr.Body)

	freeTables := a.SymbolTable.FreeSymbols
	// 占时将 captured 变为false
	// 等free都被编译完后，再变为true
	for _, free := range freeTables {
		free.Captured = false
	}
	localVars := a.SymbolTable.MaxNum()
	a.leaveFunction()

	return &ir.FnExpr{
		Token:     expr.Token,
		Name:      name,
		UnName:    expr.UnName,
		Params:    params,
		Body:      body,
		LocalVars: localVars,
		Frees:     freeTables,
	}
}

func (a *Analyzer) analyzeStructLiteralExpr(expr *ast.StructLiteralExpr) *ir.StructLiteral {
	tag := a.analyzeIdentExpr(expr.Tag)
	pairs := make([]*ir.Pair, len(expr.Elements))

	for i, elem := range expr.Elements {
		key := &ir.StringLiteral{Token: elem.Key.GetToken(), Value: elem.Key.Value}
		value := a.analyzeExpr(elem.Value)
		pairs[i] = &ir.Pair{Key: key, Value: value}
	}

	return &ir.StructLiteral{Token: expr.Token, Tag: tag, Pairs: pairs}
}

func (a *Analyzer) analyzeMatchExpr(expr *ast.MatchExpr) *ir.MatchExpr {
	subject := a.analyzeExpr(expr.Subject)
	arms := a.analyzeMatchArms(expr.Arms)

	return &ir.MatchExpr{Token: expr.Token, Subject: subject, MatchArms: arms}
}

func (a *Analyzer) analyzeMatchArms(arms []*ast.MatchArmExpr) []*ir.MatchArmExpr {
	iarms := make([]*ir.MatchArmExpr, len(arms))
	for i, arm := range arms {
		iarm := a.analyzeMatchArm(arm)
		iarms[i] = iarm
	}

	return iarms
}

func (a *Analyzer) analyzeMatchArm(arm *ast.MatchArmExpr) *ir.MatchArmExpr {
	a.enterBlock()
	pattern := a.analyzePattern(arm.Pattern)
	guard := a.analyzeExpr(arm.Guard)
	body := a.analyzeStmt(arm.Body)
	// 空块，添加null
	if block, ok := body.(*ir.BlockStmt); ok {
		if len(block.Stmts) == 0 {
			block.Stmts = append(block.Stmts, &ir.ExprStmt{Token: block.Token, Expr: &ir.NullLiteral{Token: block.Token}})
		}
	}
	a.leaveBlock()

	return &ir.MatchArmExpr{Token: arm.Token, Pattern: pattern, Guard: guard, Body: body}
}
