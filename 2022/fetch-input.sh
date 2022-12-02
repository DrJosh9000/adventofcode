
day=$(date +%-d)

curl --cookie "session=${AOC_SESSION}" \
  -o "inputs/${day}.txt" \
  "https://adventofcode.com/2022/day/${day}/input"