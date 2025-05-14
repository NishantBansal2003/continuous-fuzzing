package fuzz

import fuzzing "fuzzing/fuzzing"

// ! Need to specify storage repo to fetch corpus as this is mutated
// ! Looks like nothing is need to done, I ran the CI again and it fteches the
// ! corpus from previous run.
func Fuzz(data []byte) int {
	v := string(data)

	// As per the documentation, this isn't always equal to v, so it makes
	// no sense to check for equality. It can still be interesting to find
	// panics in it though.
	err := fuzzing.ParseString(v)

	if err != nil {
		// if err.Error() == "Invalid string" {
		// 	return 0
		// }
		panic(err)
	}

	return 0
}
