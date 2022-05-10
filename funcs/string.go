package funcs

import "strconv"

var (
	String2Bool = func(str string) (b bool) { b, _ = strconv.ParseBool(str); return }

	String2Int   = func(str string) int { i, _ := strconv.ParseInt(str, 10, 32); return int(i) }
	String2Byte  = func(str string) byte { i, _ := strconv.ParseInt(str, 10, 8); return byte(i) }
	String2Int8  = func(str string) int8 { i, _ := strconv.ParseInt(str, 10, 8); return int8(i) }
	String2Int16 = func(str string) int16 { i, _ := strconv.ParseInt(str, 10, 16); return int16(i) }
	String2Int32 = func(str string) int32 { i, _ := strconv.ParseInt(str, 10, 32); return int32(i) }
	String2Int64 = func(str string) (i int64) { i, _ = strconv.ParseInt(str, 10, 64); return }

	String2UInt   = func(str string) uint { i, _ := strconv.ParseUint(str, 10, 32); return uint(i) }
	String2UInt8  = func(str string) uint8 { i, _ := strconv.ParseUint(str, 10, 8); return uint8(i) }
	String2UInt16 = func(str string) uint16 { i, _ := strconv.ParseUint(str, 10, 16); return uint16(i) }
	String2UInt32 = func(str string) uint32 { i, _ := strconv.ParseUint(str, 10, 32); return uint32(i) }
	String2UInt64 = func(str string) (i uint64) { i, _ = strconv.ParseUint(str, 10, 64); return }

	String2Float32 = func(str string) float32 { f, _ := strconv.ParseFloat(str, 32); return float32(f) }
	String2Float64 = func(str string) (f float64) { f, _ = strconv.ParseFloat(str, 64); return }
)
