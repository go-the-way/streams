package funcs

import "fmt"

var (
	// Int2BinaryString
	// 10 -> 1010
	Int2BinaryString = func(in int) string { return fmt.Sprintf("%b", in) }
	// Int2OctalString
	// 10 -> 12
	Int2OctalString = func(in int) string { return fmt.Sprintf("%o", in) }
	// Int2String
	// 10 -> 10
	Int2String = func(in int) string { return fmt.Sprintf("%d", in) }
	// Int2HexString
	// 15 -> f
	Int2HexString = func(in int) string { return fmt.Sprintf("%x", in) }
)
