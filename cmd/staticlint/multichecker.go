package main

import (
	"github.com/gordonklaus/ineffassign/pkg/ineffassign"
	"github.com/kisielk/errcheck/errcheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"

	"urlshortener/internal/analyzer"
)

func main() {

	var mychecks []*analysis.Analyzer

	for _, v := range staticcheck.Analyzers {
		mychecks = append(mychecks, v.Analyzer)
	}

	for _, v := range stylecheck.Analyzers {
		if v.Analyzer.Name == "ST1000" {
			mychecks = append(mychecks, v.Analyzer)
		}
	}

	mychecks = append(mychecks,
		printf.Analyzer,
		shift.Analyzer,
		shadow.Analyzer,
		analyzer.ExitInMainAnalyzer,
		errcheck.Analyzer,
		ineffassign.Analyzer,
	)

	multichecker.Main(
		mychecks...,
	)
}
