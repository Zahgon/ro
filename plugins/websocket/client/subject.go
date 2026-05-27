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

package rowebsocketclient

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/samber/ro"
)

type Serializer[T any] func(T) ([]byte, error)

type Deserializer[T any] func([]byte) (T, error)

type WebsocketSubjectConfig[In any, Out any] struct {
	URL          string
	Headers      map[string]string
	Serializer   Serializer[In]
	Deserializer Deserializer[Out]
	Dialer       *websocket.Dialer

	// Connector is a function that returns a Subject[Out].
	// This is useful when you want to use a different Subject implementation.
	// For example, you could use a ReplaySubject to replay the last N messages.
	OutputConnector func() ro.Subject[Out]
	// ResetOnError    bool
	// ResetOnComplete bool
}

// NewWebsocketSubject creates a websocket subject that can both send and receive messages from a websocket endpoint.
func NewWebsocketSubject[In any, Out any](config WebsocketSubjectConfig[In, Out]) *websocketSubject[In, Out] {
	_ = "STUB: not implemented"
	return nil
}

// Set default output connector

var _ ro.Subject[string] = (*websocketSubject[string, string])(nil)
var _ ro.Observer[string] = (*websocketSubject[string, int])(nil)
var _ ro.Observable[string] = (*websocketSubject[int, string])(nil)

type websocketSubject[In any, Out any] struct {
	config WebsocketSubjectConfig[In, Out]
	input  ro.Observer[[]byte]
	output ro.Subject[Out]
	conn   *websocket.Conn
	mu     sync.RWMutex
}

// Implements ro.Observable[Out]
func (ws *websocketSubject[In, Out]) Subscribe(destination ro.Observer[Out]) ro.Subscription {
	_ = "STUB: not implemented"
	return *new(ro.Subscription)
}

// Implements ro.Observable[Out]
func (ws *websocketSubject[In, Out]) SubscribeWithContext(ctx context.Context, destination ro.Observer[Out]) ro.Subscription {
	_ = "STUB: not implemented"
	return *new(ro.Subscription)
}

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) Next(value In) { _ = "STUB: not implemented"; return }

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) NextWithContext(ctx context.Context, value In) {
	_ = "STUB: not implemented"
	return
}

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) Error(err error) { _ = "STUB: not implemented"; return }

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) ErrorWithContext(ctx context.Context, err error) {
	_ = "STUB: not implemented"
	return
}

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) Complete() { _ = "STUB: not implemented"; return }

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) CompleteWithContext(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) IsClosed() bool { _ = "STUB: not implemented"; return false }

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) HasThrown() bool { _ = "STUB: not implemented"; return false }

// Implements ro.Observer[In]
func (ws *websocketSubject[In, Out]) IsCompleted() bool { _ = "STUB: not implemented"; return false }

// Implements ro.Subject[Out]
func (ws *websocketSubject[In, Out]) HasObserver() bool { _ = "STUB: not implemented"; return false }

// Implements ro.Subject[Out]
func (ws *websocketSubject[In, Out]) CountObservers() int { _ = "STUB: not implemented"; return 0 }

// Implements ro.Subject[Out]
func (ws *websocketSubject[In, Out]) AsObservable() ro.Observable[Out] {
	_ = "STUB: not implemented"

	// Implements ro.Subject[In]
	return nil
}

func (ws *websocketSubject[In, Out]) AsObserver() ro.Observer[In] {
	_ = "STUB: not implemented"

	// Connect establishes the WebSocket connection
	return nil
}

func (ws *websocketSubject[In, Out]) connect() (ro.Observer[[]byte], ro.Subject[Out], error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Already connected

// Set up headers

// Dial the WebSocket connection

// conn.SetWriteDeadline(time.Now().Add(?))

// Start reading messages

func (ws *websocketSubject[In, Out]) readMessages(conn *websocket.Conn, output ro.Subject[Out]) {
	_ = "STUB: not implemented"
	return
}

// Deserialize and emit the message
