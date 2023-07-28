package mytypes_test

import (
	"fmt"
	"mytypes"
	"strings"
	"testing"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18)
	got := input.Twice()
	if want != got {
		t.Errorf("Twice: %d , want:%d, got:%d", input, want, got)
	}
}

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	sampleText := mytypes.MyString("Hello, Gophers!")
	want := 15
	got := sampleText.Len()
	if want != got {
		t.Errorf("String: %q - want length:%d, got:%d", sampleText, want, got)
	}
}

func TestStringBuilder(t *testing.T) {
	t.Parallel()
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("Gophers!")
	want := "Hello, Gophers !"
	got := sb.String()
	if want != got {
		t.Errorf("want:%q, got:%q", want, got)
	}
	wantLength := 14
	gotLength := sb.Len()
	if wantLength != gotLength {
		t.Errorf("wantLength:%d, gotLength:%d", wantLength, gotLength)
	}

}

// func TestMyBuilder(t *testing.T) {
// 	t.Parallel()
// 	var mb mytypes.MyBuilder
// 	want := 15
// 	got := mb.Len()
// 	if want != got {
// 		t.Errorf("want:%d, got:%d", want, got)
// 	}
// }

func TestHello(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilder
	want := "Hello, Gophers!"
	got := mb.Hello()
	if want != got {
		t.Errorf("want:%q, got:%q", want, got)
	}

}

func TestMyBuilderStruct(t *testing.T) {
	t.Parallel()
	var mb mytypes.MyBuilderStruct
	mb.Contents.WriteString("Hello, ")
	mb.Contents.WriteString("Gophers!")
	want := "Hello, Gophers!"
	got := mb.Contents.String()
	if want != got {
		t.Errorf("want:%q, got:%q", want, got)
	}
	wantLen := 15
	gotLen := mb.Contents.Len()
	if wantLen != gotLen {
		t.Errorf("wantLen:%d, gotLen:%d", wantLen, gotLen)
	}
}

func TestStringUppercaser(t *testing.T) {
	t.Parallel()
	var su mytypes.MyBuilderStruct
	su.Contents.WriteString("Hello, Gophers! ")
	want := "HELLO, GOPHERS!"
	got := su.ToUpper()
	fmt.Println(su.Contents.Len())
	if want != got {
		t.Errorf("want:%q, got:%q", want, got)
	}
}

func TestDouble(t *testing.T) {
	t.Parallel()
	/*var x int = 12
	want := 24
	mytypes.Double(&x)*/
	x := mytypes.MyInt(12)
	want := mytypes.MyInt(24)
	p := &x
	p.Double()
	if want != x {
		t.Errorf("want:%d, got:%d", want, x)
	}
}
