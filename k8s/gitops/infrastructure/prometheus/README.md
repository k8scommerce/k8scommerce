## Prometheus Deployment on K8S

To **deploy** Prometheus on K8S cluster, you can run: `kubectl apply -k ./sources`

To **delete** Prometheus deployment, you can run: `kubectl delete -k ./sources`

## Access Prometheus Resources

The Prometheus server can be accessed via port `80` on the following DNS name from within the cluster:
```
prometheus-nab-giftcard-server.prometheus.svc.cluster.local
```

You can also grab the Prometheus server URL by running these commands in the same shell:
```
  export POD_NAME=$(kubectl get pods \
    --namespace prometheus -l "app=prometheus,component=server" \
    -o jsonpath="{.items[0].metadata.name}")

  kubectl --namespace prometheus port-forward $POD_NAME 9090

```

The Prometheus alertmanager can be accessed via port `80` on the following DNS name from within the cluster:
```
prometheus-nab-giftcard-alertmanager.prometheus.svc.cluster.local
```


Grab the Alertmanager URL by running these commands in the same shell:
```
  export POD_NAME=$(kubectl get pods \
    --namespace prometheus -l "app=prometheus,component=alertmanager" \
    -o jsonpath="{.items[0].metadata.name}")

  kubectl --namespace prometheus port-forward $POD_NAME 9093
```


The Prometheus PushGateway can be accessed via port `9091` on the following DNS name from within your cluster:
```
prometheus-nab-giftcard-pushgateway.prometheus.svc.cluster.local
```

Grab the PushGateway URL by running these commands in the same shell:
```
  export POD_NAME=$(kubectl get pods \
    --namespace prometheus -l "app=prometheus,component=pushgateway" \
    -o jsonpath="{.items[0].metadata.name}")
    
  kubectl --namespace prometheus port-forward $POD_NAME 9091
```

For more information on running Prometheus, visit: https://prometheus.io/