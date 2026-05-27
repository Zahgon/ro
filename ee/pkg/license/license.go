// Copyright 2025 samber.
//
// Licensed as an Enterprise License (the "License"); you may not use
// this file except in compliance with the License. You may obtain
// a copy of the License at:
//
// https://github.com/samber/ro/blob/main/licenses/LICENSE.ee.md
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rolicense

import (
	_ "embed"
	"time"

	"github.com/samber/ro/internal/xatomic"
)

//go:embed public.key
var embeddedPublicKey []byte

// licensePrefix is a prefix on license strings to make them easily recognized.
const licensePrefix = "ro-"
const v0LicensePrefix = licensePrefix + "00"

type Type string

const (
	TypeCommunity  Type = "community"
	TypeEnterprise Type = "enterprise"
)

type Environment string

const (
	EnvironmentDevelopment Environment = "development"
	EnvironmentProduction  Environment = "production"
)

var defaultLicense = &License{
	Type:           TypeCommunity,
	LicenseID:      "",
	OrganizationID: "",
	Environment:    EnvironmentDevelopment,
	ExpiresAt:      time.Time{},
	Online:         false,
}

type License struct {
	Type           Type        `json:"t"`
	LicenseID      string      `json:"id"`
	OrganizationID string      `json:"own"`
	Environment    Environment `json:"env"`
	ExpiresAt      time.Time   `json:"exp"` // zero value -> no expiration

	// Online or offline: when online, the license is validated remotely.
	// For future use.
	Online bool `json:"o"`
}

// IsExpired checks if a license has expired
func (l *License) IsExpired() bool { _ = "STUB: not implemented"; return false }

// GenerateLicense creates a signed license string using the provided private key
func GenerateLicense(license *License, privateKeyPEM []byte) (string, error) {
	_ = "STUB: not implemented"
	// Parse the private key using the keys package
	return "", nil
}

// Serialize the license to JSON

// Encode the container

// Add the license prefix

// ParseLicense decodes and verifies a license string using the embedded public key
func ParseLicense(key string) (*License, error) {
	_ = "STUB: not implemented"
	// Check if the license has the correct prefix
	return nil, nil
}

// Remove the prefix and decode the container

// Decode the license data

// Unmarshal the license

// On expiration, the license will be set to nil on the next call to GetLicense().
var currentLicense = xatomic.NewPointer[License](nil)

// SetLicense sets the current license.
// On invalid license or expired license, an error is returned. This error might be
// ignored by the caller, since the license is not critical to the application in some cases.
func SetLicense(key string) (err error) { _ = "STUB: not implemented"; return nil }

// GetLicense returns the current license.
// If the license is expired, it will be set to nil on the next call to GetLicense().
// Warning: calling this function too often might cause performance issues.
func GetLicense() *License { _ = "STUB: not implemented"; return nil }

// Commented, because it's too expensive to call IsExpired() too often.
// if license.IsExpired() {
// 	currentLicense.Store(defaultLicense)
// 	return defaultLicense
// }

// IsEnterpriseEnabled returns true if the current license is an enterprise license.
// Warning: calling this function too often might cause performance issues.
func IsEnterpriseEnabled() bool { _ = "STUB: not implemented"; return false }
