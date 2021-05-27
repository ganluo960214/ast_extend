package ast_extend

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func FindFuncByNameTestFunc() {}

type FindFuncByNameTestStruct struct{}

func (FindFuncByNameTestStruct) Receiver()         {}
func (*FindFuncByNameTestStruct) ReceiverPointer() {}

var (
	FindFuncByNameTestFile *ast.File
)

func initTestFindFuncByName(tb testing.TB) {
	var err error
	FindFuncByNameTestFile, err = parser.ParseFile(token.NewFileSet(), "find_func_by_name_test.go", nil, parser.ParseComments)
	if err != nil {
		tb.Fatal(err)
	}
}
func FindFuncByNameCaseFunction(tb testing.TB) {
	funcName := "FindFuncByNameTestFunc"
	d, ok := FindFuncByName(FindFuncByNameTestFile, "", funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv != nil {
		tb.Fatal("error func")
	}
}
func FindFuncByNameCaseReceiverFunction(tb testing.TB) {
	receiverType := "FindFuncByNameTestStruct"
	funcName := "Receiver"
	d, ok := FindFuncByName(FindFuncByNameTestFile, receiverType, funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv.List[0].Type.(*ast.Ident).Name != receiverType {
		tb.Fatal("error func")
	}
}
func FindFuncByNameCaseReceiverPointerFunction(tb testing.TB) {
	receiverType := "FindFuncByNameTestStruct"
	funcName := "ReceiverPointer"
	d, ok := FindFuncByName(FindFuncByNameTestFile, receiverType, funcName)
	if ok == false {
		tb.Fatal("not found function")
	}
	if d.Name.Name != funcName || d.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name != receiverType {
		tb.Fatal("error func")
	}
}

func TestFindFuncByNameCaseNormalFunction(t *testing.T) {
	initTestFindFuncByName(t)
	FindFuncByNameCaseFunction(t)
}
func BenchmarkFindFuncByNameCaseNormalFunction(b *testing.B) {
	initTestFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		FindFuncByNameCaseFunction(b)
	}
	b.ReportAllocs()
}

func TestFindFuncByNameCaseReceiverFunction(t *testing.T) {
	initTestFindFuncByName(t)
	FindFuncByNameCaseReceiverFunction(t)
}
func BenchmarkFindFuncByNameCaseReceiverFunction(b *testing.B) {
	initTestFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		FindFuncByNameCaseReceiverFunction(b)
	}
	b.ReportAllocs()
}

func TestFindFuncByNameCaseReceiverPointerFunction(t *testing.T) {
	initTestFindFuncByName(t)
	FindFuncByNameCaseReceiverPointerFunction(t)
}
func BenchmarkFindFuncByNameCaseReceiverPointerFunction(b *testing.B) {
	initTestFindFuncByName(b)
	for i := 0; i < b.N; i++ {
		FindFuncByNameCaseReceiverPointerFunction(b)
	}
	b.ReportAllocs()
}
