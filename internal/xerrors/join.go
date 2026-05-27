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

package xerrors

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Join returns an error that wraps the given errors.
// Any nil error values are discarded.
// Join returns nil if every value in errs is nil.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline
// between each string.
//
// A non-nil error returned by Join implements the Unwrap() []error method.
func Join(errs ...error) error { _ = "STUB: not implemented"; return nil }

//nolint:staticcheck

type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	_ = "STUB: not implemented"
	// Since Join returns nil if every value in errs is nil,
	// e.errs cannot be empty.
	return ""
}

// At this point, b has at least one byte '\n'.

// return unsafe.String(&b[0], len(b))

func (e *joinError) Unwrap() []error { _ = "STUB: not implemented"; return nil }
