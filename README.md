[![Build Status](https://drone.winkoz.com/api/badges/winkoz/plonk/status.svg?ref=refs/heads/base)](https://drone.winkoz.com/winkoz/plonk)

# plonk

An opinionated kubernetes deployment orchestrator. Plonk uses templates to build comprehensive deploy(manifest) files and executes them against a kubernetes cluster. These templates range from `http-service`, `grafana`, `prometheus`, `redis`, `postgres` and anything that you need to build a distributed infrastructure on top of kubernetes.

---

# How to Install

Download the code and execute the following commands in the path where you downloaded it

```console

# Build plonk
$ make build

# Set the bin path to your PATH variable
$ export PATH=PATH:/path/to/the/bin/folder/inside/the/plonk/directory

# check the version
$ plonk -v 
```

You should see the following output: 

```console
Plonk 0.0.1
```

# Getting Started

## Setup plonk in your project

Plonk requires a specific folder structure and files to work, in order to simplify the setup it has the ability to bootstrap the necessary files into your project

```console
$ cd /path/to/your/project
$ plonk init name-of-your-project
```

Now your should see the following structure:
```
your-project/
├── deploy
│   ├── secrets
│   │   ├── base.yaml
│   │   └── production.yaml
│   └── variables
│       ├── base.yaml
│       └── production.yaml
└── plonk.yaml
```

### Modify the configuration files

You need to change the configuration files, for this example we are going to use a [demo image](https://hub.docker.com/r/strm/helloworld-http/) of an http server that returns a `hello world` message. Your files should look like this:

```yaml
# variables/base.yaml
build:
  NAME: "plonk-test"
  NAMESPACE: "$ENV-plonk-test"
  HOSTNAME: "$ENV.your.domain.url"
  DOCKER_IMAGE: "strm/helloworld-http:latest"
  CONTAINER_PORT: 80
environment:
  APP_ENV: "$ENV"
```

##  Verify the manifest file

Check the manifest file that will be executed against the cluster:

```console
$ plonk diff
```

The previous command should create a manifest file in `deploy/deploy.yaml` with a structure similar to the following: 

```yaml
apiVersion: v1
kind: Namespace
metadata:
  name: plonk-test-production
  labels:
    name: plonk-test-production

---
apiVersion: v1
kind: Service
metadata:
  name: plonk-test-production-service
  namespace: plonk-test-production
  labels:
    app: plonk-test-production
    component: http-service
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 80
  selector:
    app: plonk-test-production
    component: http-service

---
apiVersion: apps/v1
...
```


## Deploy your application to the cluster

The following command will build a docker image, publish that image to a remote docker repository and then execute `kubectl apply -f deploy/deploy.yaml` against the cluster you currently have configured with a manifest file built from the templates you have defined in your `plonk.yml`.

```console
$ plonk deploy
```

After this you can check the list of pods with:

```console
$ plonk show
{
  "apiVersion": "v1",
  "items": [
      {
          "apiVersion": "v1",
          "kind": "Pod",
          "metadata": {
              "annotations": {
                  "kubectl.kubernetes.io/restartedAt": "2021-02-07T22:06:35Z"
              },
              "creationTimestamp": "2021-02-07T22:06:35Z",
              "generateName": "grafana-production-deployment-dbcdfc859-",
              "labels": {
                  "app": "grafana-production",
                  "componen
...
```

If you want to skip the docker build and publish steps you can execute the command with the flag `--skip-build-n-publish`

```console
$ plonk deploy --skip-build-n-publish
```

### Build your application

Whenever Plonk tries to deploy it will do so with an image tag generated based on the current git head, this commands builds the docker image and tags it with that tag name.

```console
$ plonk build [env]
$ plonk build production
```

After that you should be able to verify the tag locally with something like
```console
$ docker images --format "table {{.ID}}\t{{.Repository}}\t{{.Tag}}" [your registry address]]/[the project name]
IMAGE ID       REPOSITORY                       TAG
e3c12ec20d7d   registry.winkoz.com/winkoz-web   production-5f0436344ab18ad9047f69d7bbff1389ee91dabf
```

### Publish your application

The orchestrator cluster or service will need to fetch the tag from some repository, be it docker hub or a private one. This command pushes a built tag to that repository:

```console
$ plonk publish [env]
$ plonk publish production
```

You can check the tag on your docker repository, the tag will be printed by the command like:
```console
$ plonk publish [env]
$ plonk publish production
• info		 Publish tag: registry.winkoz.com/winkoz-web:production-5f0436344ab18ad9047f69d7bbff1389ee91dabf file=publishCommand.go line=63
• info		 Publish executed successfully. file=publishCommand.go line=64
```

# Configure your Kubectl command

Plonk uses `kubectl` to execute commands against your k8s cluster, please make sure you have your `kubectl` correctly configured and can access your cluster. Please [see this](https://kubernetes.io/docs/tasks/tools/install-kubectl/) if you have more questions. 


If you are wrapping your kubectl in another executable or docker container you can configure this on the `plonk.yaml` file by changing the `command` property: 

```yaml
name: plonk-test
command: your-command-goes-here
```
