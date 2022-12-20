package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		str  string
		want bool
	}{
		{"google.com", true},
		{"http://google.com", true},
		{"https://google.com", true},
		{"www.google.com", true},
		{"htp://google.com",false},
	}

	for _, tt := range tests {
		// t.Run enables running "subtests", one for each
		// table entry. These are shown separately
		// when executing `go test -v`.
		testname := fmt.Sprintf("URL Candidate:%s",tt.str)
		t.Run(testname, func(t *testing.T) {
			ans := isArgWebsite(tt.str)
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
