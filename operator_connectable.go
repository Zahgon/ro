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

// ShareConfig is the configuration for the Share operator.
type ShareConfig[T any] struct {
	Connector           func() Subject[T]
	ResetOnError        bool
	ResetOnComplete     bool
	ResetOnRefCountZero bool
}

// Share creates a new Observable that multicasts (shares) the original
// Observable. As long as there is at least one subscription to the
// multicasted Observable, the source Observable will be subscribed and
// emitting data. When all subscribers have unsubscribed, the source
// Observable will be unsubscribed.
//
// This is an alias for ShareWithConfig with default configuration.
// Play: https://go.dev/play/p/C34fv02jAIH
func Share[T any]() func(Observable[T]) Observable[T] { _ = "STUB: not implemented"; return nil }

// ShareWithConfig creates a new Observable that multicasts (shares) the
// original Observable. As long as there is at least one subscription to the
// multicasted Observable, the source Observable will be subscribed and
// emitting data. When all subscribers have unsubscribed, the source
// Observable will be unsubscribed.
//
// The configuration allows to customize the behavior of the shared
// Observable:
//   - `Connector` is a factory function that creates a new Subject for each
//     subscription. The Subject can be any type of Subject, such as a
//     ReplaySubject, a BehaviorSubject, a ReplaySubject, etc.
//   - `ResetOnError` determines whether the shared Observable should be reset
//     when an error is emitted.
//   - `ResetOnComplete` determines whether the shared Observable should be reset
//     when it completes.
//   - `ResetOnRefCountZero` determines whether the shared Observable should be
//     reset when the reference count reaches zero.
//
// Play: https://go.dev/play/p/C34fv02jAIH
func ShareWithConfig[T any](config ShareConfig[T]) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Subscriptions to `source` can be concurrent, so we protect shared
// objects against race conditions.

// var subject atomic.Pointer[Subject[T]]

// subscription between the source and the subject

// not an atomic counter, because it is protected by mutex

// atomic.Bool is not available in Go 1.18
// atomic.Bool is not available in Go 1.18

// Unsafe: must be called in a mutex lock.

// Unsafe: must be called in a mutex lock.

// never nil

// `currentSubject` is a backup (local reference) of `subject`
// to manipulate it even after reset.

// Expected to be non-blocking.
// This is the subscription between the subject and the new observer.

// We need to handle errors and completion so we added a
// proxy observer between source and subject.

// Subscription between the source and the subject.

// ShareReplayConfig is the configuration for the ShareReplay operator.
type ShareReplayConfig struct {
	ResetOnRefCountZero bool
}

// ShareReplay creates a new Observable that multicasts (shares) the original
// Observable and replays a specified number of items to any future
// subscribers. As long as there is at least one subscription to the
// multicasted Observable, the source Observable will be subscribed and
// emitting data. When all subscribers have unsubscribed, the source
// Observable will be unsubscribed.
//
// This is an alias for ShareReplayWithConfig with default configuration.
// Play: https://go.dev/play/p/QmsDbChzRgu
func ShareReplay[T any](bufferSize int) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ShareReplayWithConfig creates a new Observable that multicasts (shares) the
// original Observable and replays a specified number of items to any future
// subscribers. As long as there is at least one subscription to the
// multicasted Observable, the source Observable will be subscribed and
// emitting data. When all subscribers have unsubscribed, the source
// Observable will be unsubscribed.
//
// The configuration allows to customize the behavior of the shared
// Observable:
//   - `bufferSize` is the number of items to replay to future subscribers.
//   - `ResetOnRefCountZero` determines whether the shared Observable should be
//     reset when the reference count reaches zero.
//
// Play: https://go.dev/play/p/QmsDbChzRgu
func ShareReplayWithConfig[T any](bufferSize int, config ShareReplayConfig) func(Observable[T]) Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
