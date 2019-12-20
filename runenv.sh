#!/bin/bash

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
exec env - "${vars[@]}" "$@"
