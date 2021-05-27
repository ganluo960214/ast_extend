package ast_extend

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

type AstExtFindStructByNameTestStruct struct{}

var (
	AstExtFindStructByNameTestFile *ast.File
)

func initAstExtFindStructByName(tb testing.TB) {
	var err error
	AstExtFindStructByNameTestFile, err = parser.ParseFile(token.NewFileSet(), "ast_ext_find_struct_test.go", nil, parser.ParseComments)
	if err != nil {
		tb.Fatal(err)
	}
}
func AstExtFindStructByNameCase(tb testing.TB) {
	sn := "AstExtFindStructByNameTestStruct"
	ts, _, ok := AstExtFindStructByName(AstExtFindStructByNameTestFile, sn)
	if ok == false {
		tb.Fatalf("not found %s struct", sn)
	}
	if ts.Name.Name != sn {
		tb.Fatalf("not found %s struct", sn)
	}
}
func TestAstExtFindStructByName(t *testing.T) {
	initAstExtFindStructByName(t)
	AstExtFindStructByNameCase(t)
}
func BenchmarkAstExtFindStructByName(b *testing.B) {
	initAstExtFindStructByName(b)
	for i := 0; i < b.N; i++ {
		AstExtFindStructByNameCase(b)
	}
	b.ReportAllocs()
}
