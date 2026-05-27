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

	"github.com/samber/ro/internal/xsync"
)

// Subscriber implements the Observer and Subscription interfaces. While the Observer is
// the public API for consuming the values of an Observable, all Observers get
// converted to a Subscriber, in order to provide Subscription-like capabilities
// such as `Unsubscribe()`. Subscriber is a common type in samber/ro, and crucial for
// implementing operators, but it is rarely used as a public API.
type Subscriber[T any] interface {
	Subscription
	Observer[T]
}

var _ Subscriber[int] = (*subscriberImpl[int])(nil)

// NewSubscriber creates a new Subscriber from an Observer. If the Observer
// is already a Subscriber, it is returned as is. Otherwise, a new Subscriber
// is created that wraps the Observer.
//
// The returned Subscriber will unsubscribe from the destination Observer when
// Unsubscribe() is called.
//
// This method is safe for concurrent use.
//
// It is rarely used as a public API.
func NewSubscriber[T any](destination Observer[T]) Subscriber[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewSafeSubscriber creates a new Subscriber from an Observer. If the Observer
// is already a Subscriber, it is returned as is. Otherwise, a new Subscriber
// is created that wraps the Observer.
//
// The returned Subscriber will unsubscribe from the destination Observer when
// Unsubscribe() is called.
//
// This method is safe for concurrent use.
//
// It is rarely used as a public API.
func NewSafeSubscriber[T any](destination Observer[T]) Subscriber[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewUnsafeSubscriber creates a new Subscriber from an Observer. If the Observer
// is already a Subscriber, it is returned as is. Otherwise, a new Subscriber
// is created that wraps the Observer.
//
// The returned Subscriber will unsubscribe from the destination Observer when
// Unsubscribe() is called.
//
// This method is not safe for concurrent use.
//
// It is rarely used as a public API.
func NewUnsafeSubscriber[T any](destination Observer[T]) Subscriber[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewEventuallySafeSubscriber creates a new Subscriber from an Observer. If the Observer
// is already a Subscriber, it is returned as is. Otherwise, a new Subscriber
// is created that wraps the Observer.
//
// The returned Subscriber will unsubscribe from the destination Observer when
// Unsubscribe() is called.
//
// This method is safe for concurrent use, but concurrent messages are dropped.
//
// It is rarely used as a public API.
func NewEventuallySafeSubscriber[T any](destination Observer[T]) Subscriber[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewSubscriberWithConcurrencyMode creates a new Subscriber from an Observer. If the Observer
// is already a Subscriber, it is returned as is. Otherwise, a new Subscriber
// is created that wraps the Observer.
//
// The returned Subscriber will unsubscribe from the destination Observer when
// Unsubscribe() is called.
//
// It is rarely used as a public API.
func NewSubscriberWithConcurrencyMode[T any](destination Observer[T], mode ConcurrencyMode) Subscriber[T] {
	_ = "STUB: not implemented"
	// Spinlock is ignored because it is too slow when chaining operators. Spinlock should be used
	// only for short-lived local locks.
	return nil
}

// newSubscriberImpl creates a new subscriber implementation with the specified
// synchronization behavior and destination observer.
func newSubscriberImpl[T any](mode ConcurrencyMode, mu xsync.Mutex, backpressure Backpressure, destination Observer[T]) Subscriber[T] {
	_ = "STUB: not implemented"
	// Protect against multiple encapsulation layers.
	return nil
}

// KindNext

type subscriberImpl[T any] struct {
	// While mutex is used for synchronization of producer, status is used for storing state of
	// the subscriber. Using the mutex for reading the status would have create a dead lock if
	// an Observer calls Unsubscribe(), IsClosed(), HasThrown(), IsCompleted() synchronously.
	//
	// 0 - KindNext
	// 1 - KindError
	// 2 - KindComplete
	status       int32
	backpressure Backpressure

	_ [59]byte // padding to prevent false sharing

	// Mutex are much much faster than channels.
	//
	// Also, generators has been added in go1.23. A different implem of Observable/Observer
	// might reduce latency induced by mutexes.
	//
	// It could be interesting to implement a lock-free version of this,
	// with message drop instead of backpressure, and when SLO must be kept under
	// control (real-time streams?).
	mu          xsync.Mutex
	destination Observer[T]

	Subscription

	mode ConcurrencyMode
}

// Implements Observer.
func (s *subscriberImpl[T]) Next(v T) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *subscriberImpl[T]) NextWithContext(ctx context.Context, v T) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *subscriberImpl[T]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *subscriberImpl[T]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *subscriberImpl[T]) Complete() { _ = "STUB: not implemented"; return }

// Implements Observer.
func (s *subscriberImpl[T]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Implements Observer.
func (s *subscriberImpl[T]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *subscriberImpl[T]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *subscriberImpl[T]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

// Implements Observer.
func (s *subscriberImpl[T]) Unsubscribe() { _ = "STUB: not implemented"; return }

func (s *subscriberImpl[T]) unsubscribe() {
	_ = "STUB: not implemented"
	// s.Subscription.Unsubscribe() is protected against concurrent calls.
	return
}
