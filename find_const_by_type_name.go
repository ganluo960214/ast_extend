package ast_extend

import (
	"go/ast"
	"go/token"
)

// FindConstByTypeName find all const in *ast.File
//
// Parameter
// TypeName : find the const value
//
// Return
// []*ast.ValueSpec : return all found const value
//
// Example
// type Type byte
//
// const (
//   TypeCA Type = iota
//   TypeCB
//   NotTypeCA = 2
//   NotTypeCB = iota
//   TypeCC    = Type(iota)
//   TypeCD
//   NotTypeCC      = int(iota)
//   NotTypeCD      = byte(7)
//   TypeCE    Type = 8
//   TypeCF    Type = 9
// )
// */
func FindConstByTypeName(f *ast.File, typeName string) []*ast.ValueSpec {
	constVS := []*ast.ValueSpec{}

	for _, d := range f.Decls {
		gd, ok := d.(*ast.GenDecl)
		if ok == false || gd.Tok != token.CONST {
			continue
		}

		constType := ""
		for _, s := range gd.Specs {
			vs, ok := s.(*ast.ValueSpec)
			if ok == false {
				continue
			}

			// make sure type
			if vs.Type == nil {
				if len(vs.Values) == 0 {
					if constType != typeName {
						continue
					}
				} else {
					constType = ""
					ce, ok := vs.Values[0].(*ast.CallExpr)
					if ok == false {
						continue
					}
					i, ok := ce.Fun.(*ast.Ident)
					if ok == false {
						continue
					}
					constType = i.Name
					if constType != typeName {
						continue
					}
				}
			} else {
				i, ok := vs.Type.(*ast.Ident)
				if ok == false {
					continue
				}

				constType = i.Name
				if constType != typeName {
					continue
				}
			}

			constVS = append(constVS, vs)
		}
	}

	return constVS
}
