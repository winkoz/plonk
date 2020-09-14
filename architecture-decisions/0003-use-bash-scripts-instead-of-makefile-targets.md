# 3. Use bash scripts instead of Makefile targets

Date: 2020-09-08

## Status

Accepted

## Context

When introducing the `Makefile` targets to support the creation of `ADR`s we stumble across some issueus that made it quite difficult to manipulate and verify properly.

## Decision

Instead of working around `Makefile` and Gmake we will leverage complex tasks by means of bash scripts that will be called from the `Makefile`.

## Consequences

`Bash` scripts are more flexible and easier to implement and support which will make it easier for us to maintain and add more functionality.
