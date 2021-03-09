# plonk!

An opinionated [`kubernetes`][kubernetes] deployment orchestrator. Plonk uses templates to build comprehensive deploy(manifest) files and executes them against a kubernetes cluster. These templates range from `http-service`, `grafana`, `prometheus`, `redis`, `postgres` and anything that you need to build a distributed infrastructure on top of kubernetes.

## Commands

* `plonk init <new project>` - Creates a new project.
* `plonk deploy <environment>` - Deploys the configured project using the passed in environment (or `production` if none was specified).
* `plonk show <environment>` - Retrieves the running pods for the specified environment (or `production` if none was specified).
* `plonk logs <environment> <component>` - Retrieves the logs for the specified environment (or `production` if none was specified) and for the specific component provided (if none was provided `plonk` will list all the available ones).
* `plonk destroy <environment>` - Deletes completely the generated `namespace` for the passed in environment (or `production` if none was specified). 
    
!!! danger 
    `plonk destroy` is a highly destructive command. `plonk` will ask for verification before executing; once the command has been confirmed it is unlikely that it can be stopped or reverted

## Templates

`plonk!` comes with a list of preconfigured templates:

- [alertmanager][alertmanager-docs]
- [docker-registry][docker-registry-docs]
- [drone][drone-docs]
- [gocd][gocd-docs]
- [grafana][grafana-docs]
- [memcached][memcached-docs]
- [prometheus][prometheus-docs]
- [http-service][http-service-docs]

<!-- Below are all the links listed in this page -->
[alertmanager-docs]:#
[docker-registry-docs]:#
[drone-docs]:#
[gocd-docs]:#
[grafana-docs]:#
[memcached-docs]:#
[prometheus-docs]:#
[http-service-docs]:#
[kubernetes]:http://kubernetes.github.io