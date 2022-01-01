# ecr-credentials-sync

This is based on the [documentation here](https://fluxcd.io/docs/guides/image-update/#imagerepository-cloud-providers-authentication).

This should be a temporary hack workaround.  Currently Flux doesn't correctly implement authentication to all the cloud provider repositories (including AWS ECR), so we need to run this `DaemonSet` that periodically refreshes the token.

Without this, we run into [this issue](https://github.com/fluxcd/image-reflector-controller/issues/139).

Once they implement native authentication, then we can likely remove this `DaemonSet` from the cluster, as the node itself has permissions to pull/read from all of our repositories.

This has to be run per namespace, b/c the `ImageRepository` and `ImagePolicy` are per namespace.
