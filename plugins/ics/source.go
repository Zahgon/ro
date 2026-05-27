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

package roics

import (
	ics "github.com/arran4/golang-ical"
	"github.com/samber/ro"
)

// NewICSFileReader reads events from one or more ICS files.
// @TODO: add glob support
func NewICSFileReader(paths ...string) ro.Observable[*ics.VEvent] {
	_ = "STUB: not implemented"
	return nil
}

// NewICSURLReader reads events from one or more ICS URLs.
func NewICSURLReader(urls ...string) ro.Observable[*ics.VEvent] {
	_ = "STUB: not implemented"
	return nil
}
