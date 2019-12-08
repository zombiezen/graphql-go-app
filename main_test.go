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
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"zombiezen.com/go/graphql-server/graphql"
	"zombiezen.com/go/log/testlog"
)

func TestGreeting(t *testing.T) {
	app, err := newApplication("schema.graphql", "client")
	if err != nil {
		t.Fatal(err)
	}
	ctx := testlog.WithTB(context.Background(), t)
	response := app.server.Execute(ctx, graphql.Request{
		Query: `{ greeting }`,
	})
	if len(response.Errors) > 0 {
		t.Fatal(response.Errors)
	}
	got := response.Data.ValueFor("greeting").GoValue()
	want := "Hello, World!"
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("greeting (-want +got):\n%s", diff)
	}
}

func TestMain(m *testing.M) {
	testlog.Main(nil)
	os.Exit(m.Run())
}
