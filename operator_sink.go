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

// ToSlice collects all items from the observable into a slice. It is a sink
// operator so it emit a single value. It emits the slice when the source
// completes. If the source is empty, it emits an empty slice.
// Play: https://go.dev/play/p/kxbU_PzpN6t
func ToSlice[T any]() func(Observable[T]) Observable[[]T] { _ = "STUB: not implemented"; return nil }

// @TODO: use the context.Context from the last Next notification ?

// ToMap collects all items from the observable into a map. It is a sink
// operator so it emit a single value. It emits the map when the source
// completes. If the source is empty, it emits an empty map.
// Play: https://go.dev/play/p/FiF83XYB0ba
func ToMap[T any, K comparable, V any](project func(item T) (K, V)) func(Observable[T]) Observable[map[K]V] {
	_ = "STUB: not implemented"
	return nil
}

// ToMapWithContext collects all items from the observable into a map. It is a sink
// operator so it emit a single value. It emits the map when the source
// completes. If the source is empty, it emits an empty map.
// Play: https://go.dev/play/p/FiF83XYB0ba
func ToMapWithContext[T any, K comparable, V any](project func(ctx context.Context, item T) (K, V)) func(Observable[T]) Observable[map[K]V] {
	_ = "STUB: not implemented"
	return nil
}

// ToMapI collects all items from the observable into a map. It is a sink
// operator so it emit a single value. It emits the map when the source
// completes. If the source is empty, it emits an empty map.
// Play: https://go.dev/play/p/FiF83XYB0ba
func ToMapI[T any, K comparable, V any](mapper func(item T, index int64) (K, V)) func(Observable[T]) Observable[map[K]V] {
	_ = "STUB: not implemented"
	return nil
}

// ToMapIWithContext collects all items from the observable into a map. It is a sink
// operator so it emit a single value. It emits the map when the source
// completes. If the source is empty, it emits an empty map.
// Play: https://go.dev/play/p/FiF83XYB0ba
func ToMapIWithContext[T any, K comparable, V any](mapper func(ctx context.Context, item T, index int64) (K, V)) func(Observable[T]) Observable[map[K]V] {
	_ = "STUB: not implemented"
	return nil
}

// ToChannel materializes and forward all items from the observable into a
// channel. It is a sink operator so it emit a single value. It emits the
// channel when the source completes. If the source is empty, it emits an
// empty channel. The channel will be closed when the source completes or
// emit an error.
// Play: https://go.dev/play/p/WMKa26sirV0
func ToChannel[T any](size int) func(Observable[T]) Observable[<-chan Notification[T]] {
	_ = "STUB: not implemented"
	return nil
}

// Send the channel to the observer, because
// it's going to detach the upstream from the downstream.
// The next operator might be long-running.

// This is a workaround to avoid a race condition between the
// destination.NextWithContext() and the destination.CompleteWithContext()
// on empty source.

// Send the channel to the observer, after the goroutine is started.
// Because the observer might call be long-running.
// But on empty source, the destination.CompleteWithContext() might be
// called before the goroutine is started.
