package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	// src is the input for which we want to print the AST.
	src := `
						package main
						func main() {
								println("Hello, World!")
						}
`

	// Create the AST by parsing src.
	// A FileSet represents a set of source files. Methods of file sets are synchronized; multiple goroutines may invoke them concurrently.
	// 多个文件的集合?
	fset := token.NewFileSet() // positions are relative to fset
	// func ParseFile(fset *token.FileSet, filename string, src interface{}, mode Mode) (f *ast.File, err error)
	// type Mode uint
	// const (
	// 	PackageClauseOnly Mode             = 1 << iota // stop parsing after package clause
	// 	ImportsOnly                                    // stop parsing after import declarations
	// 	ParseComments                                  // parse comments and add them to AST
	// 	Trace                                          // print a trace of parsed productions
	// 	DeclarationErrors                              // report declaration errors
	// 	SpuriousErrors                                 // same as AllErrors, for backward-compatibility
	// 	AllErrors         = SpuriousErrors             // report all errors (not just the first 10 on different lines)
	// )
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)
}
