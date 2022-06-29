package main

import (
	"strings"
	"testing"

	"gopkg.in/olebedev/go-duktape.v3"
)

func FuzzDuktape(f *testing.F) {
	testcases := []string{"print(1)"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		ctx := duktape.New()
		ctx.PevalString(`var console = {log:print,warn:print,error:print,info:print}`)
		if err := ctx.PevalString(string(orig)); err != nil {
			if !strings.HasPrefix(err.Error(), "ReferenceError") {
				if !strings.HasPrefix(err.Error(), "Compilation failed") {
					if !strings.HasPrefix(err.Error(), "SyntaxError") {
						if !strings.HasPrefix(err.Error(), "TypeError") {
							if !strings.HasPrefix(err.Error(), "RangeError") {
								t.Errorf(err.Error())
							}
						}
					}
				}
			}
		}

		// rev := Reverse(orig)
		// doubleRev := Reverse(rev)
		// if orig != doubleRev {
		// 	t.Errorf("Before: %q, after: %q", orig, doubleRev)
		// }
		// if utf8.ValidString(orig) && !utf8.ValidString(rev) {
		// 	t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		// }
	})
}
