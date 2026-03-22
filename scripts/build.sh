#!/usr/bin/env bash
set -euo pipefail

mkdir -p bin
go build -o bin/belajar-golang ./cmd/belajar-golang
