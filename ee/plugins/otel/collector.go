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
	"github.com/samber/ro/ee/internal/introspection"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const (
	serviceName = "github.com/samber/ro"

	labelNamePipe             = "pipe.name"
	LabelNamePipePosition     = "pipe.position"
	labelNameOperator         = "operator.name"
	labelNameOperatorPosition = "operator.position"
	labelNameOperatorIndex    = "operator.index"
)

func newOtelCollector(opts CollectorConfig, pipeDescription *introspection.FunDesc) *otelCollector {
	_ = "STUB: not implemented"
	return nil
}

type otelCollector struct {
	config CollectorConfig

	tracer trace.Tracer
	meter  metric.Meter
	logger log.Logger

	// @TODO:
	//   - SubscriptionDurationSeconds (summary): time between subscription and complete/error/unsuscribe
	//   - OperatorNotificationInflationRate (gauge): from an operator POV, the difference between the rx and tx notifications.
	SubscriptionsTotal            metric.Int64Counter
	NotificationsInTotal          metric.Int64Counter
	NotificationsOutTotal         metric.Int64Counter
	NotificationLagSeconds        metric.Float64Histogram
	OperatorProcessingTimeSeconds metric.Float64Histogram
}

func (c *otelCollector) initTracer() { _ = "STUB: not implemented"; return }

func (c *otelCollector) initMeter() { _ = "STUB: not implemented"; return }

func (c *otelCollector) initLogger() { _ = "STUB: not implemented"; return }

func (c *otelCollector) registerMetrics() { _ = "STUB: not implemented"; return }
