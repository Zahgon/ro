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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

func concatAttributes[T any](kv ...[]T) []T { _ = "STUB: not implemented"; return nil }

func traceWithAttributes(kv ...[]attribute.KeyValue) trace.SpanStartEventOption {
	_ = "STUB: not implemented"
	return *new(trace.SpanStartEventOption)
}

func metricWithAttributes(kv ...[]attribute.KeyValue) metric.MeasurementOption {
	_ = "STUB: not implemented"
	return *new(metric.MeasurementOption)
}

func newRecord(msg string, severity log.Severity, kv ...log.KeyValue) log.Record {
	_ = "STUB: not implemented"
	return *new(log.Record)
}
