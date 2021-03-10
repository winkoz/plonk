#!/bin/bash
# Portions Copyright 2016 The Kubernetes Authors All rights reserved.
# Portions Copyright 2018 AspenMesh
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# Based on:
# https://github.com/kubernetes/minikube/tree/master/deploy/docker/localkube-dind

mount --make-shared /

export CNI_BRIDGE_NETWORK_OFFSET="0.0.1.0"
/dindnet &> /var/log/dind.log 2>&1 < /dev/null &

dockerd \
  --host=unix:///var/run/docker.sock \
  --host=tcp://0.0.0.0:2375 \
  &> /var/log/docker.log 2>&1 < /dev/null &


if ! kind get clusters | grep -q "plonk-tester"; then
  kind create cluster --name=plonk-tester
  sleep 4
fi