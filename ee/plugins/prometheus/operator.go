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

// IncCounterOnNext is a pipe operator that increments a counter
// when a new Next() notification is sent to the destination observer.
//
// It adds a short lock for each Next() notification: 30ns for the
// prometheus/client_golang locks.
func IncCounterOnNext[T any](counter prometheus.Counter) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnError is a pipe operator that increments a counter
// when a new Error() notification is sent to the destination observer.
//
// It adds a short lock for each Error() notification: 30ns for the
// prometheus/client_golang locks.
func IncCounterOnError[T any](counter prometheus.Counter) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnComplete is a pipe operator that increments a counter
// when a new Complete() notification is sent to the destination observer.
//
// It adds a short lock for each Complete() notification: 30ns for the
// prometheus/client_golang locks.
func IncCounterOnComplete[T any](counter prometheus.Counter) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnSubscription is a pipe operator that increments a counter
// when a new subscription is created.
//
// It adds a short lock for each subscription: 30ns for the
// prometheus/client_golang locks.
func IncCounterOnSubscription[T any](counter prometheus.Counter) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ObserveNextLag is a pipe operator that tracks the time it takes for a notification
// to traverse from the source observable to the destination observer.
// It mesures the time the source pauses while waiting for the destination
// to process the notification.
//
// It adds a short lock for each Next() notification:
//   - 2x 15ns for the time tracking
//   - 30ns for the prometheus/client_golang locks
func ObserveNextLag[T any](summaryOrHistogram prometheus.Observer) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// @TODO: track error and completion processing time?

type checkpointCtx struct{}

// observeBeforePipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - ObserveNextLag
//
// Aggregating avoid the need to add a short lock for each notification.
func observeBeforePipe[T any](counterOnNext prometheus.Counter, summaryOrHistogram prometheus.Observer) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Operator: Operator: IncCounterOnNext

/////////////////////////////

// Operator: observeOnNotification

/////////////////////////////

// Operator: ObserveNextLag

// observeOperatorProcessingTime is a pipe operator that observes the processing time of an operator.
func observeOperatorProcessingTime[T any](summaryOrHistogram prometheus.ObserverVec, operatorName string, operatorPosition string, operatorIndex int) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// @TODO: track error and completion processing time?

// observeAfterPipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - IncCounterOnSubscription
//
// Aggregating avoid the need to add a short lock for each notification.
func observeAfterPipe[T any](counterOnNext prometheus.Counter, counterOnSubscription prometheus.Counter) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Operator: IncCounterOnSubscription

// Operator: IncCounterOnNext
