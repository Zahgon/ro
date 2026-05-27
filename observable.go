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
	"sync"
)

// Backpressure is a type that represents the backpressure strategy to use.
type Backpressure int8

const (
	// BackpressureBlock blocks the source observable when the destination is not ready to receive more values.
	BackpressureBlock Backpressure = iota
	// BackpressureDrop drops the source observable when the destination is not ready to receive more values.
	BackpressureDrop
)

// ConcurrencyMode is a type that represents the concurrency mode to use.
type ConcurrencyMode int8

// ConcurrencyModeSafe is a concurrency mode that is safe to use.
// Spinlock is ignored because it is too slow when chaining operators. Spinlock should be used
// only for short-lived local locks.
const (
	ConcurrencyModeSafe ConcurrencyMode = iota
	ConcurrencyModeUnsafe
	ConcurrencyModeEventuallySafe
)

// Observable is the producer of values. It is the source of values that are
// emitted to Observers.
// Observable is a representation of any set of values over any amount of time.
//
// The primary method of an Observable is subscribe, which is used to attach an
// Observer to the Observable. Once an Observer is subscribed, the Observable
// may begin to emit items to the Observer. An Observable may emit any number
// of items (including zero items), then may either complete or error, but not
// both. Upon completion or error, the Observable will not emit any more items.
//
// An Observable may call an Observer's methods synchronously or asynchronously.
//
// An Observable is not a stream. It is a factory for streams.
type Observable[T any] interface {
	// Subscribe subscribes an Observer to the Observable. The Observer will begin
	// to receive items emitted by the Observable. The Observer may receive any
	// number of items (including zero items), then may either complete or error,
	// but not both. Upon completion or error, the Observer will not receive any
	// more items.
	//
	// The Subscribe method returns a Subscription that can be used to unsubscribe
	// the Observer from the Observable. The Subscription may be used to cancel the
	// subscription, and to wait for the subscription to complete.
	//
	// The Subscription might be already disposed when the Subscribe method returns.
	// In this case, the Teardown function is not called.
	//
	// The Subscribe method may call the Observer's methods synchronously or
	// asynchronously. The Observer is responsible for handling concurrency and
	// synchronization.
	Subscribe(destination Observer[T]) Subscription
	SubscribeWithContext(ctx context.Context, destination Observer[T]) Subscription
}

var _ Observable[int] = (*observableImpl[int])(nil)

