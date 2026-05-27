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
	"sync"
)

// Teardown is a function that cleans up resources, such as closing
// a file or a network connection. It is called when the Subscription is closed.
// It is part of a Subscription, and is returned by the Observable creation.
// It will be called only once, when the Subscription is canceled.
type Teardown func()

// Unsubscribable represents any type that can be unsubscribed from.
// It provides a common interface for cancellation operations.
type Unsubscribable interface {
	Unsubscribe()
}

// Subscription represents an ongoing execution of an `Observable`, and has
// a minimal API which allows you to cancel that execution.
type Subscription interface {
	Unsubscribable

	Add(teardown Teardown)
	AddUnsubscribable(unsubscribable Unsubscribable)
	IsClosed() bool
	Wait() // Note: using .Wait() is not recommended.
}

var _ Subscription = (*subscriptionImpl)(nil)

// NewSubscription creates a new Subscription. When `teardown` is nil, nothing
// is added. When the subscription is already disposed, the `teardown` callback
// is triggered immediately.
func NewSubscription(teardown Teardown) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Pre-allocate for common case

type subscriptionImpl struct {
	done       bool
	mu         sync.Mutex // Should be a RWMutex because of the .IsClosed() method, but sync.RWMutex is 30% slower.
	finalizers []func()
}

// Add receives a finalizer to execute upon unsubscription. When `teardown`
// is nil, nothing is added. When the subscription is already disposed, the `teardown`
// callback is triggered immediately.
//
// This method is thread-safe.
//
// Implements Subscription.
func (s *subscriptionImpl) Add(teardown Teardown) { _ = "STUB: not implemented"; return }

// not protected against panics

// AddUnsubscribable merges multiple subscriptions into one. The method does nothing
// if `unsubscribable` is nil.
//
// This method is thread-safe.
//
// Implements Subscription.
func (s *subscriptionImpl) AddUnsubscribable(unsubscribable Unsubscribable) {
	_ = "STUB: not implemented"
	return
}

// Unsubscribe disposes the resources held by the subscription. May, for
// instance, cancel an ongoing `Observable` execution or cancel any other
// type of work that started when the `Subscription` was created.
//
// This method is thread-safe. Finalizers are executed in sequence.
//
// Implements Unsuscribable.
func (s *subscriptionImpl) Unsubscribe() { _ = "STUB: not implemented"; return }

// Note: we prefer not running this in parallel.

// protected against panics

// OnUnhandledError(err)

// Error is triggered after the recursive call to finalizers
// because we want to execute all finalizers before panicking.

// errors.Join has been introduced in go 1.20

// IsClosed returns true if the subscription has been disposed
// or if unsubscription is in progress.
//
// Implements Subscription.
func (s *subscriptionImpl) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Wait blocks until a `Subscription` is canceled. It can be used for
// blocking until an `Observable` throws an error or completes.
//
// Please use it carefully. Calling this method is against the Reactive
// Programming Manifesto. This method might be deleted in the future.
//
// Note: using .Wait() is not recommended.
//
// Implements Subscription.
func (s *subscriptionImpl) Wait() { _ = "STUB: not implemented"; return }

// There is no guarantee that this callback will be the last finalizer
// added to this subscription.

// execFinalizer runs the finalizer and catches any panics, converting them to errors.
func execFinalizer(finalizer func()) (err error) { _ = "STUB: not implemented"; return nil }

// @TODO: Add methods Remove + RemoveSubscription.
// Currently, Go does not support function address comparison, so we cannot
// remove a finalizer from the list.
