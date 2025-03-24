package fuzzing

import (
	"testing"
)

func FuzzParseString(f *testing.F) {
	// f.Add("1, Name, 23rd Fifth Street")
	// f.Add("2, Name2, 1 Main street")
	//! corpusDir = "testdata/fuzz" this is found in fuzz std directory
	//! so the fuzz corpus is feed from f.Add and corpusDir
	f.Fuzz(func(t *testing.T, s string) {
		err := ParseString(s)
		if err != nil {
			// Added after the main code logic is changed
			if err.Error() == "Invalid string" {
				return
			}
			t.Errorf("%v", err)
		}
	})
}
