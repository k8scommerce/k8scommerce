# Flux

```sh
flux bootstrap github \
--owner=$GITHUB_USER \
--branch=master \
--repository=services \
--path=./k8s/gitops/clusters/local-dev \
--personal=false
```