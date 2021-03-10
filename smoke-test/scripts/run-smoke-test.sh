#!/bin/bash

cd smoke-test/$1
container_id=$(docker container ls | grep smoke-tests | awk '{print $1}')
sed "s/command: kubectl/command: docker exec $container_id kubectl/g" plonk.yaml.bk > plonk.yaml

pods=$(docker exec -it --privileged $container_id kubectl get nodes)
while grep -q "localhost:8080" $pods
do
    pods=$(docker exec -it --privileged $container_id kubectl get nodes)
done

plonk deploy