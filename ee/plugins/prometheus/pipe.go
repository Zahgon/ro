// Copyright 2025 samber.
//
// Licensed as an Enterprise License (the "License"); you may not use
// this file except in compliance with the License. You may obtain
// a copy of the License at:
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.ee.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package roprometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/samber/ro"
)

// wrapPipeWithObservability is a wrapper around PipeOpX that adds multiple observability operators to the wrapPipeWithObservability.
func wrapPipeWithObservability[First any, Last any](collector *prometheusCollector, operators func(ro.Observable[First]) ro.Observable[Last]) func(ro.Observable[First]) ro.Observable[Last] {
	_ = "STUB: not implemented"

	// // Add input notification counter between source and first operator.
	// IncCounterOnNext[First](collector.NotificationsInTotal.With(prometheus.Labels{})),
	// // Track the time it takes for a notification to traverse from the source observable to the destination observer.
	// ObserveNextLag[First](collector.NotificationLagSeconds.With(prometheus.Labels{})),
	// // Track the time it takes for an operator to process a notification.
	// InitOperatorProcessingTimeObserver[First](),
	return nil
}

// observeBeforePipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - InitOperatorProcessingTimeObserver
//   - ObserveNextLag

// PipeX(...)

// // Add output notification counter between last operator and final subscriber.
// IncCounterOnNext[Last](collector.NotificationsOutTotal.With(prometheus.Labels{})),
// // Add subscriptions counter at the end: it will be the first called operator on subscription.
// IncCounterOnSubscription[Last](collector.SubscriptionsTotal),

// observeAfterPipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - IncCounterOnSubscription

