package gochecksumtype

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMissingOne(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "missing_one")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestMissingTwo(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "missing_two")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestMissingWithPanic(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "missing_with_panic")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestNoMissing(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "no_missing")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestDefaultExhaustive(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "default_exhaustive")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestNotSealed(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "not_sealed")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestNotInterface(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "not_interface")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestSharedInterface(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "shared_interface")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}

func TestAllLeaves(t *testing.T) {
	testdata := analysistest.TestData()
	results := analysistest.Run(t, testdata, Analyzer, "all_leaves")
	if len(results) == 0 {
		t.Fatal("no results")
	}
	for _, r := range results {
		if r.Err != nil {
			t.Errorf("unexpected error: %v", r.Err)
		}
	}
}
