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

package rootel

import (
	"github.com/samber/ro"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// IncCounterOnNext is a pipe operator that increments a counter
// when a new Next() notification is sent to the destination observer.
func IncCounterOnNext[T any](counter metric.Int64Counter, attributes []attribute.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LogOnNext is a pipe operator that logs a message
// when a new Next() notification is sent to the destination observer.
func LogOnNext[T any](logger log.Logger, severity log.Severity, attributes []log.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnError is a pipe operator that increments a counter
// when a new Error() notification is sent to the destination observer.
func IncCounterOnError[T any](counter metric.Int64Counter, attributes []attribute.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LogOnError is a pipe operator that logs a message
// when a new Error() notification is sent to the destination observer.
func LogOnError[T any](logger log.Logger, severity log.Severity, attributes []log.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnComplete is a pipe operator that increments a counter
// when a new Complete() notification is sent to the destination observer.
func IncCounterOnComplete[T any](counter metric.Int64Counter, attributes []attribute.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LogOnComplete is a pipe operator that logs a message
// when a new Complete() notification is sent to the destination observer.
func LogOnComplete[T any](logger log.Logger, severity log.Severity, attributes []log.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// IncCounterOnSubscription is a pipe operator that increments a counter
// when a new subscription is created.
func IncCounterOnSubscription[T any](counter metric.Int64Counter, attributes []attribute.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// LogOnSubscription is a pipe operator that logs a message
// when a new subscription is created.
func LogOnSubscription[T any](logger log.Logger, severity log.Severity, attributes []log.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// StartTraceOnSubscription is a pipe operator that create a new OTEL trace
// when a new subscription is created.
func StartTraceOnSubscription[T any](collector *otelCollector) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Create a new OTEL trace for each subscription.

// TraceOnError is a pipe operator that records an error in the current OTEL trace.
func TraceOnError[T any](collector *otelCollector) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// ObserveNextLag is a pipe operator that tracks the time it takes for a notification
// to traverse from the source observable to the destination observer.
// It mesures the time the source pauses while waiting for the destination
// to process the notification.
func ObserveNextLag[T any](tracer trace.Tracer, operatorName string, histogram metric.Float64Histogram, attributes []attribute.KeyValue) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// @TODO: track error and completion processing time?

type checkpointCtx struct{}

// observeBeforePipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - ObserveNextLag
//   - TraceOnError
//   - StartTraceOnSubscription
//
// Aggregating avoid the need to add a short lock for each notification.
func observeBeforePipe[T any](collector *otelCollector) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Count the number of messages entering the instrumented ro.PipeX

// Start the checkpoint for the next operator

// Forward the event to the next operator

// Measure the processing time of the operator, using the previous checkpoint

// observeOnNotification creates an operator that measures the processing duration of operators using OTEL metrics
func observeOnNotification[T any](collector *otelCollector, operatorName string, operatorPosition string, operatorIndex int) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Log the call to the operator

// Trace the call to the operator

// Measure the processing time of the operator, using the previous checkpoint

// Update the checkpoint for the next operator

// Forward the event to the next operator

// @TODO: track error and completion processing time?

// observeAfterPipe is the aggregation of the following operators:
//   - IncCounterOnNext
//   - IncCounterOnSubscription
//
// Aggregating avoid the need to add a short lock for each notification.
func observeAfterPipe[T any](collector *otelCollector) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// Instrument subscription

// Count the number of messages leaving the instrumented ro.PipeX

// Log the error in the last operator of the pipeline

// Record the error in the current OTEL trace

// Log completion (different from unsubscription)

// Log unsubscription (different from completion)
