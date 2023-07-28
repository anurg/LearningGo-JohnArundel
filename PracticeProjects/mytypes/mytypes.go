package mytypes

import "strings"

type MyInt int
type MyString string
type MyBuilder strings.Builder
type MyBuilderStruct struct {
	Contents strings.Builder
}
type StringUppercaser struct {
	Contents strings.Builder
}

func (i MyInt) Twice() MyInt {
	return 2 * i
}

func (s MyString) Len() int {
	return len(s)
}

func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}

func (mb MyBuilderStruct) ToUpper() string {
	return strings.ToUpper(mb.Contents.String())
}

func Double(x *int) {
	*x *= 2
}

func (x *MyInt) Double() {
	*x *= 2
}
