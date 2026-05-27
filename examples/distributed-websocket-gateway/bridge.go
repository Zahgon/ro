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

package main

import (
	"github.com/samber/lo"
	"github.com/samber/ro"
)

var bridgeInstance *bridge

func initStreams() { _ = "STUB: not implemented"; return }

func closeStreams() { _ = "STUB: not implemented"; return }

func newEchange() *bridge { _ = "STUB: not implemented"; return nil }

// websocket->redis (downstream)

// redis->websocket (upstream)

type bridge struct {
	upstream      ro.Subject[lo.Tuple2[string, string]]
	downstream    ro.Subject[lo.Tuple2[string, string]]
	subscriptions ro.Subscription
}

func (e *bridge) Publish(roomID string, msg string) { _ = "STUB: not implemented"; return }

func (e *bridge) Subscribe(roomID string, destination ro.Observer[string]) ro.Subscription {
	_ = "STUB: not implemented"
	return *new(ro.Subscription)
}

// exclude messages from other rooms

func (e *bridge) Close() { _ = "STUB: not implemented"; return }
