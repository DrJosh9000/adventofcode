#!/bin/zsh
set -eu

year="$(date +%-Y)"
solution="$(mktemp)"
trap "{ rm "${solution}" }" EXIT;

go build -o "${solution}" "./${year}/$1"
(time "${solution}") | tee /dev/stderr | tail -n1 | pbcopy
