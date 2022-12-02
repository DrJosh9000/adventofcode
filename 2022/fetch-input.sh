#!/bin/bash

set -eu

day=${1:-$(date +%-d)}

echo "Fetching inputs/${day}.txt..."
curl --cookie "session=${AOC_SESSION}" \
  -o "inputs/${day}.txt" \
  "https://adventofcode.com/2022/day/${day}/input"