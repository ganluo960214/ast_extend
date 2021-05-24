package ast_extend

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func AstExtFindFuncByNameTestFunc() {}

type AstExtFindFuncByNameTestStruct struct{}

func (AstExtFindFuncByNameTestStruct) Receiver()         {}
func (*AstExtFindFuncByNameTestStruct) ReceiverPointer() {}

var (
	AstExtFindFuncByNameTestFile *ast.File
)

func initAstExtFindFuncByName(tb testing.TB) {
	var err error
	AstExtFindFuncByNameTestFile, err = parser.ParseFile(token.NewFileSet(), "ast_ext_find_funct_test.go", nil, parser.ParseComments)
	if err != nil {
		tb.Fatal(err)
	}
}
func AstExtFindFuncByNameCaseFunction(tb testing.TB) {
	funcName := "AstExtFindFuncByNameTestFunc"
	d, ok := AstExtFindFuncByName(AstExtFindFuncByNameTestFile, "", funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv != nil {
		tb.Fatal("error func")
	}
}
func AstExtFindFuncByNameCaseReceiverFunction(tb testing.TB) {
	receiverType := "AstExtFindFuncByNameTestStruct"
	funcName := "Receiver"
	d, ok := AstExtFindFuncByName(AstExtFindFuncByNameTestFile, receiverType, funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv.List[0].Type.(*ast.Ident).Name != receiverType {
		tb.Fatal("error func")
	}
}
func AstExtFindFuncByNameCaseReceiverPointerFunction(tb testing.TB) {
	receiverType := "AstExtFindFuncByNameTestStruct"
	funcName := "ReceiverPointer"
	d, ok := AstExtFindFuncByName(AstExtFindFuncByNameTestFile, receiverType, funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name != receiverType {
		tb.Fatal("error func")
	}
}

func TestAstExtFindFuncByNameCaseNormalFunction(t *testing.T) {
	initAstExtFindFuncByName(t)
	AstExtFindFuncByNameCaseFunction(t)
}
func BenchmarkAstExtFindFuncByNameCaseNormalFunction(b *testing.B) {
	initAstExtFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		AstExtFindFuncByNameCaseFunction(b)
	}
	b.ReportAllocs()
}

func TestAstExtFindFuncByNameCaseReceiverFunction(t *testing.T) {
	initAstExtFindFuncByName(t)
	AstExtFindFuncByNameCaseReceiverFunction(t)
}
func BenchmarkAstExtFindFuncByNameCaseReceiverFunction(b *testing.B) {
	initAstExtFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		AstExtFindFuncByNameCaseReceiverFunction(b)
	}
	b.ReportAllocs()
}

func TestAstExtFindFuncByNameCaseReceiverPointerFunction(t *testing.T) {
	initAstExtFindFuncByName(t)
	AstExtFindFuncByNameCaseReceiverPointerFunction(t)
}
func BenchmarkAstExtFindFuncByNameCaseReceiverPointerFunction(b *testing.B) {
	initAstExtFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		AstExtFindFuncByNameCaseReceiverPointerFunction(b)
	}
	b.ReportAllocs()
}
