# Ast Extend

easy use *ast.File.

# Catalog

## [Function](#Function)
### [AstExtFindFuncByName](#AstExtFindFuncByName)




## Function
### AstExtFindFuncByName
```
AstExtFindFuncByName(f *ast.File, ReceiverType, FuncName string) (*ast.FuncDecl, bool)

Parameter
ReceiverType : method pass the receiver type if there is a receiver or empty string
FuncName     : try to find the function name


Return
*ast.FuncDecl : return the found the function or nil
bool          : return true if found function or false
```

Example Function Without Receiver 
```go
func F() {}
func AstExtFindFuncByName(f,"","F")
```

Example Function With Receiver
```go
type idea struct{}
func (idea) F(){}
func (*idea) F(){}
func (i idea) F(){}
func (i *idea) F(){}
func AstExtFindFuncByName(f,"idea","F")
```
