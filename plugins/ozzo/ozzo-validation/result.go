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

package roozzovalidation

// Monad for validation results, inspired by github.com/samber/mo.Result
type Result[T any] struct {
	isErr bool
	value T
	err   error
}

func (r Result[T]) Unwrap() T { _ = "STUB: not implemented"; return *new(T) }

func (r Result[T]) UnwrapOr(defaultValue T) T { _ = "STUB: not implemented"; return *new(T) }

func (r Result[T]) IsOk() bool { _ = "STUB: not implemented"; return false }

func (r Result[T]) IsError() bool { _ = "STUB: not implemented"; return false }

func (r Result[T]) Error() error { _ = "STUB: not implemented"; return nil }

func (r Result[T]) Get() (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

func Ok[T any](value T) Result[T] { _ = "STUB: not implemented"; return nil }

func Err[T any](err error) Result[T] { _ = "STUB: not implemented"; return nil }
