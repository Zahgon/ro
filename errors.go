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

package ro

import (
	"errors"
	"time"
)

// @TODO: custom error type ?
func recoverValueToError(e any) error { _ = "STUB: not implemented"; return nil }

func recoverUnhandledError(cb func()) { _ = "STUB: not implemented"; return }

var (
	//nolint:revive
	ErrRangeWithStepWrongStep                       = errors.New("ro.RangeWithStep: step must be greater than 0")
	ErrRangeWithStepAndIntervalWrongStep            = errors.New("ro.RangeWithStepAndInterval: step must be greater than 0")
	ErrFirstEmpty                                   = errors.New("ro.First: empty")
	ErrLastEmpty                                    = errors.New("ro.Last: empty")
	ErrHeadEmpty                                    = errors.New("ro.First: empty")
	ErrTailEmpty                                    = errors.New("ro.Last: empty")
	ErrTakeWrongCount                               = errors.New("ro.Take: count must be greater or equal to 0")
	ErrTakeLastWrongCount                           = errors.New("ro.TakeLast: count must be greater than 0")
	ErrSkipWrongCount                               = errors.New("ro.Skip: count must be greater or equal to 0")
	ErrSkipLastWrongCount                           = errors.New("ro.SkipLast: count must be greater than 0")
	ErrElementAtWrongNth                            = errors.New("ro.ElementAt: nth must be greater or equal to 0")
	ErrElementAtNotFound                            = errors.New("ro.ElementAt: nth element not found")
	ErrElementAtOrDefaultWrongNth                   = errors.New("ro.ElementAtOrDefault: nth must be greater or equal to 0")
	ErrRepeatWrongCount                             = errors.New("ro.Repeat: count must be greater or equal to 0")
	ErrRepeatWithIntervalWrongCount                 = errors.New("ro.RepeatWithInterval: count must be greater or equal to 0")
	ErrRepeatWithWrongCount                         = errors.New("ro.RepeatWith: count must be greater or equal to 0")
	ErrBufferWithCountWrongSize                     = errors.New("ro.BufferWithCount: size must be greater than 0")
	ErrBufferWithTimeWrongDuration                  = errors.New("ro.BufferWithTime: duration must be greater than 0")
	ErrBufferWithTimeOrCountWrongSize               = errors.New("ro.BufferWithTimeOrCount: size must be greater than 0")
	ErrBufferWithTimeOrCountWrongDuration           = errors.New("ro.BufferWithTimeOrCount: duration must be greater than 0")
	ErrClampLowerLessThanUpper                      = errors.New("ro.Clamp: lower must be less than or equal to upper")
	ErrToChannelWrongSize                           = errors.New("ro.ErrToChannelWrongSize: size must be greater or equal to 0")
	ErrPoolWrongSize                                = errors.New("ro.Pool: size must be greater than 0")
	ErrSubscribeOnWrongBufferSize                   = errors.New("ro.SubscribeOn: buffer size must be greater than 0")
	ErrObserveOnWrongBufferSize                     = errors.New("ro.ObserveOn: buffer size must be greater than 0")
	ErrDetachOnWrongMode                            = errors.New("ro.detachOn: unexpected detach mode")
	ErrUnicastSubjectConcurrent                     = errors.New("ro.UnicastSubject: a single subscriber accepted")
	ErrConnectableObservableMissingConnectorFactory = errors.New("ro.ConnectableObservable: missing connector factory")
)

func newUnsubscriptionError(err error) error { _ = "STUB: not implemented"; return nil }

type unsubscriptionError struct {
	err error
}

func (e *unsubscriptionError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *unsubscriptionError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newObservableError(err error) error { _ = "STUB: not implemented"; return nil }

type observableError struct {
	err error
}

func (e *observableError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *observableError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newObserverError(err error) error { _ = "STUB: not implemented"; return nil }

type observerError struct {
	err error
}

func (e *observerError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *observerError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newTimeoutError(duration time.Duration) error { _ = "STUB: not implemented"; return nil }

type timeoutError struct {
	duration time.Duration
}

func (e *timeoutError) Error() string { _ = "STUB: not implemented"; return "" }

func newCastError[T, U any]() error { _ = "STUB: not implemented"; return nil }

type castError[T any, U any] struct{}

func (e *castError[T, U]) Error() string { _ = "STUB: not implemented"; return "" }

func newPipeError(msg string, args ...any) error { _ = "STUB: not implemented"; return nil }

type pipeError struct {
	err error
}

func (e *pipeError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *pipeError) Unwrap() error { _ = "STUB: not implemented"; return nil }
