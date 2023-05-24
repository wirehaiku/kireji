#! /usr/bin/env fish
# scripts/build.fish: Build cross-platform binaries using Gox.

set arch   "386 amd64"
set os     "darwin linux windows"
set osarch "!darwin/386"

gox -arch=$arch -os=$os -osarch=$osarch -verbose  .
