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

package rotesting

import (
	"context"
	"testing"

	"github.com/samber/ro"
)

// @TODO: Add new methods:
// - ExpectDurationEpsilon
// - ExpectDurationLessThan
// - ExpectDurationGreaterThan
// - ExpectDurationInRange

var _ AssertSpec[int] = (*assertImpl[int])(nil)

type assertImpl[T any] struct {
	t          *testing.T
	assertions []gotestingAssertion[T]
	source     ro.Observable[T]
}

type gotestingAssertion[T any] struct {
	notification ro.Notification[T]
	msgAndArgs   []any
}

// Assert creates a new instance of test. It is used to assert the behavior of an
// observable sequence.
//
// Inspired by Flux.
func Assert[T any](t *testing.T) AssertSpec[T] {
	_ = "STUB: not implemented" //nolint:thelper
	return nil
}

func (t *assertImpl[T]) popAssertion() (gotestingAssertion[T], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (t *assertImpl[T]) equal(expected, actual any, msgAndArgs ...any) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:errcheck,forcetypeassert

func (t *assertImpl[T]) hasErrorOrCompletionNotification() bool {
	_ = "STUB: not implemented"
	return false
}

// Source sets the source observable to test.
func (t *assertImpl[T]) Source(source ro.Observable[T]) AssertSpec[T] {
	_ = "STUB: not implemented"
	return nil
}

// ExpectNext expects the next value to be emitted by the source observable.
// It fails the test if the next value is not emitted. If the source observable
// emits an error or completes, it fails the test.
func (t *assertImpl[T]) ExpectNext(value T, msgAndArgs ...any) AssertSpec[T] {
	_ = "STUB: not implemented"
	return nil
}

// ExpectNextSeq expects the next values to be emitted by the source observable.
// It fails the test if the next values are not emitted. If the source observable
// emits an error or completes, it fails the test.
func (t *assertImpl[T]) ExpectNextSeq(values ...T) AssertSpec[T] {
	_ = "STUB: not implemented"
	return nil
}

// msgAndArgs:   []any{"expected '%v' value", (any)(values[i])},

// ExpectError expects the source observable to emit an error. It fails the test
// if the source observable emits a value or completes. If the source observable
// emits an error, it compares the error with the expected error. If the error
// is not equal to the expected error, it fails the test.
func (t *assertImpl[T]) ExpectError(err error, msgAndArgs ...any) AssertSpec[T] {
	_ = "STUB: not implemented"
	return nil
}

// ExpectComplete expects the source observable to complete. It fails the test
// if the source observable emits a value or an error.
func (t *assertImpl[T]) ExpectComplete(msgAndArgs ...any) AssertSpec[T] {
	_ = "STUB: not implemented"
	return nil
}

// Verify subscribes to the source observable and verifies the assertions.
// It fails the test if the source observable emits a value, an error, or completes
// before all assertions are verified.
func (t *assertImpl[T]) Verify() { _ = "STUB: not implemented"; return }

// VerifyWithContext subscribes to the source observable and verifies the assertions.
// It fails the test if the source observable emits a value, an error, or completes
// before all assertions are verified.
func (t *assertImpl[T]) VerifyWithContext(ctx context.Context) { _ = "STUB: not implemented"; return }
