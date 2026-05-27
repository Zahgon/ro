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

package rolog

import (
	"github.com/samber/ro"
)

func Log[T any]() func(ro.Observable[T]) ro.Observable[T] { _ = "STUB: not implemented"; return nil }

func LogWithPrefix[T any](prefix string) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// bearer:disable go_lang_logger_leak

// bearer:disable go_lang_logger_leak

// bearer:disable go_lang_logger_leak

func FatalOnError[T any]() func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func FatalOnErrorWithPrefix[T any](prefix string) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

// bearer:disable go_lang_logger_leak
