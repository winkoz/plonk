#!/bin/bash

docker="docker run -it -v `PWD`:/plonk -w /plonk brianskarda/adr-tools-docker:latest"
cmd="$docker adr $@"
echo "$cmd"
eval "$cmd" || exit 0
