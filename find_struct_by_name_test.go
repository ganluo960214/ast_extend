package ast_extend

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

type FindStructByNameTestStruct struct{}

var (
	FindStructByNameTestFile *ast.File
)

func initTestFindStructByName(tb testing.TB) {
	var err error
	FindStructByNameTestFile, err = parser.ParseFile(token.NewFileSet(), "find_struct_by_name_test.go", nil, parser.ParseComments)
	if err != nil {
		tb.Fatal(err)
	}
}
func FindStructByNameCase(tb testing.TB) {
	sn := "FindStructByNameTestStruct"
	ts, _, ok := FindStructByName(FindStructByNameTestFile, sn)
	if ok == false {
		tb.Fatalf("not found %s struct", sn)
	}
	if ts.Name.Name != sn {
		tb.Fatalf("not found %s struct", sn)
	}
}
func TestFindStructByName(t *testing.T) {
	initTestFindStructByName(t)
	FindStructByNameCase(t)
}
func BenchmarkFindStructByName(b *testing.B) {
	initTestFindStructByName(b)
	for i := 0; i < b.N; i++ {
		FindStructByNameCase(b)
	}
	b.ReportAllocs()
}
