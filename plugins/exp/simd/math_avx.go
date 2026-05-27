//go:build go1.26 && goexperiment.simd && amd64

package rosimd

import (
	"simd/archsimd"

	"github.com/samber/ro"
)

// Int8x16 variants

// AddInt8x16 adds a scalar value to each element of int8x16 vectors.
func AddInt8x16[T ~int8](number T) func(ro.Observable[*archsimd.Int8x16]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt8x16 subtracts a scalar value from each element of int8x16 vectors.
func SubInt8x16[T ~int8](number T) func(ro.Observable[*archsimd.Int8x16]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt8x16 clamps each element of int8x16 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt8x16[T ~int8](minValue, maxValue T) func(ro.Observable[*archsimd.Int8x16]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt8x16 applies a minimum bound to each element of int8x16 vectors.
func MinInt8x16[T ~int8](minValue T) func(ro.Observable[*archsimd.Int8x16]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt8x16 applies a maximum bound to each element of int8x16 vectors.
func MaxInt8x16[T ~int8](maxValue T) func(ro.Observable[*archsimd.Int8x16]) ro.Observable[*archsimd.Int8x16] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt8x16 reduces int8x16 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int8 arithmetic,
// which can only hold values from -128 to 127. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumInt8x16[T ~int8]() func(ro.Observable[*archsimd.Int8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt8x16 reduces int8x16 vectors to their minimum value.
func ReduceMinInt8x16[T ~int8]() func(ro.Observable[*archsimd.Int8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt8x16 reduces int8x16 vectors to their maximum value.
func ReduceMaxInt8x16[T ~int8]() func(ro.Observable[*archsimd.Int8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int16x8 variants

// AddInt16x8 adds a scalar value to each element of int16x8 vectors.
func AddInt16x8[T ~int16](number T) func(ro.Observable[*archsimd.Int16x8]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt16x8 subtracts a scalar value from each element of int16x8 vectors.
func SubInt16x8[T ~int16](number T) func(ro.Observable[*archsimd.Int16x8]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt16x8 clamps each element of int16x8 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt16x8[T ~int16](minValue, maxValue T) func(ro.Observable[*archsimd.Int16x8]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt16x8 applies a minimum bound to each element of int16x8 vectors.
func MinInt16x8[T ~int16](minValue T) func(ro.Observable[*archsimd.Int16x8]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt16x8 applies a maximum bound to each element of int16x8 vectors.
func MaxInt16x8[T ~int16](maxValue T) func(ro.Observable[*archsimd.Int16x8]) ro.Observable[*archsimd.Int16x8] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt16x8 reduces int16x8 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int16 arithmetic,
// which can only hold values from -32768 to 32767. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumInt16x8[T ~int16]() func(ro.Observable[*archsimd.Int16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt16x8 reduces int16x8 vectors to their minimum value.
func ReduceMinInt16x8[T ~int16]() func(ro.Observable[*archsimd.Int16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt16x8 reduces int16x8 vectors to their maximum value.
func ReduceMaxInt16x8[T ~int16]() func(ro.Observable[*archsimd.Int16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int32x4 variants

// AddInt32x4 adds a scalar value to each element of int32x4 vectors.
func AddInt32x4[T ~int32](number T) func(ro.Observable[*archsimd.Int32x4]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt32x4 subtracts a scalar value from each element of int32x4 vectors.
func SubInt32x4[T ~int32](number T) func(ro.Observable[*archsimd.Int32x4]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ClampInt32x4 clamps each element of int32x4 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampInt32x4[T ~int32](minValue, maxValue T) func(ro.Observable[*archsimd.Int32x4]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MinInt32x4 applies a minimum bound to each element of int32x4 vectors.
func MinInt32x4[T ~int32](minValue T) func(ro.Observable[*archsimd.Int32x4]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MaxInt32x4 applies a maximum bound to each element of int32x4 vectors.
func MaxInt32x4[T ~int32](maxValue T) func(ro.Observable[*archsimd.Int32x4]) ro.Observable[*archsimd.Int32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt32x4 reduces int32x4 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int32 arithmetic.
// For large sums that may exceed int32 range, consider using int64 or a different reduction strategy.
func ReduceSumInt32x4[T ~int32]() func(ro.Observable[*archsimd.Int32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinInt32x4 reduces int32x4 vectors to their minimum value.
func ReduceMinInt32x4[T ~int32]() func(ro.Observable[*archsimd.Int32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxInt32x4 reduces int32x4 vectors to their maximum value.
func ReduceMaxInt32x4[T ~int32]() func(ro.Observable[*archsimd.Int32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Int64x2 variants

// AddInt64x2 adds a scalar value to each element of int64x2 vectors.
func AddInt64x2[T ~int64](number T) func(ro.Observable[*archsimd.Int64x2]) ro.Observable[*archsimd.Int64x2] {
	_ = "STUB: not implemented"
	return nil
}

// SubInt64x2 subtracts a scalar value from each element of int64x2 vectors.
func SubInt64x2[T ~int64](number T) func(ro.Observable[*archsimd.Int64x2]) ro.Observable[*archsimd.Int64x2] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumInt64x2 reduces int64x2 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using int64 arithmetic.
// For extremely large sums that may exceed int64 range, consider using a big integer library
// or a different reduction strategy.
func ReduceSumInt64x2[T ~int64]() func(ro.Observable[*archsimd.Int64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint8x16 variants

// AddUint8x16 adds a scalar value to each element of uint8x16 vectors.
func AddUint8x16[T ~uint8](number T) func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint8x16 subtracts a scalar value from each element of uint8x16 vectors.
func SubUint8x16[T ~uint8](number T) func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint8x16 clamps each element of uint8x16 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint8x16[T ~uint8](minValue, maxValue T) func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint8x16 applies a minimum bound to each element of uint8x16 vectors.
func MinUint8x16[T ~uint8](minValue T) func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint8x16 applies a maximum bound to each element of uint8x16 vectors.
func MaxUint8x16[T ~uint8](maxValue T) func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[*archsimd.Uint8x16] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint8x16 reduces uint8x16 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint8 arithmetic,
// which can only hold values from 0 to 255. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumUint8x16[T ~uint8]() func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint8x16 reduces uint8x16 vectors to their minimum value.
func ReduceMinUint8x16[T ~uint8]() func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint8x16 reduces uint8x16 vectors to their maximum value.
func ReduceMaxUint8x16[T ~uint8]() func(ro.Observable[*archsimd.Uint8x16]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint16x8 variants

// AddUint16x8 adds a scalar value to each element of uint16x8 vectors.
func AddUint16x8[T ~uint16](number T) func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint16x8 subtracts a scalar value from each element of uint16x8 vectors.
func SubUint16x8[T ~uint16](number T) func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint16x8 clamps each element of uint16x8 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint16x8[T ~uint16](minValue, maxValue T) func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint16x8 applies a minimum bound to each element of uint16x8 vectors.
func MinUint16x8[T ~uint16](minValue T) func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint16x8 applies a maximum bound to each element of uint16x8 vectors.
func MaxUint16x8[T ~uint16](maxValue T) func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[*archsimd.Uint16x8] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint16x8 reduces uint16x8 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint16 arithmetic,
// which can only hold values from 0 to 65535. For large sums, consider using a wider
// type or a different reduction strategy.
func ReduceSumUint16x8[T ~uint16]() func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint16x8 reduces uint16x8 vectors to their minimum value.
func ReduceMinUint16x8[T ~uint16]() func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint16x8 reduces uint16x8 vectors to their maximum value.
func ReduceMaxUint16x8[T ~uint16]() func(ro.Observable[*archsimd.Uint16x8]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint32x4 variants

// AddUint32x4 adds a scalar value to each element of uint32x4 vectors.
func AddUint32x4[T ~uint32](number T) func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint32x4 subtracts a scalar value from each element of uint32x4 vectors.
func SubUint32x4[T ~uint32](number T) func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ClampUint32x4 clamps each element of uint32x4 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampUint32x4[T ~uint32](minValue, maxValue T) func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MinUint32x4 applies a minimum bound to each element of uint32x4 vectors.
func MinUint32x4[T ~uint32](minValue T) func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MaxUint32x4 applies a maximum bound to each element of uint32x4 vectors.
func MaxUint32x4[T ~uint32](maxValue T) func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[*archsimd.Uint32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint32x4 reduces uint32x4 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint32 arithmetic.
// For large sums that may exceed uint32 range, consider using uint64 or a different reduction strategy.
func ReduceSumUint32x4[T ~uint32]() func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinUint32x4 reduces uint32x4 vectors to their minimum value.
func ReduceMinUint32x4[T ~uint32]() func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxUint32x4 reduces uint32x4 vectors to their maximum value.
func ReduceMaxUint32x4[T ~uint32]() func(ro.Observable[*archsimd.Uint32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Uint64x2 variants

// AddUint64x2 adds a scalar value to each element of uint64x2 vectors.
func AddUint64x2[T ~uint64](number T) func(ro.Observable[*archsimd.Uint64x2]) ro.Observable[*archsimd.Uint64x2] {
	_ = "STUB: not implemented"
	return nil
}

// SubUint64x2 subtracts a scalar value from each element of uint64x2 vectors.
func SubUint64x2[T ~uint64](number T) func(ro.Observable[*archsimd.Uint64x2]) ro.Observable[*archsimd.Uint64x2] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumUint64x2 reduces uint64x2 vectors to their sum.
//
// WARNING: This function may overflow. The result is accumulated using uint64 arithmetic.
// For extremely large sums that may exceed uint64 range, consider using a big integer library
// or a different reduction strategy.
func ReduceSumUint64x2[T ~uint64]() func(ro.Observable[*archsimd.Uint64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float32x4 variants

// AddFloat32x4 adds a scalar value to each element of float32x4 vectors.
func AddFloat32x4[T ~float32](number T) func(ro.Observable[*archsimd.Float32x4]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// SubFloat32x4 subtracts a scalar value from each element of float32x4 vectors.
func SubFloat32x4[T ~float32](number T) func(ro.Observable[*archsimd.Float32x4]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ClampFloat32x4 clamps each element of float32x4 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampFloat32x4[T ~float32](minValue, maxValue T) func(ro.Observable[*archsimd.Float32x4]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MinFloat32x4 applies a minimum bound to each element of float32x4 vectors.
func MinFloat32x4[T ~float32](minValue T) func(ro.Observable[*archsimd.Float32x4]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// MaxFloat32x4 applies a maximum bound to each element of float32x4 vectors.
func MaxFloat32x4[T ~float32](maxValue T) func(ro.Observable[*archsimd.Float32x4]) ro.Observable[*archsimd.Float32x4] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumFloat32x4 reduces float32x4 vectors to their sum.
func ReduceSumFloat32x4[T ~float32]() func(ro.Observable[*archsimd.Float32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinFloat32x4 reduces float32x4 vectors to their minimum value.
func ReduceMinFloat32x4[T ~float32]() func(ro.Observable[*archsimd.Float32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxFloat32x4 reduces float32x4 vectors to their maximum value.
func ReduceMaxFloat32x4[T ~float32]() func(ro.Observable[*archsimd.Float32x4]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Float64x2 variants

// AddFloat64x2 adds a scalar value to each element of float64x2 vectors.
func AddFloat64x2[T ~float64](number T) func(ro.Observable[*archsimd.Float64x2]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// SubFloat64x2 subtracts a scalar value from each element of float64x2 vectors.
func SubFloat64x2[T ~float64](number T) func(ro.Observable[*archsimd.Float64x2]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// ClampFloat64x2 clamps each element of float64x2 vectors between min and max values.
// Values outside the range are clamped to the nearest valid value.
func ClampFloat64x2[T ~float64](minValue, maxValue T) func(ro.Observable[*archsimd.Float64x2]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// MinFloat64x2 applies a minimum bound to each element of float64x2 vectors.
func MinFloat64x2[T ~float64](minValue T) func(ro.Observable[*archsimd.Float64x2]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// MaxFloat64x2 applies a maximum bound to each element of float64x2 vectors.
func MaxFloat64x2[T ~float64](maxValue T) func(ro.Observable[*archsimd.Float64x2]) ro.Observable[*archsimd.Float64x2] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceSumFloat64x2 reduces float64x2 vectors to their sum.
func ReduceSumFloat64x2[T ~float64]() func(ro.Observable[*archsimd.Float64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMinFloat64x2 reduces float64x2 vectors to their minimum value.
func ReduceMinFloat64x2[T ~float64]() func(ro.Observable[*archsimd.Float64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ReduceMaxFloat64x2 reduces float64x2 vectors to their maximum value.
func ReduceMaxFloat64x2[T ~float64]() func(ro.Observable[*archsimd.Float64x2]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
