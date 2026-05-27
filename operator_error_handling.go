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
	"time"
)

// Catch catches errors on the observable to be handled by returning a new observable
// or throwing an error.
// Play: https://go.dev/play/p/0pVlxwjhdMT
func Catch[T any](finally func(err error) Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// OnErrorResumeNextWith instructs an Observable to begin emitting a second
// Observable sequence if it encounters an error or completes. It immediately
// subscribes to the next one that was passed.
// Play: https://go.dev/play/p/9XLTAOginbK
func OnErrorResumeNextWith[T any](finally ...Observable[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// `subscriptions` cancels `sub` when it unsubscribes
// but `sub` cannot unsubscribe `subscriptions`

// OnErrorReturn instructs an Observable to emit a particular item when it
// encounters an error. It will then complete the sequence.
// Play: https://go.dev/play/p/d_9xe1oedjU
func OnErrorReturn[T any](finally T) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Retry resubscribes to the source observable when it encounters an error.
// It will retry infinitely. If you want to limit the number of retries, use
// RetryWithConfig.
// Play: https://go.dev/play/p/Llj9dT9Y3Z2
func Retry[T any]() func(Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// unlimited
// disabled
// disabled because it retries infinitely

// RetryConfig is the configuration for the Retry operator.
type RetryConfig struct {
	MaxRetries     uint64
	Delay          time.Duration
	ResetOnSuccess bool
}

// RetryWithConfig resubscribes to the source observable when it encounters
// an error. If a max number of retries is set, it will retry until the max
// number of retries is reached. If a delay is set, it will wait before retrying.
// If resetOnSuccess is set, it will reset the number of retries when a value is
// emitted.
// Play: https://go.dev/play/p/GilWi5xG0lr
func RetryWithConfig[T any](opts RetryConfig) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Check for context cancellation before retrying

// Use context-aware sleep that can be cancelled

// Continue to next iteration

// Continue to next iteration

// ThrowIfEmpty throws an error if the source observable is empty. It will
// throw the error returned by the throw function. If the source observable
// emits a value, it will complete. If the source observable emits an error,
// it will propagate the error.
// Play: https://go.dev/play/p/mLCaC7p_6p4
func ThrowIfEmpty[T any](throw func() error) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoWhile repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/nEWabaItDpn
func DoWhile[T any](condition func() bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoWhileWithContext repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
func DoWhileWithContext[T any](condition func(ctx context.Context) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoWhileI repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/cxOA9gimkCq
func DoWhileI[T any](condition func(index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoWhileIWithContext repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/yMoCCnnvRRH
func DoWhileIWithContext[T any](condition func(ctx context.Context, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Source emitted an error, stop the loop

// Condition is false, stop the loop

// While repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/hMj3DBVtp73
func While[T any](condition func() bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// WhileWithContext repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
func WhileWithContext[T any](condition func(ctx context.Context) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// WhileI repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/9aAuzAspyMc
func WhileI[T any](condition func(index int64) bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// WhileIWithContext repeats the source observable while the condition is true. It will
// complete when the condition is false. It will not emit any values if the
// source observable is empty. It will not emit any values if the source observable
// emits an error.
// Play: https://go.dev/play/p/xTpqdGSxOxw
func WhileIWithContext[T any](condition func(ctx context.Context, index int64) (context.Context, bool)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Condition is false, stop the loop

// Source completed normally

// Source emitted an error, stop the loop
