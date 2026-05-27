// Copyright 2025 samber.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.apache.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ro

import (
	"time"

	"github.com/samber/lo"
)

// Of creates an Observable that emits some values you specify.
// Play: https://go.dev/play/p/Zp5LgHgvJ59
func Of[T any](values ...T) Observable[T] { _ = "STUB: not implemented"; return nil }

// Just is an alias for Of.
// Play: https://go.dev/play/p/A5S2McqqfqE
func Just[T any](values ...T) Observable[T] { _ = "STUB: not implemented"; return nil }

// Start creates an Observable that emits lazily a single value.
// Play: https://go.dev/play/p/Jz7oyagu07u
func Start[T any](cb func() T) Observable[T] { _ = "STUB: not implemented"; return nil }

// Timer creates an Observable that emits a value after a specified duration.
// Play: https://go.dev/play/p/hMkNLEqpcy3
func Timer(duration time.Duration) Observable[time.Duration] { _ = "STUB: not implemented"; return nil }

// Interval creates an Observable that emits an infinite sequence of ascending
// integers, with a constant interval between them. The first value is not emitted
// immediately, but after the first interval has passed.
// Play: https://go.dev/play/p/7yskMPPFHA7
func Interval(interval time.Duration) Observable[int64] { _ = "STUB: not implemented"; return nil }

// `ok` is not expected to be false, because the go runtime will close the channel itself

// IntervalWithInitial creates an Observable that emits an infinite sequence of ascending
// integers, with a constant interval between them. The first value is not emitted immediately,
// but after the initial interval has passed. The first interval is `initial`, and the subsequent
// intervals are `interval`. The first value is emitted after `initial` time has passed.
// Play: https://go.dev/play/p/Xhi6c336ldy
func IntervalWithInitial(initial, interval time.Duration) Observable[int64] {
	_ = "STUB: not implemented"
	return nil
}

// Synchronous initial value when first tick must be triggered immediately.

// `ok` is not expected to be false, because the go runtime will close the channel itself
// exclude initial tick when it is immediately

// `ok` is not expected to be false, because the go runtime will close the channel itself

// Range creates an Observable that emits a range of integers.
// The range is [start:end), so `start` is emitted but not `end`.
// If `start` is equal to `end`, an empty Observable is returned.
// If `start` is greater than `end`, the emitted values are in
// descending order. The step is 1.
// Play: https://go.dev/play/p/5XAXfNrtJm2
func Range(start, end int64) Observable[int64] { _ = "STUB: not implemented"; return nil }

// RangeWithStep creates an Observable that emits a range of floats.
// The range is [start:end), so `start` is emitted but not `end`.
// If `start` is equal to `end`, an empty Observable is returned.
// If `start` is greater than `end`, the emitted values are in
// descending order.
// The step must be greater than 0.
// Play: https://go.dev/play/p/EOG0tIVjUKC
func RangeWithStep(start, end, step float64) Observable[float64] {
	_ = "STUB: not implemented"
	return nil
}

// RangeWithInterval creates an Observable that emits a range of integers.
// The range is [start:end), so `start` is emitted but not `end`.
// If `start` is equal to `end`, an empty Observable is returned.
// If `start` is greater than `end`, the emitted values are in
// descending order. The interval is the time between each value.
// The first value is emitted after the first interval has passed.
// The step is 1.
// Play: https://go.dev/play/p/Y_1l6BDbMSi
func RangeWithInterval(start, end int64, interval time.Duration) Observable[int64] {
	_ = "STUB: not implemented"
	return nil
}

// RangeWithStepAndInterval creates an Observable that emits a range of floats.
// The range is [start:end), so `start` is emitted but not `end`.
// If `start` is equal to `end`, an empty Observable is returned.
// If `start` is greater than `end`, the emitted values are in
// descending order. The step must be greater than 0.
// The interval is the time between each value.
// The first value is emitted after the first interval has passed.
// Play: https://go.dev/play/p/kdAEsGwfqw9
func RangeWithStepAndInterval(start, end, step float64, interval time.Duration) Observable[float64] {
	_ = "STUB: not implemented"
	return nil
}

// Repeat creates an Observable that emits a single value multiple times.
// This is a creation operator. The pipeable equivalent is `RepeatWith`.
// Play: https://go.dev/play/p/CUvh_TYALNe
func Repeat[T any](item T, count int64) Observable[T] { _ = "STUB: not implemented"; return nil }

// RepeatWithInterval creates an Observable that emits a single value multiple times.
// The interval is the time between each value. The first value is emitted
// after the first interval has passed.
// Play: https://go.dev/play/p/4PK5Zt2sGze
func RepeatWithInterval[T any](item T, count int64, interval time.Duration) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FromChannel creates an Observable from a channel. Closing the
// channel will complete the Observable.
// Play: https://go.dev/play/p/x0u4eaOzYln
func FromChannel[T any](in <-chan T) Observable[T] { _ = "STUB: not implemented"; return nil }

// FromSlice creates an Observable from a slice. The values are emitted
// in the order they are in the slice.
// Play: https://go.dev/play/p/BNhnqoQn0tP
func FromSlice[T any](collections ...[]T) Observable[T] { _ = "STUB: not implemented"; return nil }

