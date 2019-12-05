# Copyright 2019 Ross Light
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
#     http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

FROM golang:1.13 as build
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org/
# Warm up the module cache.
# Only copy in go.mod and go.sum to increase Docker cache hit rate.
COPY go.mod go.sum /src/
WORKDIR /src
RUN go mod download
# Now build the whole tree.
COPY . /src
RUN go build

FROM node:12.13-slim AS clientbuild
COPY client /src
WORKDIR /src
RUN npm install && npm run build

FROM gcr.io/distroless/base
COPY --from=build /src/graphql-go-app /graphql-go-app
COPY --from=build /src/schema.graphql /schema.graphql
COPY --from=clientbuild /src/dist /client
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/graphql-go-app", "-client=/client", "-schema=/schema.graphql"]
