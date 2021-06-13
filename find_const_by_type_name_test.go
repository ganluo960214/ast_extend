package ast_extend

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"testing"
)

type Type byte

const (
	TypeCA Type = iota
	TypeCB
	NotTypeCA = 2
	NotTypeCB = iota
	TypeCC    = Type(iota)
	TypeCD
	NotTypeCC      = int(iota)
	NotTypeCD      = byte(7)
	TypeCE    Type = 8
	TypeCF    Type = 9
)

const (
	Test_FindConstByType_FileName = "find_const_by_type_name_test.go"
	Test_FindConstByType_TypeName = "Type"
)

var (
	Test_FindConstByType_AstFile *ast.File
)

func initFindConstByTypeName(tb testing.TB) {
	var err error
	Test_FindConstByType_AstFile, err = parser.ParseFile(
		token.NewFileSet(),
		Test_FindConstByType_FileName,
		nil,
		parser.ParseComments)
	if err != nil {
		tb.Fatal(err)
	}
}

func FindConstByTypeNameCase(tb testing.TB) {
	cvs := FindConstByTypeName(
		Test_FindConstByType_AstFile,
		Test_FindConstByType_TypeName)

	names := []string{"TypeCA", "TypeCB", "TypeCC", "TypeCD", "TypeCE", "TypeCF"}
	for i, cv := range cvs {
		if names[i] != cv.Names[0].Name {
			log.Fatalf("name not match:\nexcept:%s\tnow:%s", names[i], cv.Names[0].Name)
		}
	}

}

func TestFindConstByTypeNameCaseAll(t *testing.T) {
	initFindConstByTypeName(t)
	FindConstByTypeNameCase(t)
}

func BenchmarkFindConstByTypeName(b *testing.B) {
	initFindConstByTypeName(b)
	for i := 0; i < b.N; i++ {
		FindConstByTypeNameCase(b)
	}
	b.ReportAllocs()
}
