package fuzzing

import (
	"errors"
	"strconv"
	"strings"
)

func isInt(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}

	return len(s) <= 10
}

func ParseString(s string) error {
	parts := strings.Split(s, ",")

	// This code I got to know from the fuzz-tests
	// Now I need to update the fuzz tests else the fuzz tests won't run
	// as expected.(This is not the job of continuous fuzzing, this is devs job)
	if len(parts) < 3 {
		return errors.New("Invalid string")
	}
	if isInt(parts[0]) {
		_, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
	}
	parts[2] += ", Senegal"
	return nil
}
