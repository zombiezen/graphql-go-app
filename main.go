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
	"bytes"
	"context"
	"crypto/sha256"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
	"golang.org/x/sys/unix"
	"zombiezen.com/go/graphql-server/graphql"
	"zombiezen.com/go/graphql-server/graphqlhttp"
	"zombiezen.com/go/log"
)

// application is the root server object that serves both queries and mutations.
type application struct {
	server         *graphql.Server
	entrypointPath string
}

func newApplication(schemaPath, entrypointPath string) (*application, error) {
	schema, err := graphql.ParseSchemaFile(schemaPath, &graphql.SchemaOptions{
		// Set this to true to prevent serving documentation in production.
		IgnoreDescriptions: false,
	})
	if err != nil {
		return nil, err
	}
	app := &application{
		entrypointPath: entrypointPath,
	}
	app.server, err = graphql.NewServer(schema, app, app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Greeting handles the Query.greeting field.
func (app *application) Greeting(ctx context.Context) string {
	return "Hello, World!"
}

// Mutate handles the Mutation.mutate field.
func (app *application) Mutate(ctx context.Context, args map[string]graphql.Value) (graphql.NullString, error) {
	message := args["message"].Scalar()
	log.Infof(ctx, "Mutate message: %q", message)
	return graphql.NullString{}, nil
}

// handleGraphQL processes GraphQL requests.
func (app *application) handleGraphQL(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20 /* 1 MiB */)
	defer r.Body.Close()
	graphqlhttp.NewHandler(app.server).ServeHTTP(w, r)
}

// handleEntrypoint serves the single HTML page entrypoint to the application.
// It's rendered using "html/template" so the server can send configuration to
// the client.
func (app *application) handleEntrypoint(w http.ResponseWriter, r *http.Request) {
	// Handle non-GET/HEAD requests.
	if r.Method != http.MethodGet && r.Method != http.MethodHead && r.Method != http.MethodOptions {
		w.Header().Set("Allow", "GET, HEAD, OPTIONS")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Render entrypoint to buffer.
	ctx := r.Context()
	tmpl, err := template.ParseFiles(app.entrypointPath)
	if err != nil {
		log.Errorf(ctx, "Can't parse HTML template: %v", err)
		http.Error(w, "can't parse entrypoint", http.StatusInternalServerError)
		return
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, struct{}{}); err != nil {
		log.Errorf(ctx, "Can't render HTML template: %v", err)
		http.Error(w, "can't render entrypoint", http.StatusInternalServerError)
		return
	}

	// Serve entrypoint HTML.
	// Make sure to set the Vary header if needed.
	sum := sha256.Sum256(buf.Bytes())
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("ETag", fmt.Sprintf("\"%x\"", sum[:]))
	http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(buf.Bytes()))
}

func newRouter(app *application, clientDir string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", http.HandlerFunc(app.handleGraphQL))
	mux.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir(clientDir))))
	mux.HandleFunc("/", app.handleEntrypoint)
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	clientPath := flag.String("client", "client/dist", "path to client-side resources")
	schemaPath := flag.String("schema", "schema.graphql", "path to GraphQL schema")
	flag.Parse()
	log.SetDefault(&logWriter{
		prefix: "graphql-go-app: ",
		flag:   log.StdFlags,
		out:    os.Stderr,
	})

	ctx := context.Background()
	app, err := newApplication(*schemaPath, filepath.Join(*clientPath, "index.html"))
	if err != nil {
		log.Errorf(ctx, "Read schema: %v", err)
	}
	router := newRouter(app, *clientPath)
	srv := server.New(router, &server.Options{
		RequestLogger: requestlog.NewNCSALogger(os.Stdout, nil),
	})
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, unix.SIGINT, unix.SIGTERM)
	go func() {
		for {
			<-interrupt
			log.Infof(ctx, "Received interrupt; shutting down.")
			srv.Shutdown(context.Background())
		}
	}()
	log.Infof(ctx, "Serving on http://localhost:%s/", port)
	if err := srv.ListenAndServe(":" + port); err != nil {
		log.Errorf(ctx, "Serve: %v", err)
		os.Exit(1)
	}
}
