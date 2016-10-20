// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package core

import "go.uber.org/fx/core/metrics"

// A StubModule implements the Module interface for testing
type StubModule struct {
	Host             ServiceHost
	InitError        error
	TypeVal, NameVal string
	StartError       error
	StopError        error
	Running          bool
	TrafficReporter  metrics.TrafficReporter
}

var _ Module = &StubModule{}

// NewStubModule generates a Module for use in testing
func NewStubModule() *StubModule {
	return &StubModule{
		TrafficReporter: &metrics.LoggingTrafficReporter{},
	}
}

// Initialize fakes an init call on the module
func (s *StubModule) Initialize(host ServiceHost) error {
	s.Host = host
	return s.InitError
}

// Start mimics startup
func (s *StubModule) Start(ready chan<- struct{}) <-chan error {
	errs := make(chan error, 1)
	if s.StartError != nil {
		errs <- s.StartError
	}
	ready <- struct{}{}
	return errs
}

// Type returns the type of the module
func (s *StubModule) Type() string { return s.TypeVal }

// Name returns the name of the module
func (s *StubModule) Name() string { return s.NameVal }

// IsRunning returns the current running state
func (s *StubModule) IsRunning() bool { return s.Running }

// Stop stops the module
func (s *StubModule) Stop() error { return s.StopError }

// Reporter returns the traffic reporter for the module
func (s *StubModule) Reporter() metrics.TrafficReporter { return s.TrafficReporter }