// Pipe1 is a typesafe 🎉 implementation of Pipe, that takes a source and 1 operator.
func Pipe1[A any, B any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
) (ro.Observable[B], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe2 is a typesafe 🎉 implementation of Pipe, that takes a source and 2 operators.
func Pipe2[A any, B any, C any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
) (ro.Observable[C], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe3 is a typesafe 🎉 implementation of Pipe, that takes a source and 3 operators.
func Pipe3[A any, B any, C any, D any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
) (ro.Observable[D], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe4 is a typesafe 🎉 implementation of Pipe, that takes a source and 4 operators.
func Pipe4[A any, B any, C any, D any, E any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
) (ro.Observable[E], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe5 is a typesafe 🎉 implementation of Pipe, that takes a source and 5 operators.
func Pipe5[A any, B any, C any, D any, E any, F any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
) (ro.Observable[F], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe6 is a typesafe 🎉 implementation of Pipe, that takes a source and 6 operators.
func Pipe6[A any, B any, C any, D any, E any, F any, G any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
) (ro.Observable[G], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe7 is a typesafe 🎉 implementation of Pipe, that takes a source and 7 operators.
func Pipe7[A any, B any, C any, D any, E any, F any, G any, H any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
) (ro.Observable[H], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe8 is a typesafe 🎉 implementation of Pipe, that takes a source and 8 operators.
func Pipe8[A any, B any, C any, D any, E any, F any, G any, H any, I any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
) (ro.Observable[I], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

func Pipe9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
) (ro.Observable[J], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

func Pipe10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
) (ro.Observable[K], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe11 is a typesafe 🎉 implementation of Pipe, that takes a source and 11 operators.
func Pipe11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
) (ro.Observable[L], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe12 is a typesafe 🎉 implementation of Pipe, that takes a source and 12 operators.
func Pipe12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
) (ro.Observable[M], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe13 is a typesafe 🎉 implementation of Pipe, that takes a source and 13 operators.
func Pipe13[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
) (ro.Observable[N], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe14 is a typesafe 🎉 implementation of Pipe, that takes a source and 14 operators.
func Pipe14[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
) (ro.Observable[O], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe15 is a typesafe 🎉 implementation of Pipe, that takes a source and 15 operators.
func Pipe15[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
) (ro.Observable[P], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe16 is a typesafe 🎉 implementation of Pipe, that takes a source and 16 operators.
func Pipe16[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
) (ro.Observable[Q], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe17 is a typesafe 🎉 implementation of Pipe, that takes a source and 17 operators.
func Pipe17[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
) (ro.Observable[R], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe18 is a typesafe 🎉 implementation of Pipe, that takes a source and 18 operators.
func Pipe18[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
) (ro.Observable[S], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe19 is a typesafe 🎉 implementation of Pipe, that takes a source and 19 operators.
func Pipe19[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
) (ro.Observable[T], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe20 is a typesafe 🎉 implementation of Pipe, that takes a source and 20 operators.
func Pipe20[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
	operator20 func(ro.Observable[T]) ro.Observable[U],
) (ro.Observable[U], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe21 is a typesafe 🎉 implementation of Pipe, that takes a source and 21 operators.
func Pipe21[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
	operator20 func(ro.Observable[T]) ro.Observable[U],
	operator21 func(ro.Observable[U]) ro.Observable[V],
) (ro.Observable[V], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe22 is a typesafe 🎉 implementation of Pipe, that takes a source and 22 operators.
func Pipe22[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
	operator20 func(ro.Observable[T]) ro.Observable[U],
	operator21 func(ro.Observable[U]) ro.Observable[V],
	operator22 func(ro.Observable[V]) ro.Observable[W],
) (ro.Observable[W], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe23 is a typesafe 🎉 implementation of Pipe, that takes a source and 23 operators.
func Pipe23[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
	operator20 func(ro.Observable[T]) ro.Observable[U],
	operator21 func(ro.Observable[U]) ro.Observable[V],
	operator22 func(ro.Observable[V]) ro.Observable[W],
	operator23 func(ro.Observable[W]) ro.Observable[X],
) (ro.Observable[X], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license

// Pipe24 is a typesafe 🎉 implementation of Pipe, that takes a source and 24 operators.
func Pipe24[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any, N any, O any, P any, Q any, R any, S any, T any, U any, V any, W any, X any, Y any](
	collectorConfig CollectorConfig,
	source ro.Observable[A],
	operator1 func(ro.Observable[A]) ro.Observable[B],
	operator2 func(ro.Observable[B]) ro.Observable[C],
	operator3 func(ro.Observable[C]) ro.Observable[D],
	operator4 func(ro.Observable[D]) ro.Observable[E],
	operator5 func(ro.Observable[E]) ro.Observable[F],
	operator6 func(ro.Observable[F]) ro.Observable[G],
	operator7 func(ro.Observable[G]) ro.Observable[H],
	operator8 func(ro.Observable[H]) ro.Observable[I],
	operator9 func(ro.Observable[I]) ro.Observable[J],
	operator10 func(ro.Observable[J]) ro.Observable[K],
	operator11 func(ro.Observable[K]) ro.Observable[L],
	operator12 func(ro.Observable[L]) ro.Observable[M],
	operator13 func(ro.Observable[M]) ro.Observable[N],
	operator14 func(ro.Observable[N]) ro.Observable[O],
	operator15 func(ro.Observable[O]) ro.Observable[P],
	operator16 func(ro.Observable[P]) ro.Observable[Q],
	operator17 func(ro.Observable[Q]) ro.Observable[R],
	operator18 func(ro.Observable[R]) ro.Observable[S],
	operator19 func(ro.Observable[S]) ro.Observable[T],
	operator20 func(ro.Observable[T]) ro.Observable[U],
	operator21 func(ro.Observable[U]) ro.Observable[V],
	operator22 func(ro.Observable[V]) ro.Observable[W],
	operator23 func(ro.Observable[W]) ro.Observable[X],
	operator24 func(ro.Observable[X]) ro.Observable[Y],
) (ro.Observable[Y], prometheus.Collector) {
	_ = "STUB: not implemented"
	return nil, *new(prometheus.Collector)
}

// no license

// with license
