# 2. Adopt paths for custom templates

Date: 2020-09-08

## Status

Accepted

## Context

In order for people to be able to adopt `plonk` we need to support a wide variety of configurations and we can't provide them all out of the box, therefore there should be a way for users to modify the templates.

## Decision

Since distributing the templates as part of the binary will be required and modifying them would be too complex; we decided to enable the ability for users to pass a templates path to the command.

This way if the users do not pass the flag we default to the templates that are shipped with the binary and at the same time support customization for users to provide their own templates and even share them to build a community.

## Consequences

What becomes easier or more difficult to do and any risks introduced by the change that will need to be mitigated.
It will be easier to support customization and growth of the supported configurations by `plonk`; but adds an extra layer of "difficulty" as we now need to work under the assumption that a wrong template path could be passed, the templates folder is empty or the templates passed contain "parsing" errors that we should at least identify and exit the flow early instead of later.
