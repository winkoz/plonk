# 5. Use klog logger wrapper

Date: 2020-10-15

## Status

Accepted

## Context

With the idea in mind to not "marry" the implementation of our application to an specific logger we encapsulated the log calls in a wrapper.
Originally we were using the `prometheus` logger but that proved wrong because of how the callstack was being inferred; causing all our logs to print the callsite our log wrapper.

The immediate solution was to copy the implementation of the `prometheus` logger into our "wrapper" rendering it no longer a wrapper.

To avoid this we needed to find a better solution or a different, more powerful logger that allowed us to wrap it without forcing us to copy its implementation.

## Decision

Use [`klog`](https://github.com/kubernetes/klog) is `kubernetes` take at `glog` (`Google`'s `C++` highly efficient logger) in `go` instead of `prometheus` logger and wrap its calls in a single place in our codebase.

The logger is fast, efficient and powerful enough to have more verbosity levels giving us more control on how to log. Allows us also to log to disk and control the callstack indirection printing the call site of the log call correctly instead of our log wrapper.

Due to how short-lived our application is writing to disk doesn't work unless the application is up and running for at least 5 seconds. In order to do that we need to flush the buffer everytime we log a line.

## Consequences

Going forward all calls to log information will pass through our wrapper and internally will call `klog` with an extra depth level to avoid logging the wrong call-site.
We default also to write logs to `.plonk/logs` which requires the directory to exist prior to running the app. (This only if we want to log to a directory, otherwise we can just write to a single file).
