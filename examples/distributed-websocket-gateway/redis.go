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
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"github.com/samber/ro"
)

var (
	rdb           *redis.Client
	pubsubChannel string = "exchange"
)

func init() {
	fmt.Println("Connecting to Redis...")

	url := "redis://@localhost:6379/0"
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	rdb = redis.NewClient(opts)
}

func publishSink(roomID string, msg string) error { _ = "STUB: not implemented"; return nil }

func subscribeSource(destination ro.Observer[lo.Tuple2[string, string]]) ro.Teardown {
	_ = "STUB: not implemented"
	return *new(ro.Teardown)
}
