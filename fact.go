package gochecksumtype

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

// sumTypeFact is a package-level fact that declares a sum type and its variants.
// It is exported by a package that contains a //sumtype:decl annotation and
// imported by downstream packages to enable cross-package exhaustiveness checking.
type sumTypeFact struct {
	TypeName string   // name of the sum type interface (e.g. "MySumType")
	Variants []string // names of concrete types implementing the interface
}

var _ analysis.Fact = (*sumTypeFact)(nil)

// AFact implements [analysis.Fact].
func (*sumTypeFact) AFact() {}

func (f *sumTypeFact) String() string {
	return fmt.Sprintf("sumTypeFact{%s %v}", f.TypeName, f.Variants)
}
