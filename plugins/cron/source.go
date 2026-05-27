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

package rocron

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/samber/ro"
)

type ScheduleJob struct {
	Counter int
	Time    time.Time
}

// NewScheduler creates a new observable that emits a notification on
// each tick of the scheduler.
//
// Example: trigger a job every night at 23:42.
//
//	NewScheduler(gocron.CronJob("42 23 * * *"), false).Subscribe(...)
func NewScheduler(job gocron.JobDefinition) ro.Observable[ScheduleJob] {
	_ = "STUB: not implemented"
	return nil
}

// start the scheduler
