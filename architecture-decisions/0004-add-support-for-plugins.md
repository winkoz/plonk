# 4. Add support for plugins

Date: 2020-09-19

## Status

Accepted

## Context

A big part of this project is to allow different organizations or projects with different configurations to easily modify Plonk's actions to 
adhere to their needs, meaning different values, structure and frecuency of kubernetes manifests or other file types that we decide to end up supporting, but at the same time we don't want to expose the projects to the complexity of these files, in k8s case, the complexity of the manifests.

## Decision

For this we decided not to create templates for each project inside the projects directory but rather use the templates directly in the different Plonk's commands. For this we will add plugins, which will be a set of templates, the list of plugins used by the project will be defined in the `plonk.yml` file. These plugins will modify the behaviour of 3 components, the scaffolder, the deployer and the variables. 

Examples of plugins will be `k8s` and `GOCD` and for now they will be the only ones supported. The current plonk commands `scaffold`/`generate` and `deploy` will use the templates in different ways to perform their actions.

## Consequences

We need to add support for a templating engine, remove the "scripts" from the deploy folder and setup the templates in a way that differenciates the manifest templates from the scaffolding templates.
