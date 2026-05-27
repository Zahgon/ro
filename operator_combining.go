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

//nolint:nestif,funlen,gocyclo
package ro

import (
	"context"
	"sync"

	"github.com/samber/lo"
)

// MergeWith merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/6QpUzcdRWJl
func MergeWith[T any](observables ...Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeWith1 merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/P47lkUFpYq7
func MergeWith1[T any](obsB Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeWith2 merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/LOQ3YbuDyC9
func MergeWith2[T any](obsB, obsC Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeWith3 merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/pMQ5bNOlWj9
func MergeWith3[T any](obsB, obsC, obsD Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeWith4 merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/FvJTHVOe52s
func MergeWith4[T any](obsB, obsC, obsD, obsE Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeWith5 merges the values from all observables to a single observable result.
// It subscribes to each inner Observable, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
//
// It is a curried function that takes the first Observable as an argument.
// Play: https://go.dev/play/p/kR3rFF7Bw-i
func MergeWith5[T any](obsB, obsC, obsD, obsE, obsF Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// MergeAll converts a higher-order Observable into a first-order Observable which
// concurrently delivers all values that are emitted on the inner Observables.
// It subscribes to each inner Observable as they arrive, and emits all values
// from each inner Observable, maintaining their order. It completes when all
// inner Observables are done.
// Play: https://go.dev/play/p/m3nHZZJbwMF
func MergeAll[T any]() func(Observable[Observable[T]]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// atomic.Value has been introduced in go 1.19 and this library support go 1.18

// default value is not 0, because it counts the outer Observable `sources`

// when equal to 0, it means both the outer and inner Observables are done

// MergeMap applies a projection function to each item emitted by the source
// Observable and then merges the results into a single Observable.
// Play: https://go.dev/play/p/NwEyrLITshG
func MergeMap[T, R any](projection func(item T) Observable[R]) func(Observable[T]) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// MergeMapWithContext applies a projection function to each item emitted by the source
// Observable and then merges the results into a single Observable.
// Play: https://go.dev/play/p/i2Ru9sUdL-x
func MergeMapWithContext[T, R any](projection func(ctx context.Context, item T) Observable[R]) func(Observable[T]) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// MergeMapI applies a projection function to each item emitted by the source
// Observable and then merges the results into a single Observable.
// Play: https://go.dev/play/p/dPDI7ch4g0i
func MergeMapI[T, R any](projection func(item T, index int64) Observable[R]) func(Observable[T]) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// MergeMapIWithContext applies a projection function to each item emitted by the source
// Observable and then merges the results into a single Observable.
// Play: https://go.dev/play/p/8Ih5mCaDbB8
func MergeMapIWithContext[T, R any](projection func(ctx context.Context, item T, index int64) (context.Context, Observable[R])) func(Observable[T]) Observable[R] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatestWith combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/yq7G8eItuzO
func CombineLatestWith[A, B any](obsB Observable[B]) func(Observable[A]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// CombineLatestWith1 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/KXb19PPjCb1
func CombineLatestWith1[A, B any](obsB Observable[B]) func(Observable[A]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// 0: not done
// 1: partially done
// 2: done
// 3: error

// CombineLatestWith2 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/hPDCDwEOB84
func CombineLatestWith2[A, B, C any](obsB Observable[B], obsC Observable[C]) func(Observable[A]) Observable[lo.Tuple3[A, B, C]] {
	_ = "STUB: not implemented"
	return nil
}

// 0: not done
// 1: partially done
// 2: partially done
// 3: done
// 4: error

// CombineLatestWith3 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/PcMxo8yakQq
func CombineLatestWith3[A, B, C, D any](obsB Observable[B], obsC Observable[C], obsD Observable[D]) func(Observable[A]) Observable[lo.Tuple4[A, B, C, D]] {
	_ = "STUB: not implemented"
	return nil
}

// 0: not done
// 1: partially done
// 2: partially done
// 3: partially done
// 4: done
// 5: error

// CombineLatestWith4 combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
func CombineLatestWith4[A, B, C, D, E any](obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E]) func(Observable[A]) Observable[lo.Tuple5[A, B, C, D, E]] {
	_ = "STUB: not implemented"
	return nil
}

// 0: not done
// 1: partially done
// 2: partially done
// 3: partially done
// 4: partially done
// 5: done
// 6: error

// CombineLatestAll combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/nT1qq9ipwZL
func CombineLatestAll[T any]() func(Observable[Observable[T]]) Observable[[]T] {
	_ = "STUB: not implemented"
	return nil
}

// -1: error
// 0: done
// 1: partially done
// 2: partially done
// .: partially done
// .: partially done
// n: not done

// init

// inner subscriptions

// outer subscription

// CombineLatestAllAny combines the values from the source Observable with the latest
// values from the other Observables. It will only emit when all Observables have
// emitted at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/nKMychGg9KH
func CombineLatestAllAny() func(Observable[Observable[any]]) Observable[[]any] {
	_ = "STUB: not implemented"
	return nil
}

// ConcatWith concatenates the source Observable with other Observables. It subscribes
// to each inner Observable only after the previous one completes, maintaining their
// order. It completes when all inner Observables are done.
//
// It is a curried function that takes the other Observables as arguments.
// Play: https://go.dev/play/p/nRHRSR2yNvd
func ConcatWith[T any](obs ...Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ConcatAll concatenates the source Observable with other Observables. It subscribes
// to each inner Observable only after the previous one completes, maintaining their
// order. It completes when all inner Observables are done.
// Play: https://go.dev/play/p/zygV4Ld9tcv
func ConcatAll[T any]() func(Observable[Observable[T]]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// `subscriptions` cancels `sub` when it unsubscribes
// but `sub` cannot unsubscribe `subscriptions`

// StartWith emits the given values before emitting the values from the source Observable.
// Play: https://go.dev/play/p/vS_gIw8Ce1C
func StartWith[T any](prefixes ...T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// EndWith emits the given values after emitting the values from the source Observable.
// Play: https://go.dev/play/p/9FPyf3bqJk_n
func EndWith[T any](suffixes ...T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Pairwise emits the previous and current values as a pair of two values.
// Play: https://go.dev/play/p/0YujgFTL4e0
func Pairwise[T any]() func(Observable[T]) Observable[[]T] { _ = "STUB: not implemented"; return nil }

// RaceWith creates an Observable that mirrors the first source Observable to
// emit a next, error or complete notification from the combination of the
// Observable to which the operator is applied and supplied Observables. It
// cancels the subscriptions to all other Observables. It completes when the
// source Observable completes. If the source Observable errors, it errors with
// the same error.
//
// It is a curried function that takes the other Observables as arguments.
// Play: https://go.dev/play/p/5VzGFd62SMC
func RaceWith[T any](sources ...Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Check if a winner was determined during subscription

// No winner yet, store the subscription

// Another source won, unsubscribe this one

// If this source won, keep the subscription active

type zipDestination interface {
	ErrorWithContext(context.Context, error)
	CompleteWithContext(context.Context)
}

// This code is dity but much more concise than the original implementation.
func zipInnerSubscription[T any](subscriberCtx context.Context, obs Observable[T], mu *sync.Mutex, values *[]*T, completed *bool, onUpdate func(context.Context), destination zipDestination, subscriptions Subscription) {
	_ = "STUB: not implemented"
	return
}

// ZipWith combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/RmErtE3pHjb
func ZipWith[A, B any](obsB Observable[B]) func(Observable[A]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil

	// ZipWith1 combines the values from the source Observable with the latest values
	// from the other Observables. It emits only when all Observables have emitted
	// at least one value. It completes when the source Observable completes.
	//
	// It is a curried function that takes the other Observable as an argument.
}

func ZipWith1[A, B any](obsB Observable[B]) func(Observable[A]) Observable[lo.Tuple2[A, B]] {
	_ = "STUB: not implemented"
	return nil
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

// ZipWith2 combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/MMq82Rkb0oh
func ZipWith2[A, B, C any](obsB Observable[B], obsC Observable[C]) func(Observable[A]) Observable[lo.Tuple3[A, B, C]] {
	_ = "STUB: not implemented"
	return nil
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

// ZipWith3 combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
func ZipWith3[A, B, C, D any](obsB Observable[B], obsC Observable[C], obsD Observable[D]) func(Observable[A]) Observable[lo.Tuple4[A, B, C, D]] {
	_ = "STUB: not implemented"
	return nil
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

// ZipWith4 combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
func ZipWith4[A, B, C, D, E any](obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E]) func(Observable[A]) Observable[lo.Tuple5[A, B, C, D, E]] {
	_ = "STUB: not implemented"
	return nil
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

// ZipWith5 combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
//
// It is a curried function that takes the other Observable as an argument.
// Play: https://go.dev/play/p/OJz-AVo0-hY
func ZipWith5[A, B, C, D, E, F any](obsB Observable[B], obsC Observable[C], obsD Observable[D], obsE Observable[E], obsF Observable[F]) func(Observable[A]) Observable[lo.Tuple6[A, B, C, D, E, F]] {
	_ = "STUB: not implemented"
	return nil
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

func zipAllInnerSubscriptions[T any](outerCtx context.Context, sources []Observable[T], destination Observer[[]T]) Teardown {
	_ = "STUB: not implemented"
	return *new(Teardown)
}

// unlock before calling destination.Next to prevent long locks

// @TODO: Send the last context ?

// @TODO: Send the last context ?

// free memory

// ZipAll combines the values from the source Observable with the latest values
// from the other Observables. It emits only when all Observables have emitted
// at least one value. It completes when the source Observable completes.
// Play: https://go.dev/play/p/FcpgTItKX-Q
func ZipAll[T any]() func(Observable[Observable[T]]) Observable[[]T] {
	_ = "STUB: not implemented"
	return nil
}

// First, we consume the high-order observable...

// ...then we zip all inner observables.
