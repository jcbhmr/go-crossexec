#!/usr/bin/env bash
set -Eeuo pipefail
IFS=$'\n\t'

echo "Wow! Look at this cool app!"
echo "Beep boop"
echo
echo "Look at these arguments. Cool right?"
echo "$0" "$@"
echo
echo "Running 'env'. Those are some interesting variables!"
env
exit 0
