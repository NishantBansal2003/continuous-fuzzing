package fuzz

import "fuzzing/fuzzing"

func Fuzz(data []byte) int {
	v := string(data)

	// As per the documentation, this isn't always equal to v, so it makes
	// no sense to check for equality. It can still be interesting to find
	// panics in it though.
	err := fuzzing.ParseString(v)

	if err != nil {
		panic(err)
	}

	return 0
}
