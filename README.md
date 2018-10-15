# cgroup-metrics-reporter

Report ECS CGroup metrics to DataDog

## What does it do?

This Go program is supposed to run as a daemon service on an ECS cluster. It will perform the following steps every seconds:

1. Fetch all ECS tasks running on the instance by querying the ECS introspection endpoint.
2. For each task read the CPU statistics from the related Cgroup (`cpu.stat` file).
3. Send those statistics as metrics to a Datadog Statsd endpoint running on the same machine.

## Usage

    -listen string
        Address and Port to bind health check, in host:port format (default ":8080")
    -namespace string
        Default statsd namespace (default "local.test.")
    -statsd string
        Address and Port to send statsd metrics, in host:port format (default "127.0.0.1:8125")
    -verbose
        Enable verbose logging
