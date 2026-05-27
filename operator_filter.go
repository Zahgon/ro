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
	"context"
)

// Filter emits only those items from an Observable that pass a predicate test.
// Play: https://go.dev/play/p/gjk_wULxyEW
func Filter[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FilterWithContext emits only those items from an Observable that pass a predicate test.
// Play: https://go.dev/play/p/y4gstlmx4KR
func FilterWithContext[T any](predicate func(ctx context.Context, item T) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FilterI emits only those items from an Observable that pass a predicate test.
// Play: https://go.dev/play/p/Y5a2-AicBWO
func FilterI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FilterIWithContext emits only those items from an Observable that pass a predicate test.
// Play: https://go.dev/play/p/xjz-pViifdB
func FilterIWithContext[T any](predicate func(ctx context.Context, item T, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Distinct suppresses duplicate items in an Observable.
// Play: https://go.dev/play/p/szxp8gO0_I7
func Distinct[T comparable]() func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DistinctBy suppresses duplicate items in an Observable based on a key selector.
func DistinctBy[T any, K comparable](keySelector func(item T) K) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DistinctByWithContext suppresses duplicate items in an Observable based on a key selector.
// The context is passed to the key selector function.
func DistinctByWithContext[T any, K comparable](keySelector func(ctx context.Context, item T) (context.Context, K)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IgnoreElements does not emit any items from an Observable but mirrors its
// termination notification. It is useful for ignoring all the items from an
// Observable but you want to be notified when it completes or when it throws an error.
// Play: https://go.dev/play/p/glDG6E-gZ1V
func IgnoreElements[T any]() func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Skip suppresses the first n items emitted by an Observable.
// If the count is greater than the number of items emitted by the source Observable,
// Skip will not emit any items. If the count is zero, Skip will emit all items.
// Play: https://go.dev/play/p/AAEJaUZJuIj
func Skip[T any](count int64) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SkipWhile skips items emitted by an Observable until a specified condition
// becomes false. It will then emit all the subsequent items. If the condition
// is never false, SkipWhile will not emit any items. If the condition is false
// on the first item, SkipWhile will emit all items.
func SkipWhile[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SkipWhileWithContext skips items emitted by an Observable until a specified condition
// becomes false. It will then emit all the subsequent items. If the condition
// is never false, SkipWhile will not emit any items. If the condition is false
// on the first item, SkipWhile will emit all items.
func SkipWhileWithContext[T any](predicate func(ctx context.Context, item T) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SkipWhileI skips items emitted by an Observable until a specified condition
// becomes false. It will then emit all the subsequent items. If the condition
// is never false, SkipWhile will not emit any items. If the condition is false
// on the first item, SkipWhile will emit all items.
func SkipWhileI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SkipWhileIWithContext skips items emitted by an Observable until a specified condition
// becomes false. It will then emit all the subsequent items. If the condition
// is never false, SkipWhile will not emit any items. If the condition is false
// on the first item, SkipWhile will emit all items.
// Play: https://go.dev/play/p/oYUQuPWIytL
func SkipWhileIWithContext[T any](predicate func(ctx context.Context, item T, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SkipLast suppresses the last n items emitted by an Observable. If the count
// is greater than the number of items emitted by the source Observable, SkipLast
// will not emit any items. If the count is zero, SkipLast will emit all items.
// Play: https://go.dev/play/p/gire30ONRBB
func SkipLast[T any](count int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Use a circular buffer approach to avoid memory allocations

// Buffer is full, emit the oldest item

// SkipUntil suppresses items emitted by an Observable until a second Observable
// emits an item or completes. It will then emit all the subsequent items. If the
// second Observable is empty, SkipUntil will not emit any items. If the second
// Observable emits an item or completes, SkipUntil will emit all items.
func SkipUntil[T, S any](signal Observable[S]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Take emits only the first n items emitted by an Observable. If the count is
// greater than the number of items emitted by the source Observable, Take will
// emit all items. If the count is zero, Take will not emit any items.
// Play: https://go.dev/play/p/IC_hJMsg7yk
func Take[T any](count int64) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Warning: the `source` will never be subscribed

// TakeWhile emits items emitted by an Observable so long as a specified condition
// is true. It will then complete. If the condition is never true, TakeWhile will
// not emit any items. If the condition is true on the first item, TakeWhile will
// emit all items. If the condition is false on the first item, TakeWhile will not
// emit any items.
// Play: https://go.dev/play/p/lxV03GzOa2J
func TakeWhile[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhileWithContext emits items emitted by an Observable so long as a specified condition
// is true. It will then complete. If the condition is never true, TakeWhile will
// not emit any items. If the condition is true on the first item, TakeWhile will
// emit all items. If the condition is false on the first item, TakeWhile will not
// emit any items.
func TakeWhileWithContext[T any](predicate func(ctx context.Context, item T) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhileI emits items emitted by an Observable so long as a specified condition
// is true. It will then complete. If the condition is never true, TakeWhile will
// not emit any items. If the condition is true on the first item, TakeWhile will
// emit all items. If the condition is false on the first item, TakeWhile will not
// emit any items.
func TakeWhileI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TakeWhileIWithContext emits items emitted by an Observable so long as a specified condition
// is true. It will then complete. If the condition is never true, TakeWhile will
// not emit any items. If the condition is true on the first item, TakeWhile will
// emit all items. If the condition is false on the first item, TakeWhile will not
// emit any items.
func TakeWhileIWithContext[T any](predicate func(ctx context.Context, item T, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TakeLast emits only the last n items emitted by an Observable. If the count is
// greater than the number of items emitted by the source Observable, TakeLast will
// emit all items. If the count is zero, TakeLast will not emit any items.
// Play: https://go.dev/play/p/J0mX3NpEHzy
func TakeLast[T any](count int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Warning: the `source` will never be subscribed

// Use a circular buffer to avoid memory allocations

// Emit items in order, starting from the oldest

// TakeUntil emits items emitted by an Observable until a second Observable emits
// an item or completes. It will then complete. If the second Observable is empty,
// TakeUntil will emit all items. If the second Observable emits an item or completes,
// TakeUntil will emit all items. If the second Observable emits an item or completes,
// TakeUntil will complete.
// Play: https://go.dev/play/p/nhgYGyREW1r
func TakeUntil[T, S any](signal Observable[S]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Head emits only the first item emitted by an Observable. If the source Observable
// is empty, Head will emit an error.
// Play: https://go.dev/play/p/TmhTvpuKAp_U
func Head[T any]() func(Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// Tail emits only the last item emitted by an Observable. If the source Observable
// is empty, Tail will emit an error.
func Tail[T any]() func(Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// First emits only the first item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, First will emit an error.
// Play: https://go.dev/play/p/yneVKit6vh0
func First[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FirstWithContext emits only the first item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, First will emit an error.
func FirstWithContext[T any](predicate func(ctx context.Context, item T) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FirstI emits only the first item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, FirstI will emit an error.
func FirstI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FirstIWithContext emits only the first item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, FirstI will emit an error.
func FirstIWithContext[T any](predicate func(ctx context.Context, item T, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Last emits only the last item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, Last will emit an error.
// Play: https://go.dev/play/p/aMsvsTPbmHY
func Last[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LastWithContext emits only the last item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, Last will emit an error.
func LastWithContext[T any](predicate func(ctx context.Context, item T) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LastI emits only the last item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, LastI will emit an error.
func LastI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LastIWithContext emits only the last item emitted by an Observable that satisfies a specified
// condition. If the source Observable is empty, LastI will emit an error.
func LastIWithContext[T any](predicate func(ctx context.Context, item T, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ElementAt emits only the nth item emitted by an Observable. If the source Observable
// emits fewer than n items, ElementAt will emit an error.
// Play: https://go.dev/play/p/0YE1tCbPaDg
func ElementAt[T any](nth int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ElementAtOrDefault emits only the nth item emitted by an Observable. If the source
// Observable emits fewer than n items, ElementAtOrDefault will emit a fallback value.
// Play: https://go.dev/play/p/DWMWPXkc8x4
func ElementAtOrDefault[T any](nth int64, fallback T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
