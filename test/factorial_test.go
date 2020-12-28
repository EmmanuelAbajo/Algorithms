package test

import (
	"fmt"
	"testing"
	nb "algorithms/main/numbers"
)

type testCase struct {
	value int
	expected int
}

func TestFactorial(t *testing.T) {

	testCases := []testCase{
		{1,1},
		{0,1},
		{5,120},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test factorial: %d", tc.value), func (t *testing.T) {
			out := nb.Factorial(tc.value)

			if (out != tc.expected) {
				t.Fatalf("%d != %d",out, tc.expected)
			}
		})
	}

}