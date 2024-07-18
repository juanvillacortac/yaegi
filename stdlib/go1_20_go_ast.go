// Code generated by 'yaegi extract go/ast'. DO NOT EDIT.

//go:build go1.20 && !go1.21
// +build go1.20,!go1.21

package stdlib

import (
	"go/ast"
	"go/token"
	"reflect"
)

func init() {
	Symbols["go/ast/ast"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Bad":                        reflect.ValueOf(ast.Bad),
		"Con":                        reflect.ValueOf(ast.Con),
		"FileExports":                reflect.ValueOf(ast.FileExports),
		"FilterDecl":                 reflect.ValueOf(ast.FilterDecl),
		"FilterFile":                 reflect.ValueOf(ast.FilterFile),
		"FilterFuncDuplicates":       reflect.ValueOf(ast.FilterFuncDuplicates),
		"FilterImportDuplicates":     reflect.ValueOf(ast.FilterImportDuplicates),
		"FilterPackage":              reflect.ValueOf(ast.FilterPackage),
		"FilterUnassociatedComments": reflect.ValueOf(ast.FilterUnassociatedComments),
		"Fprint":                     reflect.ValueOf(ast.Fprint),
		"Fun":                        reflect.ValueOf(ast.Fun),
		"Inspect":                    reflect.ValueOf(ast.Inspect),
		"IsExported":                 reflect.ValueOf(ast.IsExported),
		"Lbl":                        reflect.ValueOf(ast.Lbl),
		"MergePackageFiles":          reflect.ValueOf(ast.MergePackageFiles),
		"NewCommentMap":              reflect.ValueOf(ast.NewCommentMap),
		"NewIdent":                   reflect.ValueOf(ast.NewIdent),
		"NewObj":                     reflect.ValueOf(ast.NewObj),
		"NewPackage":                 reflect.ValueOf(ast.NewPackage),
		"NewScope":                   reflect.ValueOf(ast.NewScope),
		"NotNilFilter":               reflect.ValueOf(ast.NotNilFilter),
		"PackageExports":             reflect.ValueOf(ast.PackageExports),
		"Pkg":                        reflect.ValueOf(ast.Pkg),
		"Print":                      reflect.ValueOf(ast.Print),
		"RECV":                       reflect.ValueOf(ast.RECV),
		"SEND":                       reflect.ValueOf(ast.SEND),
		"SortImports":                reflect.ValueOf(ast.SortImports),
		"Typ":                        reflect.ValueOf(ast.Typ),
		"Var":                        reflect.ValueOf(ast.Var),
		"Walk":                       reflect.ValueOf(ast.Walk),

		// type definitions
		"ArrayType":      reflect.ValueOf((*ast.ArrayType)(nil)),
		"AssignStmt":     reflect.ValueOf((*ast.AssignStmt)(nil)),
		"BadDecl":        reflect.ValueOf((*ast.BadDecl)(nil)),
		"BadExpr":        reflect.ValueOf((*ast.BadExpr)(nil)),
		"BadStmt":        reflect.ValueOf((*ast.BadStmt)(nil)),
		"BasicLit":       reflect.ValueOf((*ast.BasicLit)(nil)),
		"BinaryExpr":     reflect.ValueOf((*ast.BinaryExpr)(nil)),
		"BlockStmt":      reflect.ValueOf((*ast.BlockStmt)(nil)),
		"BranchStmt":     reflect.ValueOf((*ast.BranchStmt)(nil)),
		"CallExpr":       reflect.ValueOf((*ast.CallExpr)(nil)),
		"CaseClause":     reflect.ValueOf((*ast.CaseClause)(nil)),
		"ChanDir":        reflect.ValueOf((*ast.ChanDir)(nil)),
		"ChanType":       reflect.ValueOf((*ast.ChanType)(nil)),
		"CommClause":     reflect.ValueOf((*ast.CommClause)(nil)),
		"Comment":        reflect.ValueOf((*ast.Comment)(nil)),
		"CommentGroup":   reflect.ValueOf((*ast.CommentGroup)(nil)),
		"CommentMap":     reflect.ValueOf((*ast.CommentMap)(nil)),
		"CompositeLit":   reflect.ValueOf((*ast.CompositeLit)(nil)),
		"Decl":           reflect.ValueOf((*ast.Decl)(nil)),
		"DeclStmt":       reflect.ValueOf((*ast.DeclStmt)(nil)),
		"DeferStmt":      reflect.ValueOf((*ast.DeferStmt)(nil)),
		"Ellipsis":       reflect.ValueOf((*ast.Ellipsis)(nil)),
		"EmptyStmt":      reflect.ValueOf((*ast.EmptyStmt)(nil)),
		"Expr":           reflect.ValueOf((*ast.Expr)(nil)),
		"ExprStmt":       reflect.ValueOf((*ast.ExprStmt)(nil)),
		"Field":          reflect.ValueOf((*ast.Field)(nil)),
		"FieldFilter":    reflect.ValueOf((*ast.FieldFilter)(nil)),
		"FieldList":      reflect.ValueOf((*ast.FieldList)(nil)),
		"File":           reflect.ValueOf((*ast.File)(nil)),
		"Filter":         reflect.ValueOf((*ast.Filter)(nil)),
		"ForStmt":        reflect.ValueOf((*ast.ForStmt)(nil)),
		"FuncDecl":       reflect.ValueOf((*ast.FuncDecl)(nil)),
		"FuncLit":        reflect.ValueOf((*ast.FuncLit)(nil)),
		"FuncType":       reflect.ValueOf((*ast.FuncType)(nil)),
		"GenDecl":        reflect.ValueOf((*ast.GenDecl)(nil)),
		"GoStmt":         reflect.ValueOf((*ast.GoStmt)(nil)),
		"Ident":          reflect.ValueOf((*ast.Ident)(nil)),
		"IfStmt":         reflect.ValueOf((*ast.IfStmt)(nil)),
		"ImportSpec":     reflect.ValueOf((*ast.ImportSpec)(nil)),
		"Importer":       reflect.ValueOf((*ast.Importer)(nil)),
		"IncDecStmt":     reflect.ValueOf((*ast.IncDecStmt)(nil)),
		"IndexExpr":      reflect.ValueOf((*ast.IndexExpr)(nil)),
		"IndexListExpr":  reflect.ValueOf((*ast.IndexListExpr)(nil)),
		"InterfaceType":  reflect.ValueOf((*ast.InterfaceType)(nil)),
		"KeyValueExpr":   reflect.ValueOf((*ast.KeyValueExpr)(nil)),
		"LabeledStmt":    reflect.ValueOf((*ast.LabeledStmt)(nil)),
		"MapType":        reflect.ValueOf((*ast.MapType)(nil)),
		"MergeMode":      reflect.ValueOf((*ast.MergeMode)(nil)),
		"Node":           reflect.ValueOf((*ast.Node)(nil)),
		"ObjKind":        reflect.ValueOf((*ast.ObjKind)(nil)),
		"Object":         reflect.ValueOf((*ast.Object)(nil)),
		"Package":        reflect.ValueOf((*ast.Package)(nil)),
		"ParenExpr":      reflect.ValueOf((*ast.ParenExpr)(nil)),
		"RangeStmt":      reflect.ValueOf((*ast.RangeStmt)(nil)),
		"ReturnStmt":     reflect.ValueOf((*ast.ReturnStmt)(nil)),
		"Scope":          reflect.ValueOf((*ast.Scope)(nil)),
		"SelectStmt":     reflect.ValueOf((*ast.SelectStmt)(nil)),
		"SelectorExpr":   reflect.ValueOf((*ast.SelectorExpr)(nil)),
		"SendStmt":       reflect.ValueOf((*ast.SendStmt)(nil)),
		"SliceExpr":      reflect.ValueOf((*ast.SliceExpr)(nil)),
		"Spec":           reflect.ValueOf((*ast.Spec)(nil)),
		"StarExpr":       reflect.ValueOf((*ast.StarExpr)(nil)),
		"Stmt":           reflect.ValueOf((*ast.Stmt)(nil)),
		"StructType":     reflect.ValueOf((*ast.StructType)(nil)),
		"SwitchStmt":     reflect.ValueOf((*ast.SwitchStmt)(nil)),
		"TypeAssertExpr": reflect.ValueOf((*ast.TypeAssertExpr)(nil)),
		"TypeSpec":       reflect.ValueOf((*ast.TypeSpec)(nil)),
		"TypeSwitchStmt": reflect.ValueOf((*ast.TypeSwitchStmt)(nil)),
		"UnaryExpr":      reflect.ValueOf((*ast.UnaryExpr)(nil)),
		"ValueSpec":      reflect.ValueOf((*ast.ValueSpec)(nil)),
		"Visitor":        reflect.ValueOf((*ast.Visitor)(nil)),

		// interface wrapper definitions
		"_Decl":    reflect.ValueOf((*_go_ast_Decl)(nil)),
		"_Expr":    reflect.ValueOf((*_go_ast_Expr)(nil)),
		"_Node":    reflect.ValueOf((*_go_ast_Node)(nil)),
		"_Spec":    reflect.ValueOf((*_go_ast_Spec)(nil)),
		"_Stmt":    reflect.ValueOf((*_go_ast_Stmt)(nil)),
		"_Visitor": reflect.ValueOf((*_go_ast_Visitor)(nil)),
	}
}

