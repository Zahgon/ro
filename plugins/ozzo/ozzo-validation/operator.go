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

import (
	"errors"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/samber/ro"
)

var (
	ErrValidatable            = errors.New("value does not implement ozzo.Validatable")
	ErrValidatableWithContext = errors.New("value does not implement ozzo.ValidatableWithContext")
)

func Validate[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[Result[T]] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStruct[T any]() func(ro.Observable[T]) ro.Observable[Result[T]] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateWithContext[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[Result[T]] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStructWithContext[T any]() func(ro.Observable[T]) ro.Observable[Result[T]] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateOrError[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStructOrError[T any]() func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateOrErrorWithContext[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStructOrErrorWithContext[T any]() func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateOrSkip[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStructOrSkip[T any]() func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateOrSkipWithContext[T any](rules ...ozzo.Rule) func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}

func ValidateStructOrSkipWithContext[T any]() func(ro.Observable[T]) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
