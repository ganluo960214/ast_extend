package ast_extend

import "go/ast"

// FindFuncByName find function in *ast.File
//
// Parameter
// ReceiverType: pass the receiver type if there is a receiver or empty string
// FuncName    : try to find the function name
//
// Return
// *ast.FuncDecl: return the found the function or nil
// bool         : return true if found function or false
//
// Example
// function without receiver:
// func F() {}
// FindFuncByName(f,"","F")
// function with receiver:
// type idea struct{}
// func (idea) F(){}
// FindFuncByName(f,"idea","F")
// */
func FindFuncByName(f *ast.File, ReceiverType, FuncName string) (*ast.FuncDecl, bool) {
	for _, d := range f.Decls {
		d, ok := d.(*ast.FuncDecl)
		if ok == false {
			continue
		}

		if (ReceiverType == "" && d.Recv != nil) ||
			(ReceiverType != "" && d.Recv == nil) {
			continue
		}

		if d.Name.Name != FuncName {
			continue
		}

		if ReceiverType == "" {
			return d, true
		} else {
			switch d.Recv.List[0].Type.(type) {
			case *ast.StarExpr:
				e := d.Recv.List[0].Type.(*ast.StarExpr)
				i := e.X.(*ast.Ident)
				if i.Name == ReceiverType {
					return d, true
				}
			case *ast.Ident:
				i := d.Recv.List[0].Type.(*ast.Ident)
				if i.Name == ReceiverType {
					return d, true
				}
			}
		}

	}

	return nil, false
}
