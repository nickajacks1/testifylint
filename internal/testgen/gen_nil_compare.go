package main

import (
	"strings"
	"text/template"

	"github.com/Antonboom/testifylint/internal/checkers"
)

type NilCompareTestsGenerator struct{}

func (NilCompareTestsGenerator) Checker() checkers.Checker {
	return checkers.NewNilCompare()
}

func (g NilCompareTestsGenerator) TemplateData() any {
	var (
		checker = g.Checker().Name()
		report  = checker + ": use %s.%s"
	)

	return struct {
		CheckerName       CheckerName
		InvalidAssertions []Assertion
		ValidAssertions   []Assertion
		IgnoredAssertions []Assertion
	}{
		CheckerName: CheckerName(checker),
		InvalidAssertions: []Assertion{
			{Fn: "Equal", Argsf: "value, nil", ReportMsgf: report, ProposedFn: "Nil", ProposedArgsf: "value"},
			{Fn: "Equal", Argsf: "nil, value", ReportMsgf: report, ProposedFn: "Nil", ProposedArgsf: "value"},
			{Fn: "NotEqual", Argsf: "value, nil", ReportMsgf: report, ProposedFn: "NotNil", ProposedArgsf: "value"},
			{Fn: "NotEqual", Argsf: "nil, value", ReportMsgf: report, ProposedFn: "NotNil", ProposedArgsf: "value"},
		},
		ValidAssertions: []Assertion{
			{Fn: "Nil", Argsf: "value"},
			{Fn: "NotNil", Argsf: "value"},
		},
		IgnoredAssertions: []Assertion{
			{Fn: "Equal", Argsf: "value, value"},
			{Fn: "Equal", Argsf: "nil, nil"},
			{Fn: "NotEqual", Argsf: "value, value"},
			{Fn: "NotEqual", Argsf: "nil, nil"},
		},
	}
}

func (NilCompareTestsGenerator) ErroredTemplate() Executor {
	return template.Must(template.New("NilCompareTestsGenerator.ErroredTemplate").
		Funcs(fm).
		Parse(nilCompareTestTmpl))
}

func (NilCompareTestsGenerator) GoldenTemplate() Executor {
	return template.Must(template.New("NilCompareTestsGenerator.GoldenTemplate").
		Funcs(fm).
		Parse(strings.ReplaceAll(nilCompareTestTmpl, "NewAssertionExpander", "NewAssertionExpander.AsGolden")))
}

const nilCompareTestTmpl = header + `

package {{ .CheckerName.AsPkgName }}

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func {{ .CheckerName.AsTestName }}(t *testing.T) {
	var value any

	// Invalid.
	{
		{{- range $ai, $assrn := $.InvalidAssertions }}
			{{ NewAssertionExpander.Expand $assrn "assert" "t" nil }}
			{{ NewAssertionExpander.Expand $assrn "require" "t" nil }}
		{{- end }}
	}

	// Valid.
	{
		{{- range $ai, $assrn := $.ValidAssertions }}
			{{ NewAssertionExpander.Expand $assrn "assert" "t" nil }}
			{{ NewAssertionExpander.Expand $assrn "require" "t" nil }}
		{{- end }}
	}

	// Ignored.
	{
		{{- range $ai, $assrn := $.IgnoredAssertions }}
			{{ NewAssertionExpander.Expand $assrn "assert" "t" nil }}
			{{ NewAssertionExpander.Expand $assrn "require" "t" nil }}
		{{- end }}
	}
}
`
