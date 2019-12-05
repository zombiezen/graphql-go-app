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
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"

	"gocloud.dev/server"
	"gocloud.dev/server/requestlog"
	"golang.org/x/sys/unix"
	"zombiezen.com/go/graphql-server/graphql"
	"zombiezen.com/go/graphql-server/graphqlhttp"
	"zombiezen.com/go/log"
)

// application is the root server object that serves both queries and mutations.
type application struct {
	server          *graphql.Server
	insecureGraphQL bool
}

func newApplication(schemaPath string) (*application, error) {
	schema, err := graphql.ParseSchemaFile(schemaPath, nil)
	if err != nil {
		return nil, err
	}
	app := new(application)
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

func (app *application) handleGraphQL(w http.ResponseWriter, r *http.Request) {
	if app.insecureGraphQL {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, OPTIONS")
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20 /* 1 MiB */)
	defer r.Body.Close()
	graphqlhttp.NewHandler(app.server).ServeHTTP(w, r)
}

type fileHandler string

func (fh fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, string(fh))
}

func newRouter(app *application, clientDir string) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/graphql", http.HandlerFunc(app.handleGraphQL))
	mux.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir(clientDir))))
	mux.Handle("/", fileHandler(filepath.Join(clientDir, "index.html")))
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	clientPath := flag.String("client", "client/dist", "path to client-side resources")
	schemaPath := flag.String("schema", "schema.graphql", "path to GraphQL schema")
	insecureGraphQL := flag.Bool("insecure-graphql", false, "enable GraphQL requests from any origin")
	flag.Parse()
	log.SetDefault(&logWriter{
		prefix: "graphql-go-app: ",
		flag:   log.StdFlags,
		out:    os.Stderr,
	})

	ctx := context.Background()
	app, err := newApplication(*schemaPath)
	if err != nil {
		log.Errorf(ctx, "Read schema: %v", err)
	}
	app.insecureGraphQL = *insecureGraphQL
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
