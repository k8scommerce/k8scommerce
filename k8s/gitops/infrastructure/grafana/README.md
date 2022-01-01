# grafana

Grafana is open source visualization and analytics software. It allows you to query, visualize, alert on, and explore your metrics no matter where they are stored.

* [Homepage](https://grafana.com/oss/grafana/)
* [Releases](https://github.com/grafana/grafana/releases)

Because we use Kustomize [configMapGenerator](https://github.com/kubernetes-sigs/kustomize/blob/master/examples/configGeneration.md) in order to tie the configuration and dashboards to the deployment -- both need to be tied together in the same `kustomization.yaml` or it doesn't work.

This is why we have the `sources/pos-dev` folder that contains the deployment with the spec (as opposed to putting it up in `overlays/teams/monitoring`).
