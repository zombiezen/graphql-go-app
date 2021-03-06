#!/bin/bash
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

set -eo pipefail
if [[ $# -lt 2 ]]; then
  echo 'usage: runenv.sh FILE PROGRAM [...]'
  exit 64
fi
env_path="$1"
shift

# If the file doesn't exist, then pass through.
if [[ ! -e "$env_path" ]]; then
  exec "$@"
fi

# Read non-comment lines into vars array.
vars=()
while read -r var; do
  vars+=("$var")
done < <(grep -v '^#\|^\s*$' "$env_path")

# Run program with given environment variables.
exec env -- "${vars[@]}" "$@"
