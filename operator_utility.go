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

// Tap allows you to perform side effects for notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/oDI3d6553MI
func Tap[T any](onNext func(value T), onError func(err error), onComplete func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapWithContext allows you to perform side effects for notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/oDI3d6553MI
func TapWithContext[T any](onNext func(ctx context.Context, value T), onError func(ctx context.Context, err error), onComplete func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Do is an alias to Tap.
// Play: https://go.dev/play/p/s_BSHgxdjUR
func Do[T any](onNext func(value T), onError func(err error), onComplete func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoWithContext is an alias to Tap.
func DoWithContext[T any](onNext func(ctx context.Context, value T), onError func(ctx context.Context, err error), onComplete func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnNext allows you to perform side effects for Next notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/oDI3d6553MI
func TapOnNext[T any](onNext func(value T)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnNextWithContext allows you to perform side effects for Next notifications from the source Observable
// Play: https://go.dev/play/p/oDI3d6553MI
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
func TapOnNextWithContext[T any](onNext func(ctx context.Context, value T)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnNext is an alias to TapOnNext.
func DoOnNext[T any](onNext func(value T)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil

	// DoOnNextWithContext is an alias to TapOnNextWithContext.
}

func DoOnNextWithContext[T any](onNext func(ctx context.Context, value T)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnError allows you to perform side effects for Error notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/oDI3d6553MI
func TapOnError[T any](onError func(err error)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnErrorWithContext allows you to perform side effects for Error notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
func TapOnErrorWithContext[T any](onError func(ctx context.Context, err error)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnError is an alias to TapOnError.
func DoOnError[T any](onError func(err error)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnErrorWithContext is an alias to TapOnErrorWithContext.
func DoOnErrorWithContext[T any](onError func(ctx context.Context, err error)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnComplete allows you to perform side effects for Complete notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/z1sntT6bplM
func TapOnComplete[T any](onComplete func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnCompleteWithContext allows you to perform side effects for Complete notifications from the source Observable
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/3k25j_D1OTW
func TapOnCompleteWithContext[T any](onComplete func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnComplete is an alias to TapOnComplete.
func DoOnComplete[T any](onComplete func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnCompleteWithContext is an alias to TapOnCompleteWithContext.
func DoOnCompleteWithContext[T any](onComplete func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnSubscribe allows you to perform side effects when the source Observable is subscribed to
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/0YzsxpRkO4T
func TapOnSubscribe[T any](onSubscribe func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnSubscribeWithContext allows you to perform side effects when the source Observable is subscribed to
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
func TapOnSubscribeWithContext[T any](onSubscribe func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// triggers before the source is subscribed

// DoOnSubscribe is an alias to TapOnSubscribe.
func DoOnSubscribe[T any](onSubscribe func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// DoOnSubscribeWithContext is an alias to TapOnSubscribe.
func DoOnSubscribeWithContext[T any](onSubscribe func(ctx context.Context)) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// TapOnFinalize allows you to perform side effects when the source Observable is unsubscribed from
// without modifying the emitted items. It mirrors the source Observable and forwards its emissions
// to the provided observer.
// Play: https://go.dev/play/p/VEACE_KhdvU
func TapOnFinalize[T any](onFinalize func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// triggers after the source is unsubscribed

// DoOnFinalize is an alias to TapOnFinalize.
// Play: https://go.dev/play/p/7en6T1q33WF
func DoOnFinalize[T any](onFinalize func()) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IntervalValue is a value emitted by the `TimeInterval` operator.
type IntervalValue[T any] struct {
	Value    T
	Interval time.Duration
}

// TimeInterval emits the values emitted by the source Observable with the time elapsed between each emission.
// Play: https://go.dev/play/p/VX73ZL74hPk
func TimeInterval[T any]() func(Observable[T]) Observable[IntervalValue[T]] {
	_ = "STUB: not implemented"
	return nil
}

// TimestampValue is a value emitted by the `TimeInterval` operator.
type TimestampValue[T any] struct {
	Value     T
	Timestamp time.Duration
}

// Timestamp emits the values emitted by the source Observable with the time elapsed since the source Observable was subscribed to.
// Play: https://go.dev/play/p/cDiCr6qIE2P
func Timestamp[T any]() func(Observable[T]) Observable[TimestampValue[T]] {
	_ = "STUB: not implemented"
	return nil
}

// Delay delays the emissions of the source Observable by a given duration without modifying the emitted items.
// It mirrors the source Observable and forwards its emissions to the provided observer.
// Error and Complete notifications are delayed as well.
//
// @TODO: set queue size ?
// Play: https://go.dev/play/p/K3md7WPtZGI
func Delay[T any](duration time.Duration) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Using time.AfterFunc is not convenient because it may introduce race
// conditions and/or change message order.
// We need a double mutex to prevent message reordering:
//   - one to protect the queue and allow pushing new values while we call destination.Next()
//   - one to protect the call to destination.Next() itself

// DelayEach delays the emissions of the source Observable by a given duration without modifying the emitted items.
// Play: https://go.dev/play/p/dReP7-bffEU
func DelayEach[T any](duration time.Duration) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// RepeatWith repeats the source Observable a specified number of times.
// This is a pipeable operator. The creation operator equivalent is `Repeat`.
//
// The destination is flatten.
// Play: https://go.dev/play/p/fEKtAX9_nYe
func RepeatWith[T any](count int64) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// might do nothing if already closed

// Timeout raises an error if the source Observable does not emit any item within the specified duration.
// Play: https://go.dev/play/p/t0xKoj-_AqZ
func Timeout[T any](duration time.Duration) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// if no value is emitted, we use the subscriber context

//nolint:errcheck,forcetypeassert

// @TODO: what happens if the above line is too slow?

// Materialize converts the source Observable into a stream of Notification instances.
// Play: https://go.dev/play/p/ZHtPviPoqWK
func Materialize[T any]() func(Observable[T]) Observable[Notification[T]] {
	_ = "STUB: not implemented"
	return nil
}

// Dematerialize converts the source Observable of Notification instances back into a stream of items.
// Play: https://go.dev/play/p/oRymdDqkh25
func Dematerialize[T any]() func(Observable[Notification[T]]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// SubscribeOn schedule the upstream flow to a different goroutine. Next, Error and Complete notifications
// are sent to a queue first, then the consumer consume this queue.
// SubscribeOn converts a push-based Observable into a pullable stream with backpressure capabilities.
//
// To schedule the downstream flow to a different goroutine, refer to SubscribeOn.
//
// When an Observable emits values faster than they can be consumed, SubscribeOn buffers these values
// in a queue of specified capacity. This allows downstream consumers to pull values at their own pace
// while managing backpressure from upstream emissions.
//
// Note: Once the buffer reaches its capacity, upstream emissions will block until space becomes
// available, effectively implementing backpressure control.
//
// @TODO: add a backpressure policy ? drop vs block.
// Play: https://go.dev/play/p/WrsTUq6yxtO
func SubscribeOn[T any](bufferSize int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ObserveOn schedule the downstream flow to a different goroutine. Next, Error and Complete notifications
// are sent to a queue first, then the consumer consume this queue.
// ObserveOn converts a push-based Observable into a pullable stream with backpressure capabilities.
//
// To schedule the upstream flow to a different goroutine, refer to SubscribeOn.
//
// When an Observable emits values faster than they can be consumed, ObserveOn buffers these values
// in a queue of specified capacity. This allows downstream consumers to pull values at their own pace
// while managing backpressure from upstream emissions.
//
// Note: Once the buffer reaches its capacity, upstream emissions will block until space becomes
// available, effectively implementing backpressure control.
//
// @TODO: add a backpressure policy ? drop vs block.
// Play: https://go.dev/play/p/BpdKJ6Mya03
func ObserveOn[T any](bufferSize int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func detachOn[T any](bufferSize int, onUpstream, onDownstream bool) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// The goroutine could be used either on producer or consumer side.
// 	* ObserveOn moves the goroutine on the consumer side.
// 	* SubscribeOn moves the goroutine on the producer side.

// Serialize ensures thread-safe message passing by wrapping any observable in a ro.SafeObservable implementation.
func Serialize[T any]() func(Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }
