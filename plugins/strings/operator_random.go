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

package rostrings

import (
	"math"

	"github.com/samber/ro"
)

var (
	LowerCaseLettersCharset = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperCaseLettersCharset = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LettersCharset          = append(LowerCaseLettersCharset, UpperCaseLettersCharset...)
	NumbersCharset          = []rune("0123456789")
	AlphanumericCharset     = append(LettersCharset, NumbersCharset...)
	SpecialCharset          = []rune("!@#$%^&*()_+-=[]{}|;':\",./<>?")
	AllCharset              = append(AlphanumericCharset, SpecialCharset...)

	maximumCapacity = math.MaxInt>>1 + 1
)

// nearestPowerOfTwo returns the nearest power of two.
func nearestPowerOfTwo(cap int) int { _ = "STUB: not implemented"; return 0 }

func random(size int, charset []rune) string {
	_ = "STUB: not implemented"
	// see https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	return ""
}

// Calculate the number of bits required to represent the charset,
// e.g., for 62 characters, it would need 6 bits (since 62 -> 64 = 2^6)

// Determine the corresponding bitmask,
// e.g., for 62 characters, the bitmask would be 111111.

// Available count, since xrand.Int64() returns a non-negative number, the first bit is fixed, so there are 63 random bits
// e.g., for 62 characters, this value is 10 (63 / 6).

// Generate the random string in a loop.

// Regenerate the random number if all available bits have been used

// Select a character from the charset

// Shift the bits to the right to prepare for the next character selection,
// e.g., for 62 characters, shift by 6 bits.

// Decrease the remaining number of uses for the current random number.

// Random generates a random string of the specified size using the specified charset.
// Play: https://go.dev/play/p/7oDIGRxvrGt
func Random[T any](size int, charset []rune) func(destination ro.Observable[T]) ro.Observable[string] {
	_ = "STUB: not implemented"
	return nil
}
