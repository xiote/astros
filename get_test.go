package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	cases := []struct {
		inUrl, wantContents string
	}{
		{"http://ipcheck.re4u.co.kr", "112.223.47.109"},
	}
	for _, c := range cases {
		gotContents := string(Get(c.inUrl))
		if gotContents != c.wantContents {
			t.Errorf("Get(%q) == %q, want %q", c.inUrl, gotContents, c.wantContents)
		}
	}
}
