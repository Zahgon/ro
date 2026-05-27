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

// Observer is the consumer of an Observable. It receives notifications: Next,
// Error, and Complete. Observers are safe for concurrent calls to Next,
// Error, and Complete. It is the responsibility of the Observer to ensure
// that notifications are not forwarded after it has been closed.
type Observer[T any] interface {
	// Next receives the next value from the Observable. It is called zero or
	// more times by the Observable. The Observable may call Next synchronously
	// or asynchronously. If Next is called after the Observer has been closed,
	// the value will be dropped.
	Next(value T)
	NextWithContext(ctx context.Context, value T)
	// Error receives an error from the Observable. It is called at most once by
	// the Observable. The Observable may call Error synchronously or
	// asynchronously. If Error is called after the Observer has been closed, the
	// error will be dropped.
	Error(err error)
	ErrorWithContext(ctx context.Context, err error)
	// Complete receives a completion notification from the Observable. It is called
	// at most once by the Observable. The Observable may call Complete
	// synchronously or asynchronously. If Complete is called after the Observer has
	// been closed, the completion notification will be dropped.
	Complete()
	CompleteWithContext(ctx context.Context)

	// IsClosed returns true if the Observer has been closed, either by an error
	// or completion notification. If the Observer is closed, it will not receive
	// any more notifications.
	IsClosed() bool
	// HasThrown returns true if the Observer has received an error notification.
	HasThrown() bool
	// IsCompleted returns true if the Observer has received a completion notification.
	IsCompleted() bool
}

/************************
 *     Base Observer    *
 ************************/

var _ Observer[int] = (*observerImpl[int])(nil)

// NewObserver creates a new Observer with the provided callbacks. No context
// is provided.
func NewObserver[T any](onNext func(value T), onError func(err error), onComplete func()) Observer[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewObserverWithContext creates a new Observer with the provided callbacks. A context
// is provided to each callback.
func NewObserverWithContext[T any](onNext func(ctx context.Context, value T), onError func(ctx context.Context, err error), onComplete func(ctx context.Context)) Observer[T] {
	_ = "STUB: not implemented"
	return nil
}

type observerImpl[T any] struct {
	// 0: active
	// 1: errored
	// 2: completed
	status     int32
	onNext     func(context.Context, T)
	onError    func(context.Context, error) // @TODO: add a default onError that log the error ?
	onComplete func(context.Context)
}

func (o *observerImpl[T]) Next(value T) { _ = "STUB: not implemented"; return }

func (o *observerImpl[T]) NextWithContext(ctx context.Context, value T) {
	_ = "STUB: not implemented"
	return
}

func (o *observerImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

func (o *observerImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

func (o *observerImpl[T]) Complete() { _ = "STUB: not implemented"; return }

func (o *observerImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (o *observerImpl[T]) tryNext(ctx context.Context, value T) { _ = "STUB: not implemented"; return }

func (o *observerImpl[T]) tryError(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

func (o *observerImpl[T]) tryComplete(ctx context.Context) { _ = "STUB: not implemented"; return }

func (o *observerImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

func (o *observerImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

func (o *observerImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

/*********************
 * Partial Observers *
 *********************/

// OnNext is a partial Observer with only the Next method implemented.
// Warning: This observer will silent errors.
func OnNext[T any](onNext func(value T)) Observer[T] { _ = "STUB: not implemented"; return nil }

// OnNextWithContext is a partial Observer with only the Next method implemented.
// Warning: This observer will silent errors.
func OnNextWithContext[T any](onNext func(ctx context.Context, value T)) Observer[T] {
	_ = "STUB: not implemented"
	return nil
}

// OnError is a partial Observer with only the Error method implemented.
func OnError[T any](onError func(err error)) Observer[T] { _ = "STUB: not implemented"; return nil }

// OnErrorWithContext is a partial Observer with only the Error method implemented.
func OnErrorWithContext[T any](onError func(ctx context.Context, err error)) Observer[T] {
	_ = "STUB: not implemented"
	return nil
}

// OnComplete is a partial Observer with only the Complete method implemented.
// Warning: This observer will silent errors.
func OnComplete[T any](onComplete func()) Observer[T] { _ = "STUB: not implemented"; return nil }

// OnCompleteWithContext is a partial Observer with only the Complete method implemented.
// Warning: This observer will silent errors.
func OnCompleteWithContext[T any](onComplete func(ctx context.Context)) Observer[T] {
	_ = "STUB: not implemented"
	return nil
}

// NoopObserver is an Observer that does nothing.
// Warning: This observer will silent errors.
func NoopObserver[T any]() Observer[T] { _ = "STUB: not implemented"; return nil }

// PrintObserver is an utilitary Observer that dump notifications for debug purpose.
func PrintObserver[T any]() Observer[T] { _ = "STUB: not implemented"; return nil }
