package main

import (
	"strings"
	"testing"
)

func TestSplitField(t *testing.T) {
	cases := []struct {
		source string
		expect []string
	}{
		{source: `foo bar`, expect: []string{`foo`, `bar`}},
		{source: `foo  " bar "`, expect: []string{`foo`, ` bar `}},
		{source: `ahaha  " "`, expect: []string{`ahaha`, ` `}},
	}

	for _, case1 := range cases {
		result := splitField(case1.source)
		if len(result) != len(case1.expect) {
			t.Fatalf("expect `%s`, but `%s`",
				strings.Join(case1.expect, "`,`"), strings.Join(result, "`,`"))
		}
		for i, r := range result {
			if r != case1.expect[i] {
				t.Fatalf("expect `%s`, but `%s`",
					strings.Join(case1.expect, "`,`"), strings.Join(result, "`,`"))
			}
		}
	}
}
