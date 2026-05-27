//go:build go1.26 && goexperiment.simd && amd64

package rosimd

import (
	"simd/archsimd"

	"github.com/samber/ro"
)

// AVX ScalarTo/ToScalar variants (128-bit vectors)

// Int8x16 variants

// ScalarToInt8x16 converts a stream of scalar int8 values into int8x16 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 16 elements. When the buffer
// is full (exactly 16 values), a complete Int8x16 vector is emitted. On source completion,
// any remaining values in the buffer (less than 16) are emitted as a partial vector.
func ScalarToInt8x16[T ~int8]() func(ro.Observable[T]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// Int8x16ToScalar converts a stream of int8x16 vectors into scalar int8 values.
//
// Buffer semantics: Each Int8x16 vector is fully unpacked into 16 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 16 elements set) will emit zero values for unused lanes.
func Int8x16ToScalar[T ~int8]() func(ro.Observable[*archsimd.Int8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int16x8 variants

// ScalarToInt16x8 converts a stream of scalar int16 values into int16x8 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 8 elements. When the buffer
// is full (exactly 8 values), a complete Int16x8 vector is emitted. On source completion,
// any remaining values in the buffer (less than 8) are emitted as a partial vector.
func ScalarToInt16x8[T ~int16]() func(ro.Observable[T]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// Int16x8ToScalar converts a stream of int16x8 vectors into scalar int16 values.
//
// Buffer semantics: Each Int16x8 vector is fully unpacked into 8 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 8 elements set) will emit zero values for unused lanes.
func Int16x8ToScalar[T ~int16]() func(ro.Observable[*archsimd.Int16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int32x4 variants

// ScalarToInt32x4 converts a stream of scalar int32 values into int32x4 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 4 elements. When the buffer
// is full (exactly 4 values), a complete Int32x4 vector is emitted. On source completion,
// any remaining values in the buffer (less than 4) are emitted as a partial vector.
func ScalarToInt32x4[T ~int32]() func(ro.Observable[T]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// Int32x4ToScalar converts a stream of int32x4 vectors into scalar int32 values.
//
// Buffer semantics: Each Int32x4 vector is fully unpacked into 4 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 4 elements set) will emit zero values for unused lanes.
func Int32x4ToScalar[T ~int32]() func(ro.Observable[*archsimd.Int32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int64x2 variants

// ScalarToInt64x2 converts a stream of scalar int64 values into int64x2 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 2 elements. When the buffer
// is full (exactly 2 values), a complete Int64x2 vector is emitted. On source completion,
// any remaining values in the buffer (less than 2) are emitted as a partial vector.
func ScalarToInt64x2[T ~int64]() func(ro.Observable[T]) ro.Observable[*archsimd.Int64x2] {
	_ = "STUB: not implemented"
	return nil
}

// Int64x2ToScalar converts a stream of int64x2 vectors into scalar int64 values.
//
// Buffer semantics: Each Int64x2 vector is fully unpacked into 2 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 2 elements set) will emit zero values for unused lanes.
func Int64x2ToScalar[T ~int64]() func(ro.Observable[*archsimd.Int64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint8x16 variants

// ScalarToUint8x16 converts a stream of scalar uint8 values into uint8x16 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 16 elements. When the buffer
// is full (exactly 16 values), a complete Uint8x16 vector is emitted. On source completion,
// any remaining values in the buffer (less than 16) are emitted as a partial vector.
func ScalarToUint8x16[T ~uint8]() func(ro.Observable[T]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// Uint8x16ToScalar converts a stream of uint8x16 vectors into scalar uint8 values.
//
// Buffer semantics: Each Uint8x16 vector is fully unpacked into 16 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 16 elements set) will emit zero values for unused lanes.
func Uint8x16ToScalar[T ~uint8]() func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint16x8 variants

// ScalarToUint16x8 converts a stream of scalar uint16 values into uint16x8 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 8 elements. When the buffer
// is full (exactly 8 values), a complete Uint16x8 vector is emitted. On source completion,
// any remaining values in the buffer (less than 8) are emitted as a partial vector.
func ScalarToUint16x8[T ~uint16]() func(ro.Observable[T]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// Uint16x8ToScalar converts a stream of uint16x8 vectors into scalar uint16 values.
//
// Buffer semantics: Each Uint16x8 vector is fully unpacked into 8 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 8 elements set) will emit zero values for unused lanes.
func Uint16x8ToScalar[T ~uint16]() func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint32x4 variants

// ScalarToUint32x4 converts a stream of scalar uint32 values into uint32x4 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 4 elements. When the buffer
// is full (exactly 4 values), a complete Uint32x4 vector is emitted. On source completion,
// any remaining values in the buffer (less than 4) are emitted as a partial vector.
func ScalarToUint32x4[T ~uint32]() func(ro.Observable[T]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// Uint32x4ToScalar converts a stream of uint32x4 vectors into scalar uint32 values.
//
// Buffer semantics: Each Uint32x4 vector is fully unpacked into 4 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 4 elements set) will emit zero values for unused lanes.
func Uint32x4ToScalar[T ~uint32]() func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint64x2 variants

// ScalarToUint64x2 converts a stream of scalar uint64 values into uint64x2 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 2 elements. When the buffer
// is full (exactly 2 values), a complete Uint64x2 vector is emitted. On source completion,
// any remaining values in the buffer (less than 2) are emitted as a partial vector.
func ScalarToUint64x2[T ~uint64]() func(ro.Observable[T]) ro.Observable[*archsimd.Uint64x2] {
	_ = "STUB: not implemented"
	return nil
}

// Uint64x2ToScalar converts a stream of uint64x2 vectors into scalar uint64 values.
//
// Buffer semantics: Each Uint64x2 vector is fully unpacked into 2 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 2 elements set) will emit zero values for unused lanes.
func Uint64x2ToScalar[T ~uint64]() func(ro.Observable[*archsimd.Uint64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float32x4 variants

// ScalarToFloat32x4 converts a stream of scalar float32 values into float32x4 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 4 elements. When the buffer
// is full (exactly 4 values), a complete Float32x4 vector is emitted. On source completion,
// any remaining values in the buffer (less than 4) are emitted as a partial vector.
func ScalarToFloat32x4[T ~float32]() func(ro.Observable[T]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// Float32x4ToScalar converts a stream of float32x4 vectors into scalar float32 values.
//
// Buffer semantics: Each Float32x4 vector is fully unpacked into 4 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 4 elements set) will emit zero values for unused lanes.
func Float32x4ToScalar[T ~float32]() func(ro.Observable[*archsimd.Float32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float64x2 variants

// ScalarToFloat64x2 converts a stream of scalar float64 values into float64x2 vectors.
//
// Buffer semantics: Values are accumulated into a buffer of 2 elements. When the buffer
// is full (exactly 2 values), a complete Float64x2 vector is emitted. On source completion,
// any remaining values in the buffer (less than 2) are emitted as a partial vector.
func ScalarToFloat64x2[T ~float64]() func(ro.Observable[T]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// Float64x2ToScalar converts a stream of float64x2 vectors into scalar float64 values.
//
// Buffer semantics: Each Float64x2 vector is fully unpacked into 2 scalar values, emitting
// each element individually. All elements from each input vector are emitted in order.
// Partial vectors (fewer than 2 elements set) will emit zero values for unused lanes.
func Float64x2ToScalar[T ~float64]() func(ro.Observable[*archsimd.Float64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
