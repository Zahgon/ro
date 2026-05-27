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

package rohttpclient

import (
	"net/http"

	"github.com/samber/ro"
)

// HTTPRequest sends a http request and returns the response. It's a pull-based operator.
//
// A http status code >= 400 is not considered an error.
//
// Don't forget to call resp.Body.Close() when you're done with the response.
func HTTPRequest(req *http.Request, client *http.Client) ro.Observable[*http.Response] {
	_ = "STUB: not implemented"
	return nil
}

func HTTPRequestJSON[T any](req *http.Request, client *http.Client) ro.Observable[T] {
	_ = "STUB: not implemented"
	return nil
}
