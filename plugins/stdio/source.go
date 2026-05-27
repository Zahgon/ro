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

package rostdio

import (
	"io"

	"github.com/samber/ro"
)

const IOReaderBufferSize = 1024

// NewIOReader creates an observable that reads bytes from an io.Reader.
// Play: https://go.dev/play/p/b75Poy3EVYn
func NewIOReader(reader io.Reader) ro.Observable[[]byte] { _ = "STUB: not implemented"; return nil }

// NewIOReaderLine creates an observable that reads lines from an io.Reader.
// Play: https://go.dev/play/p/oMv2jYVSLqd
func NewIOReaderLine(reader io.Reader) ro.Observable[[]byte] { _ = "STUB: not implemented"; return nil }

// NewStdReader creates an observable that reads bytes from standard input.
func NewStdReader() ro.Observable[[]byte] { _ = "STUB: not implemented"; return nil }

// NewStdReaderLine creates an observable that reads lines from standard input.
func NewStdReaderLine() ro.Observable[[]byte] { _ = "STUB: not implemented"; return nil }

// NewPrompt creates an observable that reads user input after displaying a prompt.
func NewPrompt(prompt string) ro.Observable[[]byte] { _ = "STUB: not implemented"; return nil }

// Print the prompt to stdout

// Read from stdin

// Send the input as a byte slice
