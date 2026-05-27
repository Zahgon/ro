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

// bearer:disable go_gosec_blocklist_md5

// bearer:disable go_gosec_blocklist_sha1

// Convenience variables for easy access
var (
	// StringHash provides hash functions for strings
	StringHash = StringHashers{}
	// BytesHash provides hash functions for byte slices
	BytesHash = BytesHashers{}
)

// StringHashers provides common hash functions for strings
type StringHashers struct{}

// FNV64a returns a hash function using FNV-1a 64-bit
func (StringHashers) FNV64a() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// FNV64 returns a hash function using FNV-1 64-bit
func (StringHashers) FNV64() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// FNV32a returns a hash function using FNV-1a 32-bit
func (StringHashers) FNV32a() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// FNV32 returns a hash function using FNV-1 32-bit
func (StringHashers) FNV32() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// SHA256 returns a hash function using SHA-256
func (StringHashers) SHA256() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// SHA1 returns a hash function using SHA-1
func (StringHashers) SHA1() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// bearer:disable go_lang_weak_hash_sha1, go_gosec_crypto_weak_crypto

// SHA512 returns a hash function using SHA-512
func (StringHashers) SHA512() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// MD5 returns a hash function using MD5
func (StringHashers) MD5() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// bearer:disable go_lang_weak_hash_md5, go_gosec_crypto_weak_crypto

// MapHash returns a hash function using maphash
func (StringHashers) MapHash() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// CRC32 returns a hash function using CRC-32
func (StringHashers) CRC32() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// CRC64 returns a hash function using CRC-64
func (StringHashers) CRC64() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// Adler32 returns a hash function using Adler-32
func (StringHashers) Adler32() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// Jenkins returns a hash function using Jenkins hash (one-at-a-time)
func (StringHashers) Jenkins() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// DJB2 returns a hash function using DJB2 algorithm
func (StringHashers) DJB2() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// hash * 33 + c

// SDBM returns a hash function using SDBM algorithm
func (StringHashers) SDBM() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// Loselose returns a hash function using the "lose lose" algorithm
func (StringHashers) Loselose() func(string) uint64 { _ = "STUB: not implemented"; return nil }

// BytesHashers provides common hash functions for byte slices
type BytesHashers struct{}

// FNV64a returns a hash function using FNV-1a 64-bit for byte slices
func (BytesHashers) FNV64a() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// FNV64 returns a hash function using FNV-1 64-bit for byte slices
func (BytesHashers) FNV64() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// FNV32a returns a hash function using FNV-1a 32-bit for byte slices
func (BytesHashers) FNV32a() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// FNV32 returns a hash function using FNV-1 32-bit for byte slices
func (BytesHashers) FNV32() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// SHA256 returns a hash function using SHA-256 for byte slices
func (BytesHashers) SHA256() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// SHA1 returns a hash function using SHA-1 for byte slices
func (BytesHashers) SHA1() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// bearer:disable go_lang_weak_hash_sha1, go_gosec_crypto_weak_crypto

// SHA512 returns a hash function using SHA-512 for byte slices
func (BytesHashers) SHA512() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// MD5 returns a hash function using MD5 for byte slices
func (BytesHashers) MD5() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// bearer:disable go_lang_weak_hash_md5, go_gosec_crypto_weak_crypto

// MapHash returns a hash function using maphash for byte slices
func (BytesHashers) MapHash() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// CRC32 returns a hash function using CRC-32 for byte slices
func (BytesHashers) CRC32() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// CRC64 returns a hash function using CRC-64 for byte slices
func (BytesHashers) CRC64() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// Adler32 returns a hash function using Adler-32 for byte slices
func (BytesHashers) Adler32() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// Jenkins returns a hash function using Jenkins hash (one-at-a-time) for byte slices
func (BytesHashers) Jenkins() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// DJB2 returns a hash function using DJB2 algorithm for byte slices
func (BytesHashers) DJB2() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// hash * 33 + c

// SDBM returns a hash function using SDBM algorithm for byte slices
func (BytesHashers) SDBM() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }

// Loselose returns a hash function using the "lose lose" algorithm for byte slices
func (BytesHashers) Loselose() func([]byte) uint64 { _ = "STUB: not implemented"; return nil }
