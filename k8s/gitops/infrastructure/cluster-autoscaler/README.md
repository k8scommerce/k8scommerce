# cluster-autoscaler

`cluster-autoscaler` manages scaling additional nodes up if we run out of pod slots.  It can also scale them down.

* [Homepage](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler)
* [AWS EKS Docs on Cluster Autoscaler](https://docs.aws.amazon.com/eks/latest/userguide/cluster-autoscaler.html)
* [Cluster Autoscaler Docs on AWS EKS](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/cloudprovider/aws/README.md)
* [Cluster Autoscaler FAQ](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md)
* [Releases](https://github.com/kubernetes/autoscaler/releases) (look for only `cluster-autoscaler` releases, not helm charts or others; they recommend release number match kubernetes release number).

## **IMPORTANT**

We have this setup to use [IAM Roles for Service Accounts](https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html).  We don't need to create the `ServiceAccount` for `cluster-autoscaler` because this is created by `eksctl` when the cluster is created.  We just need to bind to it.

You'll need to / want to override some settings on the deployment for a given cluster including, possibly:

* `nodeSelector` - pins the deployment to a type of node
* [deployment.command](https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-the-parameters-to-ca)
  * Note that the one command line arg `node-group-auto-discovery` specifies the tag which calls out the cluster name, so it is specific to a cluster.
