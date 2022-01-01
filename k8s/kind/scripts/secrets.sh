#!/bin/sh

set -o errexit


 # Create a TLS secret on disk and encrypt it with Mozilla SOPS.
  # Files are expected to be PEM-encoded.
flux create secret tls localhost-cert \
--namespace=flux-system \
--cert-file=./certs/localhost.crt \
--key-file=./certs/localhost.key 

flux create secret tls local-regsitry-cert \
--namespace=flux-system \
--cert-file=./certs/local-regsitry.crt \
--key-file=./certs/local-regsitry.key 



    # \
    # --export > ./certs/localhost.yaml

# sops --encrypt --encrypted-regex '^(data|stringData)$' \
#     --in-place ./certs/localhost.yaml



# flux create secret tls localhost-cert --ca-file ./ca.pem
flux create image repository product-repo \
--namespace=flux-system \
--cert-ref localhost-cert \
--image localhost:5000/product --interval 1m


# this combination works
flux create secret tls local-registry-cert \
--ca-file ./certs/local-registry.pem \
-n flux-system 

flux create image repository product-repo \
--cert-ref local-registry-cert \
--image local-registry:5000/product \
--interval 1m \
-n flux-system


--from-file=key=./minica/platform-services.cognizant-ai.net/key.pem \
--from-file=cert=./minica/platform-services.cognizant-ai.net/cert.pem


