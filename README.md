# GraphQL Go web application template

This repository contains a small web application using GraphQL and Go on the
backend and React with [Apollo Client][] in TypeScript for the frontend. It's
designed as a starting point for bootstrapping new web applications with Go and
TypeScript.

[Apollo Client]: https://www.apollographql.com/docs/react/

## Getting Started

1. Use the GitHub web interface to [generate a repository][] based on
   this template.
2. Clone the repository to your local machine.
3. Use the `set-project-name.sh` script to use your project name.

```shell
git clone https://github.com/MYNAME/PROJECT.git
cd PROJECT
./set-project-name.sh github.com/MYNAME/PROJECT
```

You can also try running this repository with [Google Cloud Run][]:

[![Run on Google Cloud](https://deploy.cloud.run/button.svg)](https://deploy.cloud.run)

[generate a repository]: https://github.com/zombiezen/graphql-go-app/generate
[Google Cloud Run]: https://cloud.google.com/run

## Features

Go server:

-  HTTP request logs sent to stdout
-  Debug logging sent to stderr
-  Graceful termination
-  Unit test example

TypeScript client:

-  Preconfigured [Apollo Client][]
-  [GraphQL Playground][] available by visiting `/client/playground.html`
-  Unit test harness using [Jest][] and [Enzyme][]

Development:

-  Provided [Dockerfile][] for consistent builds and deployment to any
   container-based environment like Kubernetes
-  Provided Heroku configuration for simple deployment
-  Continuous Integration (CI) via [GitHub Actions][]

[Dockerfile]: https://github.com/zombiezen/graphql-go-app/blob/master/Dockerfile
[Enzyme]: https://airbnb.io/enzyme/
[GitHub Actions]: https://github.com/features/actions
[GraphQL Playground]: https://github.com/prisma-labs/graphql-playground
[Jest]: https://jestjs.io/

## License

[Apache 2.0](https://github.com/zombiezen/graphql-go-app/blob/master/LICENSE)
