package gochecksumtype

import (
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

// checks exhaustiveness of sum type
// switch statements. Sum types are declared with a //sumtype:decl comment
// above a sealed interface type.
func newAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:             "sumtype",
		Doc:              "check exhaustiveness of sum type switch statements",
		Run:              run,
		Flags:            newFlags(),
		FactTypes:        []analysis.Fact{new(sumTypeFact)},
		URL:              "",
		RunDespiteErrors: false,
		Requires:         nil,
		ResultType:       nil,
	}
}

var Analyzer = newAnalyzer()

func run(pass *analysis.Pass) (any, error) {
	decls, err := findSumTypeDecls(pass.Files)
	if err != nil {
		return nil, err
	}

	defs, defErrs := findSumTypeDefs(pass, decls)
	for _, e := range defErrs {
		pass.Reportf(e.(Error).Pos(), "%s", e.Error())
	}

	// Export facts so downstream packages can check exhaustiveness
	// against sum types defined here.
	for _, def := range defs {
		variantNames := make([]string, len(def.Variants))
		for i, v := range def.Variants {
			variantNames[i] = v.Name()
		}
		pass.ExportPackageFact(&sumTypeFact{
			TypeName: def.Decl.TypeName,
			Variants: variantNames,
		})
	}

	// Import sum type facts from dependencies.
	for _, pkg := range pass.Pkg.Imports() {
		var fact sumTypeFact
		if pass.ImportPackageFact(pkg, &fact) {
			if def := factToTypeDef(pkg, &fact); def != nil {
				defs = append(defs, *def)
			}
		}
	}

	cfg := cfgFromFlags(pass.Analyzer.Flags)

	// Check exhaustiveness for all type switches in this package.
	for _, astfile := range pass.Files {
		for _, err := range checkFile(pass, astfile, defs, cfg) {
			pass.Reportf(err.(Error).Pos(), "%s", err.Error())
		}
	}

	return nil, nil
}

// factToTypeDef reconstructs a [sumTypeDef] from an imported fact.
func factToTypeDef(pkg *types.Package, fact *sumTypeFact) *sumTypeDef {
	obj := pkg.Scope().Lookup(fact.TypeName)
	if obj == nil {
		return nil
	}
	iface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		return nil
	}
	def := &sumTypeDef{
		Decl:     sumTypeDecl{TypeName: fact.TypeName, Pos: token.NoPos},
		Ty:       iface,
		Variants: make([]types.Object, 0, len(fact.Variants)),
	}
	for _, name := range fact.Variants {
		vObj := pkg.Scope().Lookup(name)
		if vObj != nil {
			def.Variants = append(def.Variants, vObj)
		}
	}
	return def
}
