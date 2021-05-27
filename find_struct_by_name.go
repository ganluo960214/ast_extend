package ast_extend

import (
	"go/ast"
	"go/token"
)

// FindStructByName find struct in *ast.File
//
// Parameter
// StructName: try to find the struct name
//
// Return
// *ast.TypeSpec  : struct declaration
// *ast.StructType: struct content declaration
// bool           : is found struct
//
// Example
// func F() {}
// FindStructByName(f,"F")
//
func FindStructByName(f *ast.File, StructName string) (*ast.TypeSpec, *ast.StructType, bool) {
	for _, d := range f.Decls {
		gd, ok := d.(*ast.GenDecl)
		if ok == false || gd.Tok != token.TYPE {
			continue
		}

		for _, s := range gd.Specs {
			ts, ok := s.(*ast.TypeSpec)
			if ok == false {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if ok == false && ts.Name.Name != StructName {
				continue
			}

			return ts, st, true
		}
	}
	return nil, nil, false
}
