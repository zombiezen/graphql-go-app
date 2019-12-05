#!/bin/bash
# Copyright 2019 Ross Light
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

set -e

if [[ $# -ne 1 || "$1" == '--help' ]]; then
  echo "usage: $(basename "$0") IMPORTPATH"
  exit 64
fi
IMPORTPATH="$1"
NAME="$(basename "$1")"

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
cd "$DIR"
find . -type f | grep -v '/\.[^/]\|/node_modules/\|set-project-name.sh\|README.md' | \
  xargs sed -i -e "s:zombiezen.com/go/graphql-go-app:$IMPORTPATH:g" -e "s/graphql-go-app/$NAME/g" \
    .dockerignore \
    .gitignore \
    .graphqlconfig
