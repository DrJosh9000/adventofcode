#!/bin/zsh

set -eu

year="$(date +%-Y)"
day="${1:-$(date +%-d)}"

echo "Fetching ${year}/inputs/${day}.txt..."
curl --cookie COOKIE \
  -o "${year}/inputs/${day}.txt" \
  "https://adventofcode.com/${year}/day/${day}/input"
