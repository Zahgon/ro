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

// All determines whether all elements of an observable sequence satisfy a condition.
// Play: https://go.dev/play/p/t22F_crlA-l
func All[T any](predicate func(T) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// AllWithContext determines whether all elements of an observable sequence satisfy a condition.
// Play: https://go.dev/play/p/NEA7Zi7yVNh
func AllWithContext[T any](predicate func(ctx context.Context, item T) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// AllI determines whether all elements of an observable sequence satisfy a condition.
func AllI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// AllIWithContext determines whether all elements of an observable sequence satisfy a condition.
// Play: https://go.dev/play/p/UkOzE4wQXPG
func AllIWithContext[T any](predicate func(ctx context.Context, item T, index int64) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// Contains determines whether an observable sequence contains a specified element with an equality comparer.
// Play: https://go.dev/play/p/ldteqqGsMWM
func Contains[T any](predicate func(item T) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// ContainsWithContext determines whether an observable sequence contains a specified element with an equality comparer.
// Play: https://go.dev/play/p/RPHkyiLrFVW
func ContainsWithContext[T any](predicate func(ctx context.Context, item T) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// ContainsI determines whether an observable sequence contains a specified element with an equality comparer.
func ContainsI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// ContainsIWithContext determines whether an observable sequence contains a specified element with an equality comparer.
// Play: https://go.dev/play/p/TkLfujMVNJb
func ContainsIWithContext[T any](predicate func(ctx context.Context, item T, index int64) bool) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}

// Find returns the first element of an observable sequence that satisfies the condition.
// Play: https://go.dev/play/p/2f5rn0HoKeq
func Find[T any](predicate func(item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FindWithContext returns the first element of an observable sequence that satisfies the condition.
// Play: https://go.dev/play/p/BVm-Grgv11w
func FindWithContext[T any](predicate func(ctx context.Context, item T) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FindI returns the first element of an observable sequence that satisfies the condition.
func FindI[T any](predicate func(item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// FindIWithContext returns the first element of an observable sequence that satisfies the condition.
// Play: https://go.dev/play/p/X8oT_CF9IKM
func FindIWithContext[T any](predicate func(ctx context.Context, item T, index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// return zero value or error ?

// Iif determines which one of two observables to return based on a condition.
// Play: https://go.dev/play/p/t-sNgL5EZA-
func Iif[T any](predicate func() bool, source1, source2 Observable[T]) func() Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DefaultIfEmpty emits a default value if the source observable emits no items.
// Play: https://go.dev/play/p/WDh807OLPWv
func DefaultIfEmpty[T any](defaultValue T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DefaultIfEmptyWithContext emits a default value if the source observable emits no items.
func DefaultIfEmptyWithContext[T any](defaultCtx context.Context, defaultValue T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SequenceEqual determines whether two observable sequences are equal by comparing the elements pairwise.
// Play: https://go.dev/play/p/cBIQlH01byQ
func SequenceEqual[T comparable](obsB Observable[T]) func(Observable[T]) Observable[bool] {
	_ = "STUB: not implemented"
	return nil
}
