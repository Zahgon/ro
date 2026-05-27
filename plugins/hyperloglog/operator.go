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

package rohyperloglog

import (
	"fmt"

	"github.com/samber/ro"
)

var ErrInvalidPrecision = fmt.Errorf("rohyperloglog.CountDistinct: precision has to be >= 4 and <= 18")

func CountDistinct[T comparable](precision uint8, sparse bool, hashFunc func(input T) uint64) func(ro.Observable[T]) ro.Observable[uint64] {
	_ = "STUB: not implemented"
	return nil
}

// the error is handled above

func CountDistinctReduce[T comparable](precision uint8, sparse bool, hashFunc func(input T) uint64) func(ro.Observable[T]) ro.Observable[uint64] {
	_ = "STUB: not implemented"
	return nil
}

// the error is handled above
