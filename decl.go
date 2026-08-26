package gochecksumtype

import (
	"go/ast"
	"go/token"
	"strings"
)

// sumTypeDecl is a declaration of a sum type in a Go source file.
type sumTypeDecl struct {
	// The type named by this decl.
	TypeName string
	// Position where the declaration was found.
	Pos token.Pos
}

// findSumTypeDecls searches the given AST files for sum type declarations of
// the form `//sumtype:decl`.
func findSumTypeDecls(files []*ast.File) ([]sumTypeDecl, error) {
	var decls []sumTypeDecl
	var retErr error
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			if node == nil {
				return true
			}
			decl, ok := node.(*ast.GenDecl)
			if !ok || decl.Doc == nil {
				return true
			}
			var tspec *ast.TypeSpec
			for _, spec := range decl.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				tspec = ts
			}
			for _, line := range decl.Doc.List {
				if !strings.HasPrefix(line.Text, "//sumtype:decl") {
					continue
				}
				if tspec == nil {
					retErr = notFoundError{Decl: sumTypeDecl{Pos: decl.Pos(), TypeName: ""}}
					return false
				}
				pos := tspec.Pos()
				decl := sumTypeDecl{TypeName: tspec.Name.Name, Pos: pos}
				debugf("found sum type decl: %s", decl.TypeName)
				decls = append(decls, decl)
				break
			}
			return true
		})
	}
	return decls, retErr
}