// Empty creates an Observable that emits no values and completes immediately.
// Play: https://go.dev/play/p/D1JWkPG4NFK
func Empty[T any]() Observable[T] { _ = "STUB: not implemented"; return nil }

// Never creates an Observable that emits no values and never completes.
// This is useful for testing or when combining with other Observables.
// Play: https://go.dev/play/p/GHzcVYaEvN8
func Never() Observable[struct{}] { _ = "STUB: not implemented"; return nil }

// Throw creates an Observable that emits an error and completes immediately.
// Play: https://go.dev/play/p/1TBK8LdDRJF
func Throw[T any](err error) Observable[T] {
	_ = "STUB: not implemented"
	// `nil` is a valid value for `err`
	return nil
}

// Defer creates an Observable that waits until an Observer subscribes to it,
// and then it creates an Observable for each Observer. This is useful for
// creating Observables that depend on some external state that is not
// available at the time of creation. The `cb` function is called for each
// Observer that subscribes to the Observable.
// Play: https://go.dev/play/p/wyVzordmkK0
func Defer[T any](factory func() Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Future creates an Observable that waits until an Observer subscribes to it,
// and then it emits either a value or an error, returned by the `factory` function.
//
// This is useful for creating Observables that depend on some external state
// that is not available at the time of creation. The `factory` function is called
// for each Observer that subscribes to the Observable.
func Future[T any](factory func() (T, error)) Observable[T] { _ = "STUB: not implemented"; return nil }

// Merge merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
// Play: https://go.dev/play/p/hX2xPyeO3M9
func Merge[T any](sources ...Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// CombineLatest2 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/mzpJyg7plnm
func CombineLatest2[A, B any](obsA Observable[A], obsB Observable[B]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatest3 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
func CombineLatest3[A, B, C any](obsA Observable[A], obsB Observable[B], obsC Observable[C]) Observable[lo.Tuple3[A, B, C]] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatest4 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/mzpJyg7plnm
func CombineLatest4[A, B, C, D any](obsA Observable[A], obsB Observable[B], obsC Observable[C], obsD Observable[D]) Observable[lo.Tuple4[A, B, C, D]] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatest5 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/mzpJyg7plnm
func CombineLatest5[A, B, C, D, E any](obsA Observable[A], obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E]) Observable[lo.Tuple5[A, B, C, D, E]] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatestAny combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/mzpJyg7plnm
func CombineLatestAny(sources ...Observable[any]) Observable[[]any] {
	_ = "STUB: not implemented"
	return nil
}

// Zip combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip[T any](sources ...Observable[T]) Observable[[]T] { _ = "STUB: not implemented"; return nil }

// Zip2 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip2[A, B any](obsA Observable[A], obsB Observable[B]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// Zip3 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip3[A, B, C any](obsA Observable[A], obsB Observable[B], obsC Observable[C]) Observable[lo.Tuple3[A, B, C]] {
	_ = "STUB: not implemented"
	return nil
}

// Zip4 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip4[A, B, C, D any](obsA Observable[A], obsB Observable[B], obsC Observable[C], obsD Observable[D]) Observable[lo.Tuple4[A, B, C, D]] {
	_ = "STUB: not implemented"
	return nil
}

// Zip5 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip5[A, B, C, D, E any](obsA Observable[A], obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E]) Observable[lo.Tuple5[A, B, C, D, E]] {
	_ = "STUB: not implemented"
	return nil
}

// Zip6 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/5YxbQ5jNzjQ
func Zip6[A, B, C, D, E, F any](obsA Observable[A], obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E], obsF Observable[F]) Observable[lo.Tuple6[A, B, C, D, E, F]] {
	_ = "STUB: not implemented"
	return nil
}

// Concat concatenates the source Observable with other Observables. It subscribes
// to each inner Observable only after the previous one completes, maintaining their
// order. It completes when all inner Observables are done.
// Play: https://go.dev/play/p/DFokqIXIguM
func Concat[T any](obs ...Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// Race creates an Observable that mirrors the first source Observable to
// emit a next, error or complete notification from the combination of the
// Observable sources. It cancels the subscriptions to all other Observables.
// It completes when the source Observable completes. If the source Observable
// emits an error, the error is emitted by the resulting Observable.
// Play: https://go.dev/play/p/5VzGFd62SMC
func Race[T any](sources ...Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// Amb is an alias for Race.
// Play: https://go.dev/play/p/-YvhnpQFVNS
func Amb[T any](sources ...Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// RandIntN creates an Observable that emits random int values in the range [0, n).
// The count is the number of values to emit.
// Play: https://go.dev/play/p/4m7T5j-7i3a
func RandIntN(n, count int) Observable[int] { _ = "STUB: not implemented"; return nil }

// RandFloat64 creates an Observable that emits random float64 values in the range [0, 1).
// The count is the number of values to emit.
// Play: https://go.dev/play/p/MRuy8rUpTve
func RandFloat64(count int) Observable[float64] { _ = "STUB: not implemented"; return nil }
