# Welcome to plonk!

`plonk!` is `DevOps` as a tool, allowing engineers to focus on developing their systems and applications and letting `plonk!` handle communicating and controlling their `kubernetes` cluster.

An opinionated kubernetes deployment orchestrator. Plonk uses templates to build comprehensive deploy(manifest) files and executes them against a kubernetes cluster. These templates range from `http-service`, `grafana`, `prometheus`, `redis`, `postgres` and anything that you need to build a distributed infrastructure on top of kubernetes.

---

## How to Install

Download the code and execute the following commands in the path where you downloaded it

!!! example
    
    ```bash
    # Build plonk
    $ make build

    # Set the bin path to your PATH variable
    $ export PATH=PATH:/path/to/the/bin/folder/inside/the/plonk/directory

    # check the version
    $ plonk -v 
    ```

!!! success "You should see the following output:"
    ```console
    Plonk 0.0.1
    ```

## Getting Started

### Setup plonk in your project

Plonk requires a specific folder structure and files to work, in order to simplify the setup it has the ability to bootstrap the necessary files into your project

!!! example
    ```console
    $ cd /path/to/your/project
    $ plonk init name-of-your-project
    ```

??? note "Now your should see the following structure:"
    ```console
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

!!! example
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

###  Verify the manifest file

Check the manifest file that will be executed against the cluster:

```console
$ plonk diff
```

??? success "The previous command should create a manifest file in `deploy/deploy.yaml` with a structure similar to the following:"
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

### Deploy your application to the cluster

The following command will execute `kubectl apply -f deploy/deploy.yaml` against the cluster you currently have configured.

```console
$ plonk deploy
```

!!! example "After this you can check the list of pods with:"
    ```console
    $ plonk show
    ```
    ```json
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


## Configure your Kubectl command

Plonk uses `kubectl` to execute commands against your k8s cluster, please make sure you have your `kubectl` correctly configured and can access your cluster. Please [see this](https://kubernetes.io/docs/tasks/tools/install-kubectl/) if you have more questions. 


If you are wrapping your kubectl in another executable or docker container you can configure this on the `plonk.yaml` file by changing the `command` property: 

```yaml
name: plonk-test
command: your-command-goes-here
```

