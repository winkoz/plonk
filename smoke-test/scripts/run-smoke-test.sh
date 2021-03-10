#!/bin/bash

rm -rf smoke-test/sandbox/.
cp -r smoke-test/tests/$1/* smoke-test/sandbox/.
cd smoke-test/sandbox/

container_id=$(docker container ls | grep smoke-tests | awk '{print $1}')
sed "s/command: kubectl/command: docker exec $container_id kubectl/g" plonk.yaml.bk > plonk.yaml

while read -r line ; do
    ready="$(echo $line | grep '==> /var/log/d')"
    test -z $ready || break
done < <(docker logs --tail 300 --follow $container_id)

echo "🚦 Cluster ready"

plonk deploy