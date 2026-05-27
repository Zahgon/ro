//go:build go1.26 && goexperiment.simd && amd64

package rosimd

import (
	"simd/archsimd"

	"github.com/samber/ro"
)

// AVX2 Int8x32 variants

// AddInt8x32 adds a scalar value to each element of int8x32 vectors.
func AddInt8x32[T ~int8](number T) func(ro.Observable[*archsimd.Int8x32]) ro.Observable[*archsimd.Int8x32] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt8x32 subtracts a scalar value from each element of int8x32 vectors.
func SubInt8x32[T ~int8](number T) func(ro.Observable[*archsimd.Int8x32]) ro.Observable[*archsimd.Int8x32] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt8x32 clamps each element of int8x32 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt8x32[T ~int8](minValue, maxValue T) func(ro.Observable[*archsimd.Int8x32]) ro.Observable[*archsimd.Int8x32] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt8x32 applies a minimum bound to each element of int8x32 vectors.
func MinInt8x32[T ~int8](minValue T) func(ro.Observable[*archsimd.Int8x32]) ro.Observable[*archsimd.Int8x32] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt8x32 applies a maximum bound to each element of int8x32 vectors.
func MaxInt8x32[T ~int8](maxValue T) func(ro.Observable[*archsimd.Int8x32]) ro.Observable[*archsimd.Int8x32] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt8x32 reduces int8x32 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int8 arithmetic,
// which can only hold values from -128 to 127. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumInt8x32[T ~int8]() func(ro.Observable[*archsimd.Int8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt8x32 reduces int8x32 vectors to their minimum value.
func ReduceMinInt8x32[T ~int8]() func(ro.Observable[*archsimd.Int8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt8x32 reduces int8x32 vectors to their maximum value.
func ReduceMaxInt8x32[T ~int8]() func(ro.Observable[*archsimd.Int8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float32x8 AVX2 operators (most commonly used AVX2 type)

// AddFloat32x8 adds a scalar value to each element of float32x8 vectors.
func AddFloat32x8[T ~float32](number T) func(ro.Observable[*archsimd.Float32x8]) ro.Observable[*archsimd.Float32x8] {
	_ = "STUB: not implemented"
	return nil
}

// SubFloat32x8 subtracts a scalar value from each element of float32x8 vectors.
func SubFloat32x8[T ~float32](number T) func(ro.Observable[*archsimd.Float32x8]) ro.Observable[*archsimd.Float32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ClampFloat32x8 clamps each element of float32x8 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampFloat32x8[T ~float32](minValue, maxValue T) func(ro.Observable[*archsimd.Float32x8]) ro.Observable[*archsimd.Float32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MinFloat32x8 applies a minimum bound to each element of float32x8 vectors.
func MinFloat32x8[T ~float32](minValue T) func(ro.Observable[*archsimd.Float32x8]) ro.Observable[*archsimd.Float32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MaxFloat32x8 applies a maximum bound to each element of float32x8 vectors.
func MaxFloat32x8[T ~float32](maxValue T) func(ro.Observable[*archsimd.Float32x8]) ro.Observable[*archsimd.Float32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumFloat32x8 reduces float32x8 vectors to their sum.
func ReduceSumFloat32x8[T ~float32]() func(ro.Observable[*archsimd.Float32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinFloat32x8 reduces float32x8 vectors to their minimum value.
func ReduceMinFloat32x8[T ~float32]() func(ro.Observable[*archsimd.Float32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxFloat32x8 reduces float32x8 vectors to their maximum value.
func ReduceMaxFloat32x8[T ~float32]() func(ro.Observable[*archsimd.Float32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int32x8 AVX2 operators

// AddInt32x8 adds a scalar value to each element of int32x8 vectors.
func AddInt32x8[T ~int32](number T) func(ro.Observable[*archsimd.Int32x8]) ro.Observable[*archsimd.Int32x8] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt32x8 subtracts a scalar value from each element of int32x8 vectors.
func SubInt32x8[T ~int32](number T) func(ro.Observable[*archsimd.Int32x8]) ro.Observable[*archsimd.Int32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt32x8 clamps each element of int32x8 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt32x8[T ~int32](minValue, maxValue T) func(ro.Observable[*archsimd.Int32x8]) ro.Observable[*archsimd.Int32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt32x8 applies a minimum bound to each element of int32x8 vectors.
func MinInt32x8[T ~int32](minValue T) func(ro.Observable[*archsimd.Int32x8]) ro.Observable[*archsimd.Int32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt32x8 applies a maximum bound to each element of int32x8 vectors.
func MaxInt32x8[T ~int32](maxValue T) func(ro.Observable[*archsimd.Int32x8]) ro.Observable[*archsimd.Int32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt32x8 reduces int32x8 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int32 arithmetic.
// For extremely large sums that may exceed int32 range, consider using a wider type
// or a different reduction strategy.
func ReduceSumInt32x8[T ~int32]() func(ro.Observable[*archsimd.Int32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt32x8 reduces int32x8 vectors to their minimum value.
func ReduceMinInt32x8[T ~int32]() func(ro.Observable[*archsimd.Int32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt32x8 reduces int32x8 vectors to their maximum value.
func ReduceMaxInt32x8[T ~int32]() func(ro.Observable[*archsimd.Int32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int16x16 AVX2 operators

// AddInt16x16 adds a scalar value to each element of int16x16 vectors.
func AddInt16x16[T ~int16](number T) func(ro.Observable[*archsimd.Int16x16]) ro.Observable[*archsimd.Int16x16] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt16x16 subtracts a scalar value from each element of int16x16 vectors.
func SubInt16x16[T ~int16](number T) func(ro.Observable[*archsimd.Int16x16]) ro.Observable[*archsimd.Int16x16] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt16x16 clamps each element of int16x16 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt16x16[T ~int16](minValue, maxValue T) func(ro.Observable[*archsimd.Int16x16]) ro.Observable[*archsimd.Int16x16] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt16x16 applies a minimum bound to each element of int16x16 vectors.
func MinInt16x16[T ~int16](minValue T) func(ro.Observable[*archsimd.Int16x16]) ro.Observable[*archsimd.Int16x16] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt16x16 applies a maximum bound to each element of int16x16 vectors.
func MaxInt16x16[T ~int16](maxValue T) func(ro.Observable[*archsimd.Int16x16]) ro.Observable[*archsimd.Int16x16] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt16x16 reduces int16x16 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int16 arithmetic,
// which can only hold values from -32768 to 32767. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumInt16x16[T ~int16]() func(ro.Observable[*archsimd.Int16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt16x16 reduces int16x16 vectors to their minimum value.
func ReduceMinInt16x16[T ~int16]() func(ro.Observable[*archsimd.Int16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt16x16 reduces int16x16 vectors to their maximum value.
func ReduceMaxInt16x16[T ~int16]() func(ro.Observable[*archsimd.Int16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint8x32 AVX2 operators

// AddUint8x32 adds a scalar value to each element of uint8x32 vectors.
func AddUint8x32[T ~uint8](number T) func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[*archsimd.Uint8x32] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint8x32 subtracts a scalar value from each element of uint8x32 vectors.
func SubUint8x32[T ~uint8](number T) func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[*archsimd.Uint8x32] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint8x32 clamps each element of uint8x32 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint8x32[T ~uint8](minValue, maxValue T) func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[*archsimd.Uint8x32] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint8x32 applies a minimum bound to each element of uint8x32 vectors.
func MinUint8x32[T ~uint8](minValue T) func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[*archsimd.Uint8x32] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint8x32 applies a maximum bound to each element of uint8x32 vectors.
func MaxUint8x32[T ~uint8](maxValue T) func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[*archsimd.Uint8x32] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint8x32 reduces uint8x32 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint8 arithmetic,
// which can only hold values from 0 to 255. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumUint8x32[T ~uint8]() func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint8x32 reduces uint8x32 vectors to their minimum value.
func ReduceMinUint8x32[T ~uint8]() func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint8x32 reduces uint8x32 vectors to their maximum value.
func ReduceMaxUint8x32[T ~uint8]() func(ro.Observable[*archsimd.Uint8x32]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint16x16 AVX2 operators

// AddUint16x16 adds a scalar value to each element of uint16x16 vectors.
func AddUint16x16[T ~uint16](number T) func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[*archsimd.Uint16x16] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint16x16 subtracts a scalar value from each element of uint16x16 vectors.
func SubUint16x16[T ~uint16](number T) func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[*archsimd.Uint16x16] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint16x16 clamps each element of uint16x16 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint16x16[T ~uint16](minValue, maxValue T) func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[*archsimd.Uint16x16] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint16x16 applies a minimum bound to each element of uint16x16 vectors.
func MinUint16x16[T ~uint16](minValue T) func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[*archsimd.Uint16x16] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint16x16 applies a maximum bound to each element of uint16x16 vectors.
func MaxUint16x16[T ~uint16](maxValue T) func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[*archsimd.Uint16x16] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint16x16 reduces uint16x16 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint16 arithmetic.
// For extremely large sums that may exceed uint16 range, consider using a wider type
// or a different reduction strategy.
func ReduceSumUint16x16[T ~uint16]() func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint16x16 reduces uint16x16 vectors to their minimum value.
func ReduceMinUint16x16[T ~uint16]() func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint16x16 reduces uint16x16 vectors to their maximum value.
func ReduceMaxUint16x16[T ~uint16]() func(ro.Observable[*archsimd.Uint16x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint32x8 AVX2 operators

// AddUint32x8 adds a scalar value to each element of uint32x8 vectors.
func AddUint32x8[T ~uint32](number T) func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[*archsimd.Uint32x8] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint32x8 subtracts a scalar value from each element of uint32x8 vectors.
func SubUint32x8[T ~uint32](number T) func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[*archsimd.Uint32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint32x8 clamps each element of uint32x8 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint32x8[T ~uint32](minValue, maxValue T) func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[*archsimd.Uint32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint32x8 applies a minimum bound to each element of uint32x8 vectors.
func MinUint32x8[T ~uint32](minValue T) func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[*archsimd.Uint32x8] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint32x8 applies a maximum bound to each element of uint32x8 vectors.
func MaxUint32x8[T ~uint32](maxValue T) func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[*archsimd.Uint32x8] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint32x8 reduces uint32x8 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint32 arithmetic.
// For extremely large sums that may exceed uint32 range, consider using a wider type
// or a different reduction strategy.
func ReduceSumUint32x8[T ~uint32]() func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint32x8 reduces uint32x8 vectors to their minimum value.
func ReduceMinUint32x8[T ~uint32]() func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint32x8 reduces uint32x8 vectors to their maximum value.
func ReduceMaxUint32x8[T ~uint32]() func(ro.Observable[*archsimd.Uint32x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float64x4 AVX2 operators

// AddFloat64x4 adds a scalar value to each element of float64x4 vectors.
func AddFloat64x4[T ~float64](number T) func(ro.Observable[*archsimd.Float64x4]) ro.Observable[*archsimd.Float64x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubFloat64x4 subtracts a scalar value from each element of float64x4 vectors.
func SubFloat64x4[T ~float64](number T) func(ro.Observable[*archsimd.Float64x4]) ro.Observable[*archsimd.Float64x4] {
	_ = "STUB: not implemented"
	return nil
}

// ClampFloat64x4 clamps each element of float64x4 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampFloat64x4[T ~float64](minValue, maxValue T) func(ro.Observable[*archsimd.Float64x4]) ro.Observable[*archsimd.Float64x4] {
	_ = "STUB: not implemented"
	return nil
}

// MinFloat64x4 applies a minimum bound to each element of float64x4 vectors.
func MinFloat64x4[T ~float64](minValue T) func(ro.Observable[*archsimd.Float64x4]) ro.Observable[*archsimd.Float64x4] {
	_ = "STUB: not implemented"
	return nil
}

// MaxFloat64x4 applies a maximum bound to each element of float64x4 vectors.
func MaxFloat64x4[T ~float64](maxValue T) func(ro.Observable[*archsimd.Float64x4]) ro.Observable[*archsimd.Float64x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumFloat64x4 reduces float64x4 vectors to their sum.
func ReduceSumFloat64x4[T ~float64]() func(ro.Observable[*archsimd.Float64x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinFloat64x4 reduces float64x4 vectors to their minimum value.
func ReduceMinFloat64x4[T ~float64]() func(ro.Observable[*archsimd.Float64x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxFloat64x4 reduces float64x4 vectors to their maximum value.
func ReduceMaxFloat64x4[T ~float64]() func(ro.Observable[*archsimd.Float64x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int64x4 AVX2 operators

// AddInt64x4 adds a scalar value to each element of int64x4 vectors.
func AddInt64x4[T ~int64](number T) func(ro.Observable[*archsimd.Int64x4]) ro.Observable[*archsimd.Int64x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt64x4 subtracts a scalar value from each element of int64x4 vectors.
func SubInt64x4[T ~int64](number T) func(ro.Observable[*archsimd.Int64x4]) ro.Observable[*archsimd.Int64x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt64x4 reduces int64x4 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int64 arithmetic.
// For extremely large sums that may exceed int64 range, consider using a wider type
// or a different reduction strategy.
func ReduceSumInt64x4[T ~int64]() func(ro.Observable[*archsimd.Int64x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint64x4 AVX2 operators

// AddUint64x4 adds a scalar value to each element of uint64x4 vectors.
func AddUint64x4[T ~uint64](number T) func(ro.Observable[*archsimd.Uint64x4]) ro.Observable[*archsimd.Uint64x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint64x4 subtracts a scalar value from each element of uint64x4 vectors.
func SubUint64x4[T ~uint64](number T) func(ro.Observable[*archsimd.Uint64x4]) ro.Observable[*archsimd.Uint64x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint64x4 reduces uint64x4 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint64 arithmetic.
// For extremely large sums that may exceed uint64 range, consider using a wider type
// or a different reduction strategy.
func ReduceSumUint64x4[T ~uint64]() func(ro.Observable[*archsimd.Uint64x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
