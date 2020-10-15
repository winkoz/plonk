# 5. Use apex/log wrapper

Date: 2020-10-15

## Status

Accepted

## Context

With the idea in mind to not "marry" the implementation of our application to an specific logger we encapsulated the log calls in a wrapper.
Originally we were using the `prometheus` logger but that proved wrong because of how the callstack was being inferred; causing all our logs to print the callsite our log wrapper.

The immediate solution was to copy the implementation of the `prometheus` logger into our "wrapper" rendering it no longer a wrapper.

To avoid this we needed to find a better solution or a different, more powerful logger that allowed us to wrap it without forcing us to copy its implementation.

## Decision

Use [`apex/log`][apex] as a better logging engine instead of `prometheus` logger copy.
Initially we thought about using [`klog`](https://github.com/kubernetes/klog) but it is now a "life support" project and no new big features will be added. Instead we decided to use [apex][apex] as it is a modern structured logging system with support for different handlers and ability to add custom ones.

## Consequences

Starting now we should always wrap our public facing functions with `log.StartTrace()` and `log.StopTrace()` calls to be able to log to console the time it's taking us to compute different operations.

This will print to console from `Info` level and above unfortunately which is not ideal but is a good compromise for the time being.
Besides that we are introducing the `--verbosity` flag to `plonk` which takes values from `[debug, info, warn, error, fatal]` in order to decide what level to print out; defaulting to `info`.

[apex]:https://github.com/apex/log