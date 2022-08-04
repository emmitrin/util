package collections

import (
	"strconv"
	"testing"
)

func TestConvert(t *testing.T) {
	// easily convert slices of values
	numberStrings := []string{"10", "12", "14", "16"}

	numbers, err := Convert(numberStrings, strconv.Atoi)
	if err != nil {
		t.Error(err)
	}

	if !Compare(numbers, []int{10, 12, 14, 16}) {
		t.Error("comparison error")
	}

	badNumberStrings := []string{"aaaa", "5", "0"}
	numbers, err = Convert(badNumberStrings, strconv.Atoi)

	if err == nil {
		t.Error("an error was expected")
	}

}
