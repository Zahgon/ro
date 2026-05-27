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
	"github.com/samber/ro/ee/internal/introspection"
)

const (
	labelNamePipe         = "pipe"
	LabelNamePipePosition = "pipe_position"
	// LabelNameSubscriptionID   = "subscription_id"
	labelNameOperator         = "operator"
	labelNameOperatorPosition = "operator_position"
	labelNameOperatorIndex    = "operator_index"
)

var _ prometheus.Collector = (*prometheusCollector)(nil)

func newPrometheusCollector(opts CollectorConfig, pipeDescription *introspection.FunDesc) *prometheusCollector {
	_ = "STUB: not implemented"
	return nil
}

type prometheusCollector struct {
	// @TODO:
	//   - SubscriptionDurationSeconds (summary): time between subscription and complete/error/unsuscribe
	//   - OperatorNotificationInflationRate (gauge): from an operator POV, the difference between the rx and tx notifications.
	SubscriptionsTotal            prometheus.Counter
	NotificationsInTotal          *prometheus.CounterVec
	NotificationsOutTotal         *prometheus.CounterVec
	NotificationLagSeconds        *prometheus.SummaryVec
	OperatorProcessingTimeSeconds *prometheus.SummaryVec
}

func (c *prometheusCollector) Describe(ch chan<- *prometheus.Desc) {
	_ = "STUB: not implemented"
	return
}

func (c *prometheusCollector) Collect(ch chan<- prometheus.Metric) {
	_ = "STUB: not implemented"
	return
}
