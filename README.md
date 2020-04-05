# GraphQL Go web application template

![Build](https://github.com/zombiezen/graphql-go-app/workflows/Build/badge.svg)

This repository contains a small web application using GraphQL and Go on the
backend and React with [Apollo Client][] in TypeScript for the frontend. It's
designed as a starting point for bootstrapping new web applications with Go and
TypeScript.

If you find issues, please report them on the [graphql-server issue tracker][].

[Apollo Client]: https://www.apollographql.com/docs/react/
[graphql-server issue tracker]: https://github.com/zombiezen/graphql-server/issues

## Features

Go server:

-  Reflection-based GraphQL server using [`zombiezen.com/go/graphql-server`][]
-  HTTP request logs sent to stdout
-  Debug logging sent to stderr
-  Graceful termination
-  Unit test example

TypeScript client:

-  Preconfigured [React Router][] and [Apollo Client][]
-  [GraphQL Playground][] available by visiting `/client/playground.html`
-  Unit test harness using [Jest][] and [Enzyme][]
-  Webpack configuration ready for [code splitting][] and
   [build-time GraphQL parsing][]

Development:

-  Provided [Dockerfile][] for consistent builds and deployment to any
   container-based environment like Kubernetes
-  Provided Heroku configuration for simple deployment
-  Continuous Integration (CI) via [GitHub Actions][]
-  [Tasks][VSCode Tasks] and editor settings for Visual Studio Code

[build-time GraphQL parsing]: https://www.apollographql.com/docs/react/integrations/webpack/
[code splitting]: https://reactjs.org/docs/code-splitting.html
[Dockerfile]: https://github.com/zombiezen/graphql-go-app/blob/master/Dockerfile
[Enzyme]: https://airbnb.io/enzyme/
[GitHub Actions]: https://github.com/features/actions
[GraphQL Playground]: https://github.com/prisma-labs/graphql-playground
[Jest]: https://jestjs.io/
[React Router]: https://reacttraining.com/react-router/web/
[VSCode Tasks]: https://code.visualstudio.com/docs/editor/tasks
[`zombiezen.com/go/graphql-server`]: https://pkg.go.dev/mod/zombiezen.com/go/graphql-server

## Getting Started

Before getting started, you will need [Go][] 1.13 or later, [Node.js][] 12, and
Adobe's [go-starter][] tool. You may optionally want to install [Docker][].

Use the `go-starter` tool to create your repository.

```shell
go-starter zombiezen/graphql-go-app PROJECT
cd PROJECT
```

To run the app locally:

```shell
# In one terminal:
cd client && npm install && npm run watch

# In another terminal:
go build && ./<APPLICATION_NAME>
```

[Docker]: https://www.docker.com/get-started
[generate a repository]: https://github.com/zombiezen/graphql-go-app/generate
[Go]: https://golang.org/dl/
[go-starter]: https://github.com/adobe/go-starter
[Node.js]: https://nodejs.org/en/download/

### Notable Files

Once you've got your environment set up, these are the files you will most
likely want to edit next:

-  [schema.graphql][]: Service definition written in the [GraphQL schema language][].
   If you make changes, run `npm run graphql-codegen` from the `client`
   directory to update the TypeScript type definitions.
-  [main.go][]: Server code. Can be broken up into multiple Go files, like with
   any Go package.
-  [main_test.go][]: Server test code
-  [client/src/components/App.tsx][]: Top-level `<App>` React component
-  [client/src/components/App.test.tsx][]: Unit tests for the `<App>` component
-  [client/dist/style.css][]: Stylesheet
-  [client/dist/index.html][]: Entrypoint for the TypeScript client application.

[GraphQL schema language]: https://graphql.org/learn/schema/
[main.go]: https://github.com/zombiezen/graphql-go-app/blob/master/main.go
[main_test.go]: https://github.com/zombiezen/graphql-go-app/blob/master/main_test.go
[client/dist/index.html]: https://github.com/zombiezen/graphql-go-app/blob/master/client/dist/index.html
[client/dist/style.css]: https://github.com/zombiezen/graphql-go-app/blob/master/client/dist/style.css
[client/src/components/App.tsx]: https://github.com/zombiezen/graphql-go-app/blob/master/client/src/components/App.tsx
[client/src/components/App.test.tsx]: https://github.com/zombiezen/graphql-go-app/blob/master/client/src/components/App.test.tsx
[schema.graphql]: https://github.com/zombiezen/graphql-go-app/blob/master/schema.graphql

## License

[Apache 2.0](https://github.com/zombiezen/graphql-go-app/blob/master/LICENSE)
