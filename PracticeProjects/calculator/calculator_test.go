package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 5, b: 3, want: 8},
		{a: 5, b: 13, want: 18},
		{a: 25, b: 3, want: 28},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want:%f, got:%f", tc.want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 5, b: 3, want: 2},
		{a: 5, b: 13, want: 8},
		{a: 25, b: 3, want: 22},
	}
	// using range for loop
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want:%f, got:%f", tc.want, got)
		}
	}

	// Using for loop
	// for i := 0; i < len(testCases); i++ {
	// 	got := calculator.Subtract(testCases[i].a, testCases[i].b)
	// 	if testCases[i].want != got {
	// 		t.Errorf("want:%f, got:%f", testCases[i].want, got)
	// 	}
	// }
	// Using for range loop without current element & using indexes
	// for i := range testCases {
	// 	got := calculator.Subtract(testCases[i].a, testCases[i].b)
	// 	if testCases[i].want != got {
	// 		t.Errorf("want:%f, got:%f", testCases[i].want, got)
	// 	}
	// }
}

//TODO - Multiplication

// Test Division

func closeEnough(a, b float64, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 1, b: 3, want: 0.333},
		{a: 0, b: 13, want: 0},
		{a: 25, b: 5, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf(" want no error input but got error %v", err)
		}
		// if tc.want != got {
		// 	t.Errorf("want:%f, got:%f", tc.want, got)
		// }
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("want:%f, got:%f , diff %f", tc.want, got, tc.want-got)
		}

	}
}

func TestDivideInvalidInput(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got none")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 16, want: 4},
		{a: 25, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("Want error to be nil but got error %v", err)
		}
		if tc.want != got {
			t.Errorf("want:%f, got:%f", tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	var a float64 = -4
	_, err := calculator.Sqrt(a)
	if err == nil {
		t.Error("expected eror but got no error.")
	}
}
