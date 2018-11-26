#!/bin/sh

set -e

if [ ! -d /cgroup/cpu ]; then
    mkdir -p /cgroup/cpu
    mount -t cgroup cgroup /cgroup/cpu -o rw,relatime,cpu
fi

exec "$@"