// NewObservable creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewObservable[T any](subscribe func(destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewSafeObservable creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewSafeObservable[T any](subscribe func(destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewUnsafeObservable creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewUnsafeObservable[T any](subscribe func(destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewEventuallySafeObservable creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is safe for concurrent use, but concurrent messages are dropped.
func NewEventuallySafeObservable[T any](subscribe func(destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewObservableWithContext creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewObservableWithContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewSafeObservableWithContext creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewSafeObservableWithContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewUnsafeObservableWithContext creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is not safe for concurrent use.
func NewUnsafeObservableWithContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewEventuallySafeObservableWithContext creates a new Observable. The subscribe function is called when
// the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error,
// but not both. Upon completion or error, the Observable will not emit any more
// items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// This method is safe for concurrent use, but concurrent messages are dropped.
func NewEventuallySafeObservableWithContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewObservableWithConcurrencyMode creates a new Observable with the given concurrency mode.
// The subscribe function is called when the Observable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error, but not both. Upon completion or error, the Observable will not emit any more items.
//
// The subscribe function should return a Teardown function that will be called
// when the Subscription is unsubscribed. The Teardown function should clean up
// any resources created during the subscription.
//
// The subscribe function may return a Teardown function that does nothing, if
// no cleanup is necessary. In this case, the Teardown function should return nil.
//
// The Observable will use the given concurrency mode.
//
// It is rarely used as a public API.
func NewObservableWithConcurrencyMode[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown, mode ConcurrencyMode) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

type observableImpl[T any] struct {
	mode      ConcurrencyMode
	subscribe func(ctx context.Context, destination Observer[T]) Teardown
}

// Subscribe subscribes an Observer to the Observable. The Observer will begin
// to receive items emitted by the Observable. The Observer may receive any
// number of items (including zero items), then may either complete or error,
// but not both. Upon completion or error, the Observer will not receive any
// more items.
//
// The Subscribe method returns a Subscription that can be used to unsubscribe
// the Observer from the Observable. The Subscription may be used to cancel the
// subscription, and to wait for the subscription to complete.
//
// The Subscription might be already disposed when the Subscribe method returns.
// In this case, the Teardown function is not called.
//
// The Subscribe method may call the Observer's methods synchronously or
// asynchronously. The Observer is responsible for handling concurrency and
// synchronization.
func (s *observableImpl[T]) Subscribe(destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// SubscribeWithContext subscribes an Observer to the Observable. The Observer will begin
// to receive items emitted by the Observable. The Observer may receive any
// number of items (including zero items), then may either complete or error,
// but not both. Upon completion or error, the Observer will not receive any
// more items.
//
// The Subscribe method returns a Subscription that can be used to unsubscribe
// the Observer from the Observable. The Subscription may be used to cancel the
// subscription, and to wait for the subscription to complete.
//
// The Subscription might be already disposed when the Subscribe method returns.
// In this case, the Teardown function is not called.
//
// The Subscribe method may call the Observer's methods synchronously or
// asynchronously. The Observer is responsible for handling concurrency and
// synchronization.
func (s *observableImpl[T]) SubscribeWithContext(ctx context.Context, destination Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// Warning: here, we are catching panic in subscription.Add.
// I'm not sure if it's a good idea.

// Collect collects all values emitted by the source Observable and returns them
// as a slice. It waits for the source Observable to complete before returning.
// If the source Observable emits an error, the error is returned along with the
// values collected so far.
func Collect[T any](obs Observable[T]) ([]T, error) { _ = "STUB: not implemented"; return nil, nil }

// CollectWithContext collects all values emitted by the source Observable and returns them
// as a slice. It waits for the source Observable to complete before returning.
// If the source Observable emits an error, the error is returned along with the
// values collected so far.
// @TODO: return more values, such as (isCanceled bool) or (duration time.Duration) ?
func CollectWithContext[T any](ctx context.Context, obs Observable[T]) ([]T, context.Context, error) {
	_ = "STUB: not implemented"
	return nil, *new(context.Context), nil
}

// Note: using .Wait() is not recommended.

// ConnectableObservable is an Observable that can be connected and disconnected.
// When connected, it will emit values to its observers.
//
// ConnectableObservable is useful when you want to share a single subscription to an Observable
// among multiple observers. This is useful when you want to multicast the values of an Observable.
type ConnectableObservable[T any] interface {
	Observable[T]

	// Connect connects the ConnectableObservable. When connected, the ConnectableObservable
	// will emit values to its observers. If the ConnectableObservable is already connected,
	// this method creates a new subscription and starts emitting values to its observers.
	//
	// The Connect method returns a Subscription that can be used to disconnect the
	// ConnectableObservable. The Subscription may be used to cancel the connection,
	// and to wait for the connection to complete.
	//
	// The Subscription might be already disposed when the Connect method returns.
	Connect() Subscription
	ConnectWithContext(ctx context.Context) Subscription
}

var (
	_ ConnectableObservable[int] = (*connectableObservableImpl[int])(nil)
	_ Observable[int]            = (*connectableObservableImpl[int])(nil)
)

// ConnectableConfig is the configuration for a ConnectableObservable.
type ConnectableConfig[T any] struct {
	Connector         func() Subject[T]
	ResetOnDisconnect bool
}

func defaultConnector[T any]() Subject[T] { _ = "STUB: not implemented"; return nil }

// NewConnectableObservable creates a new ConnectableObservable. The subscribe function is called when
// the ConnectableObservable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error, but not both.
// Upon completion or error, the ConnectableObservable will not emit any more items.
//
// The ConnectableObservable will use the default connector, which is a PublishSubject.
// The ConnectableObservable will reset the source when disconnected. This means that
// when the ConnectableObservable is disconnected, it will create a new source when
// reconnected.
//
// If you want to use a different connector or change the reset behavior, use
// NewConnectableObservableWithConfig.
func NewConnectableObservable[T any](subscribe func(destination Observer[T]) Teardown) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewConnectableObservableWithContext creates a new ConnectableObservable. The subscribe function is called when
// the ConnectableObservable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error, but not both.
// Upon completion or error, the ConnectableObservable will not emit any more items.
//
// The ConnectableObservable will use the default connector, which is a PublishSubject.
// The ConnectableObservable will reset the source when disconnected. This means that
// when the ConnectableObservable is disconnected, it will create a new source when
// reconnected.
//
// If you want to use a different connector or change the reset behavior, use
// NewConnectableObservableWithConfig.
func NewConnectableObservableWithContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewConnectableObservableWithConfig creates a new ConnectableObservable. The subscribe function is called when
// the ConnectableObservable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error, but not both.
// Upon completion or error, the ConnectableObservable will not emit any more items.
//
// The ConnectableObservable will use the given connector. The ConnectableObservable will reset
// the source when disconnected if ResetOnDisconnect is true. This means that when the
// ConnectableObservable is disconnected, it will create a new source when reconnected.
func NewConnectableObservableWithConfig[T any](subscribe func(destination Observer[T]) Teardown, config ConnectableConfig[T]) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

// NewConnectableObservableWithConfigAndContext creates a new ConnectableObservable. The subscribe function is called when
// the ConnectableObservable is subscribed to. The subscribe function is given an Observer,
// to which it may emit any number of items, then may either complete or error, but not both.
// Upon completion or error, the ConnectableObservable will not emit any more items.
//
// The ConnectableObservable will use the given connector. The ConnectableObservable will reset
// the source when disconnected if ResetOnDisconnect is true. This means that when the
// ConnectableObservable is disconnected, it will create a new source when reconnected.
func NewConnectableObservableWithConfigAndContext[T any](subscribe func(ctx context.Context, destination Observer[T]) Teardown, config ConnectableConfig[T]) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Connectable creates a new ConnectableObservable from an Observable. The ConnectableObservable
// will use the default connector, which is a PublishSubject. The ConnectableObservable will reset
// the source when disconnected. This means that when the ConnectableObservable is disconnected,
// it will create a new source when reconnected.
//
// If you want to use a different connector or change the reset behavior, use ConnectableWithConfig.
func Connectable[T any](source Observable[T]) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ConnectableWithConfig creates a new ConnectableObservable from an Observable. The ConnectableObservable
// will use the given connector. The ConnectableObservable will reset the source when disconnected
// if ResetOnDisconnect is true. This means that when the ConnectableObservable is disconnected,
// it will create a new source when reconnected.
func ConnectableWithConfig[T any](source Observable[T], config ConnectableConfig[T]) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

func newConnectableObservableImpl[T any](source Observable[T], config ConnectableConfig[T]) ConnectableObservable[T] {
	_ = "STUB: not implemented"
	return nil
}

type connectableObservableImpl[T any] struct {
	mu           sync.Mutex
	config       ConnectableConfig[T]
	source       Observable[T]
	subject      Subject[T]
	subscription Subscription
}

// Connect connects the ConnectableObservable. When connected, the ConnectableObservable
// will emit values to its observers. If the ConnectableObservable is already connected,
// this method creates a new subscription and starts emitting values to its observers.
//
// The Connect method returns a Subscription that can be used to disconnect the
// ConnectableObservable. The Subscription may be used to cancel the connection,
// and to wait for the connection to complete.
//
// The Subscription might be already disposed when the Connect method returns.
func (s *connectableObservableImpl[T]) Connect() Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

// ConnectWithContext connects the ConnectableObservable. When connected, the ConnectableObservable
// will emit values to its observers. If the ConnectableObservable is already connected,
// this method creates a new subscription and starts emitting values to its observers.
//
// The Connect method returns a Subscription that can be used to disconnect the
// ConnectableObservable. The Subscription may be used to cancel the connection,
// and to wait for the connection to complete.
//
// The Subscription might be already disposed when the Connect method returns.
func (s *connectableObservableImpl[T]) ConnectWithContext(ctx context.Context) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

func (s *connectableObservableImpl[T]) Subscribe(observer Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}

func (s *connectableObservableImpl[T]) SubscribeWithContext(ctx context.Context, observer Observer[T]) Subscription {
	_ = "STUB: not implemented"
	return *new(Subscription)
}
