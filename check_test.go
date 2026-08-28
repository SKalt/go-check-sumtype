package gochecksumtype

import (
	"errors"
	"flag"
	"fmt"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func (c cfg) setFlags(fs *flag.FlagSet) error {
	return errors.Join(
		fs.Set(flagDefaultSignifiesExhaustive, fmt.Sprint(c.defaultSignifiesExhaustive())),
		fs.Set(flagIncludeSharedInterfaces, fmt.Sprint(c.includeSharedInterfaces())),
	)
}

func (c cfg) String() string {
	return fmt.Sprintf("%s:%t %s:%t",
		flagDefaultSignifiesExhaustive, c.defaultSignifiesExhaustive(),
		flagIncludeSharedInterfaces, c.includeSharedInterfaces(),
	)
}

var testdata = analysistest.TestData()

func TestExpectFindings(t *testing.T) {
	type testCase struct {
		name string
		cfg  cfg
	}
	for _, tc := range []testCase{
		// tests that we detect a single missing variant.
		// see ./testdata/src/missing_one/main.go
		{name: "missing_one", cfg: defaultExhaustive},
		// tests that we detect two missing variants.
		// see ./testdata/src/missing_two/main.go
		{name: "missing_two", cfg: defaultExhaustive},
		// tests that we detect a single missing variant even
		// if we have a trivial default case that panics.
		// see ./testdata/src/missing_with_panic/main.go
		{name: "missing_with_panic", cfg: defaultExhaustive},
		// tests that we correctly detect exhaustive case analysis.
		// see ./testdata/src/no_missing/main.go
		{name: "no_missing", cfg: defaultExhaustive},
		// tests that even if we have a missing variant, a default
		// case should thwart exhaustiveness checking when [cfg.defaultSignifiesExhaustive] is true.
		// see ./testdata/src/default_exhaustive/main.go
		{name: "default_exhaustive", cfg: defaultExhaustive},
		// tests that even if we have a missing variant, a default
		// case should thwart exhaustiveness checking when [cfg.defaultSignifiesExhaustive] is false.
		{name: "not_default_exhaustive", cfg: 0},
		// tests that we report an error if one tries to declare a sum
		// type with an unsealed interface. See ./testdata/src/not_sealed/main.go
		{name: "not_sealed", cfg: defaultExhaustive},
		// tests that we report an error if one tries to declare a sum
		// type that doesn't correspond to an interface. See ./testdata/src/not_interface/main.go
		{name: "not_interface", cfg: defaultExhaustive},
		// tests that if a shared interface is declared in the switch
		// statement, we don't report an error if structs that implement the interface are not explicitly
		// declared in the switch statement.
		// see ./testdata/src/shared_interface
		{name: "subtype_in_switch", cfg: includeSharedInterfaces},
		// tests that we do not report an error if a switch statement
		// covers all leaves of the sum type, even if any SubTypes are not explicitly covered
		// see ./testdata/src/all_leaves/main.go
		{name: "all_leaves", cfg: defaultExhaustive},
		// tests that `_test.go` files are covered.
		{name: "with_tests", cfg: defaultExhaustive},
		// tests that aliases are not counted as additional variants.
		{name: "with_alias", cfg: defaultExhaustive},
		// tests that all sum types declared in an imported package are checked.
		{name: "multiple_sumtypes_user", cfg: defaultExhaustive},
	} {
		t.Run(tc.name, func(t *testing.T) {
			analyzer := newAnalyzer()
			err := tc.cfg.setFlags(&analyzer.Flags)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(cfgFromFlags(analyzer.Flags).String())
			results := analysistest.Run(t, testdata, analyzer, tc.name)
			if len(results) == 0 {
				t.Fatalf("0 results for patterns %q", tc.name)
			}
			for _, r := range results {
				if r.Err != nil {
					t.Errorf("unexpected error: %v", r.Err)
				}
			}
		})
	}
}
