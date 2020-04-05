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
# Install dependencies.
# Only copy in package.json and package-lock.json to increase Docker
# cache hit rate.
COPY client/package.json client/package-lock.json /src/
WORKDIR /src
RUN npm install
# Now build the whole tree.
COPY client /src
RUN npm run build
RUN mv /src/dist/main.js.map /src/

FROM gcr.io/distroless/base-debian10
COPY --from=build /src/<APPLICATION_NAME> /<APPLICATION_NAME>
COPY --from=build /src/schema.graphql /schema.graphql
COPY --from=clientbuild /src/dist /client
COPY --from=clientbuild /src/main.js.map /main.js.map
ENV PORT 8080
EXPOSE 8080
ENTRYPOINT ["/<APPLICATION_NAME>", "-client=/client", "-schema=/schema.graphql"]
