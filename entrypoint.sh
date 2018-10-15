#!/bin/sh

set -e

mkdir -p /cgroup/cpu
mount -t cgroup cgroup /cgroup/cpu -o rw,relatime,cpu

exec "$@"