// _go_ast_Decl is an interface wrapper for Decl type
type _go_ast_Decl struct {
	IValue interface{}
	WEnd   func() token.Pos
	WPos   func() token.Pos
}

func (W _go_ast_Decl) End() token.Pos { return W.WEnd() }
func (W _go_ast_Decl) Pos() token.Pos { return W.WPos() }

// _go_ast_Expr is an interface wrapper for Expr type
type _go_ast_Expr struct {
	IValue interface{}
	WEnd   func() token.Pos
	WPos   func() token.Pos
}

func (W _go_ast_Expr) End() token.Pos { return W.WEnd() }
func (W _go_ast_Expr) Pos() token.Pos { return W.WPos() }

// _go_ast_Node is an interface wrapper for Node type
type _go_ast_Node struct {
	IValue interface{}
	WEnd   func() token.Pos
	WPos   func() token.Pos
}

func (W _go_ast_Node) End() token.Pos { return W.WEnd() }
func (W _go_ast_Node) Pos() token.Pos { return W.WPos() }

// _go_ast_Spec is an interface wrapper for Spec type
type _go_ast_Spec struct {
	IValue interface{}
	WEnd   func() token.Pos
	WPos   func() token.Pos
}

func (W _go_ast_Spec) End() token.Pos { return W.WEnd() }
func (W _go_ast_Spec) Pos() token.Pos { return W.WPos() }

// _go_ast_Stmt is an interface wrapper for Stmt type
type _go_ast_Stmt struct {
	IValue interface{}
	WEnd   func() token.Pos
	WPos   func() token.Pos
}

func (W _go_ast_Stmt) End() token.Pos { return W.WEnd() }
func (W _go_ast_Stmt) Pos() token.Pos { return W.WPos() }

// _go_ast_Visitor is an interface wrapper for Visitor type
type _go_ast_Visitor struct {
	IValue interface{}
	WVisit func(node ast.Node) (w ast.Visitor)
}

func (W _go_ast_Visitor) Visit(node ast.Node) (w ast.Visitor) { return W.WVisit(node) }
