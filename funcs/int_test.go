package funcs

import "testing"

func TestInt(t *testing.T) {
	t.Log(Int2BinaryString(16))
	t.Log(Int2OctalString(16))
	t.Log(Int2String(16))
	t.Log(Int2HexString(16))
}
