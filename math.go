package math

import "math"

type (
	UnsignedInteger interface {
		uint | uint8 | uint16 | uint32 | uint64
	}

	SignedInteger interface {
		int | int8 | int16 | int32 | int64
	}

	Integer interface {
		SignedInteger | UnsignedInteger
	}

	Float interface {
		float32 | float64
	}

	Number interface {
		Float | Integer
	}
)

// Div divides a by b, returning the resulting value as the specified QuotientT type.
func Div[QuotientT, DividendT, DivisorT Number](a DividendT, b DivisorT) QuotientT {
	return QuotientT(a) / QuotientT(b)
}

// Remainder returns the remainder from dividing a by b, as the specified RemainderT type.
func Remainder[RemainderT, DividendT, DivisorT Number](a DividendT, b DivisorT) RemainderT {
	return RemainderT(math.Remainder(float64(a), float64(b)))
}
