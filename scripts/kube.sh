#!/bin/bash

goctl kube deploy \
    -name cart \
    -namespace k8scommerce \
    -image 127.0.0.1:5000/cart:latest \
    -secret k8scommerce-db \
    -requestCpu 128 \
    -requestMem 128 \
    -replicas 2 \
    -port 8080  \
    -minReplicas 2 \
    -maxReplicas 15 \
    -o cart.yaml

