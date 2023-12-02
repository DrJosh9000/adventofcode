#!/bin/zsh

set -eu

year="$(date +%-Y)"
day="${1:-$(date +%-d)}"

mkdir -p "${year}/inputs"

echo "Fetching ${year}/inputs/${day}.txt..."
curl --cookie "$(cat COOKIE)" \
  -o "${year}/inputs/${day}.txt" \
  "https://adventofcode.com/${year}/day/${day}/input"
