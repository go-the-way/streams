package constraint

type (
	Int interface {
		int8 | int16 | int32 | int | int64
	}
	Uint interface {
		uint8 | uint16 | uint32 | uint | uint64
	}
	Float  interface{ float32 | float64 }
	Number interface{ Int | Uint | Float }
)
