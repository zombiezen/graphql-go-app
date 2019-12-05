// Copyright 2019 Ross Light
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"io"
	"sync"

	"go.opencensus.io/trace"
	"zombiezen.com/go/log"
)

type logWriter struct {
	prefix string
	flag   log.Flags

	mu  sync.Mutex
	out io.Writer
	buf []byte
}

func (w *logWriter) Log(ctx context.Context, ent log.Entry) {
	span := trace.FromContext(ctx)
	defer w.mu.Unlock()
	w.mu.Lock()
	w.buf = append(w.buf[:0], w.prefix...)
	if span != nil {
		w.buf = append(w.buf, span.String()...)
		w.buf = append(w.buf, ':', ' ')
	}
	w.buf = ent.Append(w.buf, w.flag)
	w.buf = append(w.buf, '\n')
	w.out.Write(w.buf)
}

func (w *logWriter) LogEnabled(log.Entry) bool {
	return true
}
