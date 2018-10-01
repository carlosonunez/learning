#!/usr/bin/env bash
if ! which docker &>/dev/null
then
  >&2 echo "ERROR: Docker not found."
  exit 1
fi

docker run --rm \
  --tty \
  --volume "$PWD:/app" \
  --workdir "/app" \
  ruby:alpine ruby solution.rb $@
