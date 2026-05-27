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
	"fmt"
)

var (
	// By default, the library will ignore unhandled errors and dropped notifications.
	// You can change this behavior by setting the following variables to your own
	// error handling functions.
	//
	// Example:
	//
	// 	ro.OnUnhandledError = func(ctx context.Context, err error) {
	// 		slog.Error(fmt.Sprintf("unhandled error: %s\n", err.Error()))
	// 	}
	//
	// 	ro.OnDroppedNotification = func(ctx context.Context, notification fmt.Stringer) {
	// 		slog.Warn(fmt.Sprintf("dropped notification: %s\n", notification.String()))
	// 	}
	//
	// Note: `OnUnhandledError` and `OnDroppedNotification` are called synchronously from
	// the goroutine that emits the error or the notification. A slow callback will slow
	// down the whole pipeline.

	// OnUnhandledError is called when an error is emitted by an Observable and
	// no error handler is registered.
	OnUnhandledError = IgnoreOnUnhandledError
	// OnDroppedNotification is called when a notification is emitted by an Observable and
	// no notification handler is registered.
	OnDroppedNotification = IgnoreOnDroppedNotification
)

// IgnoreOnUnhandledError is the default implementation of `OnUnhandledError`.
func IgnoreOnUnhandledError(ctx context.Context, err error) {
	_ = "STUB: not implemented"

	// IgnoreOnDroppedNotification is the default implementation of `OnDroppedNotification`.
	return
}

func IgnoreOnDroppedNotification(ctx context.Context, notification fmt.Stringer) {
	_ = "STUB: not implemented"

	// DefaultOnUnhandledError is the default implementation of `OnUnhandledError`.
	return
}

func DefaultOnUnhandledError(ctx context.Context, err error) {
	_ = "STUB: not implemented"

	// bearer:disable go_lang_logger_leak
	return
}

var _ fmt.Stringer = (*Notification[int])(nil) // see below

// DefaultOnDroppedNotification is the default implementation of `OnDroppedNotification`.
//
// Since we cannot assign a generic callback to `OnDroppedNotification`,
// we had to use a `fmt.Stringer` instead a `Notification[T any]`.
func DefaultOnDroppedNotification(ctx context.Context, notification fmt.Stringer) {
	_ = "STUB: not implemented"
	// bearer:disable go_lang_logger_leak
	return
}

// Kind represents the kind of a Notification.
// It can be Next, Error, or Complete.
type Kind uint8

// String returns the string representation of a Kind.
func (k Kind) String() string { _ = "STUB: not implemented"; return "" }

// Kind constants.
const (
	KindNext Kind = iota
	KindError
	KindComplete
)

// Notification represents a value emitted by an Observable. It can be a Next
// value, an Error, or a Complete signal. It is used to communicate between
// Observables and Observers. It is a generic type, so it can hold any value.
type Notification[T any] struct {
	Kind  Kind
	Value T
	Err   error
}

func (n Notification[T]) String() string { _ = "STUB: not implemented"; return "" }

// NewNotificationNext creates a new Notification with a Next value.
func NewNotificationNext[T any](value T) Notification[T] { _ = "STUB: not implemented"; return nil }

// NewNotificationError creates a new Notification with an Error.
func NewNotificationError[T any](err error) Notification[T] { _ = "STUB: not implemented"; return nil }

// NewNotificationComplete creates a new Notification with a Complete signal.
func NewNotificationComplete[T any]() Notification[T] { _ = "STUB: not implemented"; return nil }

func processNotification[T any](n Notification[T], onNext func(T), onError func(error), onComplete func()) bool {
	_ = "STUB: not implemented"
	return false
}

func processNotificationWithContext[T any](ctx context.Context, n Notification[T], onNext func(context.Context, T), onError func(context.Context, error), onComplete func(context.Context)) bool {
	_ = "STUB: not implemented"
	return false
}

func processNotificationWithObserver[T any](n Notification[T], destination Observer[T]) bool {
	_ = "STUB: not implemented"
	return false
}

func processNotificationWithObserverAndContext[T any](ctx context.Context, n Notification[T], destination Observer[T]) bool {
	_ = "STUB: not implemented"
	return false
}
